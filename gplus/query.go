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
	"fmt"
	"github.com/acmestack/gorm-plus/constants"
	"reflect"
	"strings"
)

type QueryCond[T any] struct {
	SelectColumns     []string
	DistinctColumns   []string
	QueryBuilder      strings.Builder
	OrBracketBuilder  strings.Builder
	OrBracketArgs     []any
	AndBracketBuilder strings.Builder
	AndBracketArgs    []any
	QueryArgs         []any
	OrderBuilder      strings.Builder
	GroupBuilder      strings.Builder
	HavingBuilder     strings.Builder
	HavingArgs        []any
	LastCond          string
	UpdateMap         map[string]any
	ConditionMap      map[any]any
}

func NewQuery[T any]() (*QueryCond[T], *T) {
	q := &QueryCond[T]{}
	return q, q.buildColumnNameMap()
}

func NewQueryMap[T any]() (*QueryCond[T], *T) {
	q := &QueryCond[T]{}
	q.ConditionMap = make(map[any]any)
	return q, q.buildColumnNameMap()
}

func (q *QueryCond[T]) Eq(column any, val any) *QueryCond[T] {
	q.addCond(column, val, constants.Eq)
	return q
}

func (q *QueryCond[T]) Ne(column any, val any) *QueryCond[T] {
	q.addCond(column, val, constants.Ne)
	return q
}

func (q *QueryCond[T]) Gt(column any, val any) *QueryCond[T] {
	q.addCond(column, val, constants.Gt)
	return q
}

func (q *QueryCond[T]) Ge(column any, val any) *QueryCond[T] {
	q.addCond(column, val, constants.Ge)
	return q
}

func (q *QueryCond[T]) Lt(column any, val any) *QueryCond[T] {
	q.addCond(column, val, constants.Lt)
	return q
}

func (q *QueryCond[T]) Le(column any, val any) *QueryCond[T] {
	q.addCond(column, val, constants.Le)
	return q
}

func (q *QueryCond[T]) Like(column any, val any) *QueryCond[T] {
	s := fmt.Sprintf("%v", val)
	q.addCond(column, "%"+s+"%", constants.Like)
	return q
}

func (q *QueryCond[T]) NotLike(column any, val any) *QueryCond[T] {
	s := fmt.Sprintf("%v", val)
	q.addCond(column, "%"+s+"%", constants.Not+" "+constants.Like)
	return q
}

func (q *QueryCond[T]) LikeLeft(column any, val any) *QueryCond[T] {
	s := fmt.Sprintf("%v", val)
	q.addCond(column, "%"+s, constants.Like)
	return q
}

func (q *QueryCond[T]) LikeRight(column any, val any) *QueryCond[T] {
	s := fmt.Sprintf("%v", val)
	q.addCond(column, s+"%", constants.Like)
	return q
}

func (q *QueryCond[T]) IsNull(column any) *QueryCond[T] {
	columnName := getColumnName(column)
	q.buildAndIfNeed()
	cond := fmt.Sprintf("%s is null", columnName)
	q.QueryBuilder.WriteString(cond)
	return q
}

func (q *QueryCond[T]) IsNotNull(column any) *QueryCond[T] {
	columnName := getColumnName(column)
	q.buildAndIfNeed()
	cond := fmt.Sprintf("%s is not null", columnName)
	q.QueryBuilder.WriteString(cond)
	return q
}

func (q *QueryCond[T]) In(column any, val any) *QueryCond[T] {
	q.addCond(column, val, constants.In)
	return q
}

func (q *QueryCond[T]) NotIn(column any, val any) *QueryCond[T] {
	q.addCond(column, val, constants.Not+" "+constants.In)
	return q
}

func (q *QueryCond[T]) Between(column any, start, end any) *QueryCond[T] {
	columnName := getColumnName(column)
	q.buildAndIfNeed()
	cond := fmt.Sprintf("%s %s ? and ? ", columnName, constants.Between)
	q.QueryBuilder.WriteString(cond)
	q.QueryArgs = append(q.QueryArgs, start, end)
	return q
}

func (q *QueryCond[T]) NotBetween(column any, start, end any) *QueryCond[T] {
	columnName := getColumnName(column)
	q.buildAndIfNeed()
	cond := fmt.Sprintf("%s %s %s ? and ? ", columnName, constants.Not, constants.Between)
	q.QueryBuilder.WriteString(cond)
	q.QueryArgs = append(q.QueryArgs, start, end)
	return q
}

func (q *QueryCond[T]) Distinct(columns ...any) *QueryCond[T] {
	for _, v := range columns {
		if columnName, ok := columnNameCache.Load(reflect.ValueOf(v).Pointer()); ok {
			q.DistinctColumns = append(q.DistinctColumns, columnName.(string))
		}
	}
	return q
}

func (q *QueryCond[T]) And() *QueryCond[T] {
	q.QueryBuilder.WriteString(constants.And)
	q.QueryBuilder.WriteString(" ")
	q.LastCond = constants.And
	return q
}

func (q *QueryCond[T]) AndBracket(bracketQuery *QueryCond[T]) *QueryCond[T] {
	q.AndBracketBuilder.WriteString(constants.And + " " + constants.LeftBracket + bracketQuery.QueryBuilder.String() + constants.RightBracket + " ")
	q.AndBracketArgs = append(q.AndBracketArgs, bracketQuery.QueryArgs...)
	return q
}

func (q *QueryCond[T]) Or() *QueryCond[T] {
	q.QueryBuilder.WriteString(constants.Or)
	q.QueryBuilder.WriteString(" ")
	q.LastCond = constants.Or
	return q
}

func (q *QueryCond[T]) OrBracket(bracketQuery *QueryCond[T]) *QueryCond[T] {
	q.OrBracketBuilder.WriteString(constants.Or + " " + constants.LeftBracket + bracketQuery.QueryBuilder.String() + constants.RightBracket + " ")
	q.OrBracketArgs = append(q.OrBracketArgs, bracketQuery.QueryArgs...)
	return q
}

func (q *QueryCond[T]) Select(columns ...any) *QueryCond[T] {
	for _, v := range columns {
		columnName := getColumnName(v)
		q.SelectColumns = append(q.SelectColumns, columnName)
	}
	return q
}

func (q *QueryCond[T]) OrderByDesc(columns ...any) *QueryCond[T] {
	var columnNames []string
	for _, v := range columns {
		columnName := getColumnName(v)
		columnNames = append(columnNames, columnName)
	}
	q.buildOrder(constants.Desc, columnNames...)
	return q
}

func (q *QueryCond[T]) OrderByAsc(columns ...any) *QueryCond[T] {
	var columnNames []string
	for _, v := range columns {
		columnName := getColumnName(v)
		columnNames = append(columnNames, columnName)
	}
	q.buildOrder(constants.Asc, columnNames...)
	return q
}

func (q *QueryCond[T]) Group(columns ...any) *QueryCond[T] {
	for _, v := range columns {
		columnName := getColumnName(v)
		if q.GroupBuilder.Len() > 0 {
			q.GroupBuilder.WriteString(constants.Comma)
		}
		q.GroupBuilder.WriteString(columnName)
	}
	return q
}

func (q *QueryCond[T]) Having(having string, args ...any) *QueryCond[T] {
	q.HavingBuilder.WriteString(having)
	q.HavingArgs = append(q.HavingArgs, args)
	return q
}

func (q *QueryCond[T]) Set(column any, val any) *QueryCond[T] {
	columnName := getColumnName(column)
	if q.UpdateMap == nil {
		q.UpdateMap = make(map[string]any)
	}
	q.UpdateMap[columnName] = val
	return q
}

func (q *QueryCond[T]) addCond(column any, val any, condType string) {
	columnName := getColumnName(column)
	q.buildAndIfNeed()
	cond := fmt.Sprintf("%s %s ?", columnName, condType)
	q.QueryBuilder.WriteString(cond)
	q.QueryBuilder.WriteString(" ")
	q.LastCond = ""
	q.QueryArgs = append(q.QueryArgs, val)
}

func (q *QueryCond[T]) buildAndIfNeed() {
	if q.LastCond != constants.And && q.LastCond != constants.Or && q.QueryBuilder.Len() > 0 {
		q.QueryBuilder.WriteString(constants.And)
		q.QueryBuilder.WriteString(" ")
	}
}

func (q *QueryCond[T]) buildOrder(orderType string, columns ...string) {
	for _, v := range columns {
		if q.OrderBuilder.Len() > 0 {
			q.OrderBuilder.WriteString(constants.Comma)
		}
		q.OrderBuilder.WriteString(v)
		q.OrderBuilder.WriteString(" ")
		q.OrderBuilder.WriteString(orderType)
	}
}

func (q *QueryCond[T]) buildColumnNameMap() *T {
	// first try to load from cache
	modelTypeStr := reflect.TypeOf((*T)(nil)).Elem().String()
	if model, ok := modelInstanceCache.Load(modelTypeStr); ok {
		return model.(*T)
	}
	model := new(T)
	Cache(model, globalDb.Config.NamingStrategy)
	return model
}

func getColumnName(v any) string {
	var columnName string
	valueOf := reflect.ValueOf(v)
	switch valueOf.Kind() {
	case reflect.String:
		columnName = v.(string)
	case reflect.Pointer:
		if name, ok := columnNameCache.Load(valueOf.Pointer()); ok {
			columnName = name.(string)
		}
	}
	return columnName
}
