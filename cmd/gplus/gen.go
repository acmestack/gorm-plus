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
	"bytes"
	"fmt"
	"github.com/dave/jennifer/jen"
	"go/ast"
	"go/format"
	"go/types"
	"gorm.io/gorm/schema"
	"io"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
	"strings"
)

type StructInfo struct {
	StructName string
	Fields     []string
	FieldTags  []string
}

type Generator struct{}

var enableTypeMarker = markers.Must(markers.MakeDefinition("gplus:column", markers.DescribesType, false))

func (Generator) RegisterMarkers(into *markers.Registry) error {
	if err := markers.RegisterAll(into, enableTypeMarker); err != nil {
		return err
	}
	return nil
}

func enabledOnType(info *markers.TypeInfo) bool {
	if typeMarker := info.Markers.Get(enableTypeMarker.Name); typeMarker != nil {
		return typeMarker.(bool)
	}
	return false
}

func (Generator) Generate(ctx *genall.GenerationContext) error {
	for _, root := range ctx.Roots {
		ctx.Checker.Check(root, func(node ast.Node) bool {
			// ignore interfaces
			_, isIface := node.(*ast.InterfaceType)
			return !isIface
		})

		root.NeedTypesInfo()

		var allStructInfos []StructInfo
		if err := markers.EachType(ctx.Collector, root, func(info *markers.TypeInfo) {
			if !enabledOnType(info) {
				return
			}
			allStructInfos = append(allStructInfos, buildStructInfo(info, root)...)
		}); err != nil {
			root.AddError(err)
			return nil
		}

		if len(allStructInfos) > 0 {
			genFile := buildGenFile(root, allStructInfos)
			var b bytes.Buffer
			err := genFile.Render(&b)
			if err != nil {
				root.AddError(err)
				return nil
			}
			columnContent, err := format.Source(b.Bytes())
			if err != nil {
				root.AddError(err)
				return nil
			}
			writeOut(ctx, root, columnContent)
		}
	}
	return nil
}

func buildGenFile(root *loader.Package, allStructInfos []StructInfo) *jen.File {
	genFile := jen.NewFile(root.Name)
	for _, s := range allStructInfos {
		genFile.Var().Id(s.StructName + "Column").Op("=").Id("struct").Id("{")
		for _, field := range s.Fields {
			genFile.Id(field).String()
		}
		genFile.Id("}").Id("{")
		for i, field := range s.Fields {
			tag := s.FieldTags[i]
			newTag := strings.ReplaceAll(tag, "gorm:", "")
			newTag = strings.ReplaceAll(newTag, "\"", "")
			tagSetting := schema.ParseTagSetting(newTag, ";")
			columnName := tagSetting["COLUMN"]
			if columnName == "" {
				// Use NamingStrategy by default for now
				namingStrategy := schema.NamingStrategy{}
				columnName = namingStrategy.ColumnName("", field)
			}
			columnName = fmt.Sprintf("\"%s\"", columnName)
			genFile.Id(field).Op(":").Id(columnName).Id(",")
		}
		genFile.Id("}")
	}
	return genFile
}

func buildStructInfo(info *markers.TypeInfo, root *loader.Package) []StructInfo {
	var allStructInfos []StructInfo
	typeInfo := root.TypesInfo.TypeOf(info.RawSpec.Name)
	if typeInfo == types.Typ[types.Invalid] {
		root.AddError(loader.ErrFromNode(fmt.Errorf("unknown type %s", info.Name), info.RawSpec))
	}
	structType, ok := typeInfo.Underlying().(*types.Struct)
	if !ok {
		root.AddError(loader.ErrFromNode(fmt.Errorf("%s is not a struct type", info.Name), info.RawSpec))
		return allStructInfos
	}

	structInfo := StructInfo{
		StructName: info.Name,
		Fields:     make([]string, 0, structType.NumFields()),
		FieldTags:  make([]string, 0, structType.NumFields()),
	}

	for i := 0; i < structType.NumFields(); i++ {
		field := structType.Field(i)
		structInfo.Fields = append(structInfo.Fields, field.Name())
		structInfo.FieldTags = append(structInfo.FieldTags, structType.Tag(i))
	}

	allStructInfos = append(allStructInfos, structInfo)
	return allStructInfos
}

func writeOut(ctx *genall.GenerationContext, root *loader.Package, outBytes []byte) {
	outputFile, err := ctx.Open(root, "zz_gen.column.go")
	if err != nil {
		root.AddError(err)
		return
	}
	defer outputFile.Close()
	n, err := outputFile.Write(outBytes)
	if err != nil {
		root.AddError(err)
		return
	}
	if n < len(outBytes) {
		root.AddError(io.ErrShortWrite)
	}
}
