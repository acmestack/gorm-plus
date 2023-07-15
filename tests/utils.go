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

package tests

import (
	"database/sql/driver"
	"fmt"
	"go/ast"
	"gorm.io/gorm"
	"reflect"
	"strings"
	"testing"
	"time"

	"gorm.io/gorm/utils"
)

func AssertObjEqual(t *testing.T, r, e interface{}, names ...string) {
	for _, name := range names {
		got := reflect.Indirect(reflect.ValueOf(r)).FieldByName(name).Interface()
		expect := reflect.Indirect(reflect.ValueOf(e)).FieldByName(name).Interface()
		t.Run(name, func(t *testing.T) {
			AssertEqual(t, got, expect)
		})
	}
}

func AssertEqual(t *testing.T, got, expect interface{}) {
	if !reflect.DeepEqual(got, expect) {
		isEqual := func() {
			if curTime, ok := got.(time.Time); ok {
				format := "2006-01-02T15:04:05Z07:00"

				if curTime.Round(time.Second).UTC().Format(format) != expect.(time.Time).Round(time.Second).UTC().Format(format) && curTime.Truncate(time.Second).UTC().Format(format) != expect.(time.Time).Truncate(time.Second).UTC().Format(format) {
					t.Errorf("%v: expect: %v, got %v after time round", utils.FileWithLineNum(), expect.(time.Time), curTime)
				}
			} else if fmt.Sprint(got) != fmt.Sprint(expect) {
				t.Errorf("%v: expect: %#v, got %#v", utils.FileWithLineNum(), expect, got)
			}
		}

		if fmt.Sprint(got) == fmt.Sprint(expect) {
			return
		}

		if reflect.Indirect(reflect.ValueOf(got)).IsValid() != reflect.Indirect(reflect.ValueOf(expect)).IsValid() {
			t.Errorf("%v: expect: %+v, got %+v", utils.FileWithLineNum(), expect, got)
			return
		}

		if valuer, ok := got.(driver.Valuer); ok {
			got, _ = valuer.Value()
		}

		if valuer, ok := expect.(driver.Valuer); ok {
			expect, _ = valuer.Value()
		}

		if got != nil {
			got = reflect.Indirect(reflect.ValueOf(got)).Interface()
		}

		if expect != nil {
			expect = reflect.Indirect(reflect.ValueOf(expect)).Interface()
		}

		if reflect.ValueOf(got).IsValid() != reflect.ValueOf(expect).IsValid() {
			t.Errorf("%v: expect: %+v, got %+v", utils.FileWithLineNum(), expect, got)
			return
		}

		if reflect.ValueOf(got).Kind() == reflect.Slice {
			if reflect.ValueOf(expect).Kind() == reflect.Slice {
				if reflect.ValueOf(got).Len() == reflect.ValueOf(expect).Len() {
					for i := 0; i < reflect.ValueOf(got).Len(); i++ {
						name := fmt.Sprintf(reflect.ValueOf(got).Type().Name()+" #%v", i)
						t.Run(name, func(t *testing.T) {
							AssertEqual(t, reflect.ValueOf(got).Index(i).Interface(), reflect.ValueOf(expect).Index(i).Interface())
						})
					}
				} else {
					name := reflect.ValueOf(got).Type().Elem().Name()
					t.Errorf("%v expects length: %v, got %v (expects: %+v, got %+v)", name, reflect.ValueOf(expect).Len(), reflect.ValueOf(got).Len(), expect, got)
				}
				return
			}
		}

		if reflect.ValueOf(got).Kind() == reflect.Struct {
			if reflect.ValueOf(expect).Kind() == reflect.Struct {
				if reflect.ValueOf(got).NumField() == reflect.ValueOf(expect).NumField() {
					exported := false
					for i := 0; i < reflect.ValueOf(got).NumField(); i++ {
						if fieldStruct := reflect.ValueOf(got).Type().Field(i); ast.IsExported(fieldStruct.Name) {
							exported = true
							field := reflect.ValueOf(got).Field(i)
							t.Run(fieldStruct.Name, func(t *testing.T) {
								AssertEqual(t, field.Interface(), reflect.ValueOf(expect).Field(i).Interface())
							})
						}
					}

					if exported {
						return
					}
				}
			}
		}

		if reflect.ValueOf(got).Type().ConvertibleTo(reflect.ValueOf(expect).Type()) {
			got = reflect.ValueOf(got).Convert(reflect.ValueOf(expect).Type()).Interface()
			isEqual()
		} else if reflect.ValueOf(expect).Type().ConvertibleTo(reflect.ValueOf(got).Type()) {
			expect = reflect.ValueOf(got).Convert(reflect.ValueOf(got).Type()).Interface()
			isEqual()
		} else {
			t.Errorf("%v: expect: %+v, got %+v", utils.FileWithLineNum(), expect, got)
			return
		}
	}
}

func Now() *time.Time {
	now := time.Now()
	return &now
}

func buildSql(db *gorm.DB) string {
	sql := db.Statement.SQL.String()
	for _, value := range db.Statement.Vars {
		sql = strings.Replace(sql, "?", convert(value), 1)
	}
	return sql
}

func convert(value any) string {
	columnType := reflect.TypeOf(value)
	switch columnType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%v", value)
	}
	return fmt.Sprintf("'%v'", value)
}
