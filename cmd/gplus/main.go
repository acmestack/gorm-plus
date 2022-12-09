/*
 * Licensed to the AcmeStack under one or more contributor license
 * agreements. See the NOTICE file distributed with this work for
 * additional information regarding copyright ownership.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

var (
	allGenerators = map[string]genall.Generator{
		"gen": Generator{},
	}

	allOutputRules = map[string]genall.OutputRule{
		"dir":       genall.OutputToDirectory(""),
		"none":      genall.OutputToNothing,
		"stdout":    genall.OutputToStdout,
		"artifacts": genall.OutputArtifacts{},
	}

	// optionsRegistry contains all the marker definitions used to process command line options
	optionsRegistry = &markers.Registry{}
)

func init() {
	for genName, gen := range allGenerators {
		// make the generator options marker itself
		defn := markers.Must(markers.MakeDefinition(genName, markers.DescribesPackage, gen))
		if err := optionsRegistry.Register(defn); err != nil {
			panic(err)
		}
		if helpGiver, hasHelp := gen.(genall.HasHelp); hasHelp {
			if help := helpGiver.Help(); help != nil {
				optionsRegistry.AddHelp(defn, help)
			}
		}

		// make per-generation output rule markers
		for ruleName, rule := range allOutputRules {
			ruleMarker := markers.Must(markers.MakeDefinition(fmt.Sprintf("output:%s:%s", genName, ruleName), markers.DescribesPackage, rule))
			if err := optionsRegistry.Register(ruleMarker); err != nil {
				panic(err)
			}
			if helpGiver, hasHelp := rule.(genall.HasHelp); hasHelp {
				if help := helpGiver.Help(); help != nil {
					optionsRegistry.AddHelp(ruleMarker, help)
				}
			}
		}
	}

	// make "default output" output rule markers
	for ruleName, rule := range allOutputRules {
		ruleMarker := markers.Must(markers.MakeDefinition("output:"+ruleName, markers.DescribesPackage, rule))
		if err := optionsRegistry.Register(ruleMarker); err != nil {
			panic(err)
		}
		if helpGiver, hasHelp := rule.(genall.HasHelp); hasHelp {
			if help := helpGiver.Help(); help != nil {
				optionsRegistry.AddHelp(ruleMarker, help)
			}
		}
	}

	// add in the common options markers
	if err := genall.RegisterOptionsMarkers(optionsRegistry); err != nil {
		panic(err)
	}
}

type noUsageError struct{ error }

func main() {
	genCmd := newGenCmd()

	if err := genCmd.Execute(); err != nil {
		if _, noUsage := err.(noUsageError); !noUsage {
			if err := genCmd.Usage(); err != nil {
				panic(err)
			}
		}
		fmt.Fprintf(genCmd.OutOrStderr(), "run `%[1]s %[2]s -w` to see all available markers, or `%[1]s %[2]s -h` for usage\n", genCmd.CalledAs(), strings.Join(os.Args[1:], " "))
		os.Exit(1)
	}
}

func newGenCmd() *cobra.Command {
	genCmd := &cobra.Command{
		Use:   "gen",
		Short: "Generate struct column codes.",
		Long:  "Generate struct column codes.",
		Example: `Generate struct column codes.

	# Run all the generators for a given project
	gplus gen 
`,
		RunE: func(c *cobra.Command, rawOpts []string) error {
			rt, err := genall.FromOptions(optionsRegistry, rawOpts)
			if err != nil {
				return err
			}
			if len(rt.Generators) == 0 {
				return fmt.Errorf("no generators specified")
			}

			if hadErrs := rt.Run(); hadErrs {
				return noUsageError{fmt.Errorf("not all generators ran successfully")}
			}
			return nil
		},
	}
	genCmd.Flags().Bool("help", false, "print out usage and a summary of options")
	return genCmd
}
