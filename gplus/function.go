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

import "github.com/acmestack/gorm-plus/constants"

func As(columnName any, asName any) string {
	return getColumnName(columnName) + " " + constants.As + " " + getColumnName(asName)
}

func SumAs(columnName any, asName any) string {
	return buildFunction(constants.SUM, getColumnName(columnName), getColumnName(asName))
}

func AvgAs(columnName any, asName any) string {
	return buildFunction(constants.AVG, getColumnName(columnName), getColumnName(asName))
}

func MaxAs(columnName any, asName any) string {
	return buildFunction(constants.MAX, getColumnName(columnName), getColumnName(asName))
}

func MinAs(columnName any, asName any) string {
	return buildFunction(constants.MIN, getColumnName(columnName), getColumnName(asName))
}

func CountAs(columnName any, asName any) string {
	return buildFunction(constants.COUNT, getColumnName(columnName), getColumnName(asName))
}

func buildFunction(function string, columnNameStr string, asNameStr string) string {
	columnNameStr = function + constants.LeftBracket + columnNameStr + constants.RightBracket +
		" " + constants.As + " " + asNameStr
	return columnNameStr
}
