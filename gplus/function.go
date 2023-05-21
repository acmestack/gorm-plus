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

package gplus

import (
	"github.com/acmestack/gorm-plus/constants"
)

type Function struct {
	funStr string
}

func (f *Function) As(asName any) string {
	return f.funStr + " " + constants.As + " " + getColumnName(asName)
}

func (f *Function) Eq(value int64) (string, int64) {
	return buildFunStr(f.funStr, constants.Eq, value)
}

func (f *Function) Ne(value int64) (string, int64) {
	return buildFunStr(f.funStr, constants.Ne, value)
}

func (f *Function) Gt(value int64) (string, int64) {
	return buildFunStr(f.funStr, constants.Gt, value)
}

func (f *Function) Ge(value int64) (string, int64) {
	return buildFunStr(f.funStr, constants.Ge, value)
}

func (f *Function) Lt(value int64) (string, int64) {
	return buildFunStr(f.funStr, constants.Lt, value)
}

func (f *Function) Le(value int64) (string, int64) {
	return buildFunStr(f.funStr, constants.Le, value)
}

func (f *Function) Between(start int64, end int64) (string, int64, int64) {
	return f.funStr + " " + constants.Between + " ? and ?", start, end
}

func (f *Function) NotBetween(start int64, end int64) (string, int64, int64) {
	return f.funStr + " " + constants.Not + " " + constants.Between + " ? and ?", start, end
}

func Sum(columnName any) *Function {
	return &Function{funStr: addBracket(constants.SUM, getColumnName(columnName))}
}

func Avg(columnName any) *Function {
	return &Function{funStr: addBracket(constants.AVG, getColumnName(columnName))}
}

func Max(columnName any) *Function {
	return &Function{funStr: addBracket(constants.MAX, getColumnName(columnName))}
}

func Min(columnName any) *Function {
	return &Function{funStr: addBracket(constants.MIN, getColumnName(columnName))}
}

func Count(columnName any) *Function {
	return &Function{funStr: addBracket(constants.COUNT, getColumnName(columnName))}
}

func addBracket(function string, columnNameStr string) string {
	return function + constants.LeftBracket + columnNameStr + constants.RightBracket
}

func buildFunStr(funcStr string, typeStr string, value int64) (string, int64) {
	return funcStr + " " + typeStr + " ?", value
}
