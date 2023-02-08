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
	"gorm.io/gorm/schema"
	"reflect"
	"strings"
)

type Query[T any] struct {
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
	ColumnNameMap     map[uintptr]string
}

func NewQuery[T any]() (*Query[T], *T) {
	q := &Query[T]{}
	return q, q.buildColumnNameMap()
}

func (q *Query[T]) Eq(column any, val any) *Query[T] {
	q.addCond(column, val, constants.Eq)
	return q
}

func (q *Query[T]) Ne(column any, val any) *Query[T] {
	q.addCond(column, val, constants.Ne)
	return q
}

func (q *Query[T]) Gt(column any, val any) *Query[T] {
	q.addCond(column, val, constants.Gt)
	return q
}

func (q *Query[T]) Ge(column any, val any) *Query[T] {
	q.addCond(column, val, constants.Ge)
	return q
}

func (q *Query[T]) Lt(column any, val any) *Query[T] {
	q.addCond(column, val, constants.Lt)
	return q
}

func (q *Query[T]) Le(column any, val any) *Query[T] {
	q.addCond(column, val, constants.Le)
	return q
}

func (q *Query[T]) Like(column any, val any) *Query[T] {
	s := val.(string)
	q.addCond(column, "%"+s+"%", constants.Like)
	return q
}

func (q *Query[T]) NotLike(column any, val any) *Query[T] {
	s := val.(string)
	q.addCond(column, "%"+s+"%", constants.Not+" "+constants.Like)
	return q
}

func (q *Query[T]) LikeLeft(column any, val any) *Query[T] {
	s := val.(string)
	q.addCond(column, "%"+s, constants.Like)
	return q
}

func (q *Query[T]) LikeRight(column any, val any) *Query[T] {
	s := val.(string)
	q.addCond(column, s+"%", constants.Like)
	return q
}

func (q *Query[T]) IsNull(column string) *Query[T] {
	q.buildAndIfNeed()
	cond := fmt.Sprintf("%s is null", column)
	q.QueryBuilder.WriteString(cond)
	return q
}

func (q *Query[T]) IsNotNull(column string) *Query[T] {
	q.buildAndIfNeed()
	cond := fmt.Sprintf("%s is not null", column)
	q.QueryBuilder.WriteString(cond)
	return q
}

func (q *Query[T]) In(column string, val any) *Query[T] {
	q.addCond(column, val, constants.In)
	return q
}

func (q *Query[T]) NotIn(column any, val any) *Query[T] {
	q.addCond(column, val, constants.Not+" "+constants.In)
	return q
}

func (q *Query[T]) Between(column any, start, end any) *Query[T] {
	q.buildAndIfNeed()
	cond := fmt.Sprintf("%s %s ? and ? ", column, constants.Between)
	q.QueryBuilder.WriteString(cond)
	q.QueryArgs = append(q.QueryArgs, start, end)
	return q
}

func (q *Query[T]) NotBetween(column any, start, end any) *Query[T] {
	q.buildAndIfNeed()
	cond := fmt.Sprintf("%s %s %s ? and ? ", column, constants.Not, constants.Between)
	q.QueryBuilder.WriteString(cond)
	q.QueryArgs = append(q.QueryArgs, start, end)
	return q
}

func (q *Query[T]) Distinct(column ...string) *Query[T] {
	q.DistinctColumns = column
	return q
}

func (q *Query[T]) And() *Query[T] {
	q.QueryBuilder.WriteString(constants.And)
	q.QueryBuilder.WriteString(" ")
	q.LastCond = constants.And
	return q
}

func (q *Query[T]) AndBracket(bracketQuery *Query[T]) *Query[T] {
	q.AndBracketBuilder.WriteString(constants.And + " " + constants.LeftBracket + bracketQuery.QueryBuilder.String() + constants.RightBracket + " ")
	q.AndBracketArgs = append(q.AndBracketArgs, bracketQuery.QueryArgs...)
	return q
}

func (q *Query[T]) Or() *Query[T] {
	q.QueryBuilder.WriteString(constants.Or)
	q.QueryBuilder.WriteString(" ")
	q.LastCond = constants.Or
	return q
}

func (q *Query[T]) OrBracket(bracketQuery *Query[T]) *Query[T] {
	q.OrBracketBuilder.WriteString(constants.Or + " " + constants.LeftBracket + bracketQuery.QueryBuilder.String() + constants.RightBracket + " ")
	q.OrBracketArgs = append(q.OrBracketArgs, bracketQuery.QueryArgs...)
	return q
}

func (q *Query[T]) Select(columns ...any) *Query[T] {
	for _, v := range columns {
		valueOf := reflect.ValueOf(v)
		switch valueOf.Kind() {
		case reflect.String:
			columnStr := v.(string)
			q.SelectColumns = append(q.SelectColumns, columnStr)
		case reflect.Pointer:
			columnName := q.ColumnNameMap[valueOf.Pointer()]
			q.SelectColumns = append(q.SelectColumns, columnName)
		}
	}
	return q
}

func (q *Query[T]) OrderByDesc(columns ...string) *Query[T] {
	q.buildOrder(constants.Desc, columns...)
	return q
}

func (q *Query[T]) OrderByAsc(columns ...string) *Query[T] {
	q.buildOrder(constants.Asc, columns...)
	return q
}

func (q *Query[T]) Group(columns ...any) *Query[T] {
	for _, v := range columns {
		columnName := q.ColumnNameMap[reflect.ValueOf(v).Pointer()]
		if q.GroupBuilder.Len() > 0 {
			q.GroupBuilder.WriteString(constants.Comma)
		}
		q.GroupBuilder.WriteString(columnName)
	}
	return q
}

func (q *Query[T]) Having(having string, args ...any) *Query[T] {
	q.HavingBuilder.WriteString(having)
	q.HavingArgs = append(q.HavingArgs, args)
	return q
}

func (q *Query[T]) Set(column any, val any) *Query[T] {
	columnName := q.ColumnNameMap[reflect.ValueOf(column).Pointer()]
	if q.UpdateMap == nil {
		q.UpdateMap = make(map[string]any)
	}
	q.UpdateMap[columnName] = val
	return q
}

func (q *Query[T]) addCond(column any, val any, condType string) {
	columnName := q.ColumnNameMap[reflect.ValueOf(column).Pointer()]
	q.buildAndIfNeed()
	cond := fmt.Sprintf("%s %s ?", columnName, condType)
	q.QueryBuilder.WriteString(cond)
	q.QueryBuilder.WriteString(" ")
	q.LastCond = ""
	q.QueryArgs = append(q.QueryArgs, val)
}

func (q *Query[T]) buildAndIfNeed() {
	if q.LastCond != constants.And && q.LastCond != constants.Or && q.QueryBuilder.Len() > 0 {
		q.QueryBuilder.WriteString(constants.And)
		q.QueryBuilder.WriteString(" ")
	}
}

func (q *Query[T]) buildOrder(orderType string, columns ...string) {
	for _, v := range columns {
		if q.OrderBuilder.Len() > 0 {
			q.OrderBuilder.WriteString(constants.Comma)
		}
		q.OrderBuilder.WriteString(v)
		q.OrderBuilder.WriteString(" ")
		q.OrderBuilder.WriteString(orderType)
	}
}

func (q *Query[T]) buildColumnNameMap() *T {
	q.ColumnNameMap = make(map[uintptr]string)
	model := new(T)
	valueOf := reflect.ValueOf(model)
	typeOf := reflect.TypeOf(model)
	for i := 0; i < valueOf.Elem().NumField(); i++ {
		pointer := valueOf.Elem().Field(i).Addr().Pointer()
		field := typeOf.Elem().Field(i)
		tagSetting := schema.ParseTagSetting(field.Tag.Get("gorm"), ";")
		name, ok := tagSetting["COLUMN"]
		if ok {
			q.ColumnNameMap[pointer] = name
		} else {
			namingStrategy := schema.NamingStrategy{}
			name = namingStrategy.ColumnName("", field.Name)
			q.ColumnNameMap[pointer] = name
		}
	}
	return model
}
