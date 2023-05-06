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
	"reflect"
	"strings"
	"sync"

	"github.com/acmestack/gorm-plus/constants"
	"gorm.io/gorm/schema"
)

var columnNameMapCache sync.Map
var modelInstanceCache sync.Map

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
	ConditionMap      map[any]any
}

func NewQuery[T any]() (*Query[T], *T) {
	q := &Query[T]{}
	return q, q.buildColumnNameMap()
}

func NewQueryMap[T any]() (*Query[T], *T) {
	q := &Query[T]{}
	q.ConditionMap = make(map[any]any)
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
	s := fmt.Sprintf("%v", val)
	q.addCond(column, "%"+s+"%", constants.Like)
	return q
}

func (q *Query[T]) NotLike(column any, val any) *Query[T] {
	s := fmt.Sprintf("%v", val)
	q.addCond(column, "%"+s+"%", constants.Not+" "+constants.Like)
	return q
}

func (q *Query[T]) LikeLeft(column any, val any) *Query[T] {
	s := fmt.Sprintf("%v", val)
	q.addCond(column, "%"+s, constants.Like)
	return q
}

func (q *Query[T]) LikeRight(column any, val any) *Query[T] {
	s := fmt.Sprintf("%v", val)
	q.addCond(column, s+"%", constants.Like)
	return q
}

func (q *Query[T]) IsNull(column any) *Query[T] {
	columnName := q.getColumnName(column)
	q.buildAndIfNeed()
	cond := fmt.Sprintf("%s is null", columnName)
	q.QueryBuilder.WriteString(cond)
	return q
}

func (q *Query[T]) IsNotNull(column any) *Query[T] {
	columnName := q.getColumnName(column)
	q.buildAndIfNeed()
	cond := fmt.Sprintf("%s is not null", columnName)
	q.QueryBuilder.WriteString(cond)
	return q
}

func (q *Query[T]) In(column any, val any) *Query[T] {
	q.addCond(column, val, constants.In)
	return q
}

func (q *Query[T]) NotIn(column any, val any) *Query[T] {
	q.addCond(column, val, constants.Not+" "+constants.In)
	return q
}

func (q *Query[T]) Between(column any, start, end any) *Query[T] {
	columnName := q.getColumnName(column)
	q.buildAndIfNeed()
	cond := fmt.Sprintf("%s %s ? and ? ", columnName, constants.Between)
	q.QueryBuilder.WriteString(cond)
	q.QueryArgs = append(q.QueryArgs, start, end)
	return q
}

func (q *Query[T]) NotBetween(column any, start, end any) *Query[T] {
	columnName := q.getColumnName(column)
	q.buildAndIfNeed()
	cond := fmt.Sprintf("%s %s %s ? and ? ", columnName, constants.Not, constants.Between)
	q.QueryBuilder.WriteString(cond)
	q.QueryArgs = append(q.QueryArgs, start, end)
	return q
}

func (q *Query[T]) Distinct(columns ...any) *Query[T] {
	for _, v := range columns {
		columnName := q.ColumnNameMap[reflect.ValueOf(v).Pointer()]
		q.DistinctColumns = append(q.DistinctColumns, columnName)
	}
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
		columnName := q.getColumnName(v)
		q.SelectColumns = append(q.SelectColumns, columnName)
	}
	return q
}

func (q *Query[T]) OrderByDesc(columns ...any) *Query[T] {
	var columnNames []string
	for _, v := range columns {
		columnName := q.getColumnName(v)
		columnNames = append(columnNames, columnName)
	}
	q.buildOrder(constants.Desc, columnNames...)
	return q
}

func (q *Query[T]) OrderByAsc(columns ...any) *Query[T] {
	var columnNames []string
	for _, v := range columns {
		columnName := q.getColumnName(v)
		columnNames = append(columnNames, columnName)
	}
	q.buildOrder(constants.Asc, columnNames...)
	return q
}

func (q *Query[T]) Group(columns ...any) *Query[T] {
	for _, v := range columns {
		columnName := q.getColumnName(v)
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
	columnName := q.getColumnName(column)
	if q.UpdateMap == nil {
		q.UpdateMap = make(map[string]any)
	}
	q.UpdateMap[columnName] = val
	return q
}

func (q *Query[T]) addCond(column any, val any, condType string) {
	columnName := q.getColumnName(column)
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
	// first try to load from cache
	modelTypeStr := reflect.TypeOf((*T)(nil)).Elem().String()
	if model, ok := modelInstanceCache.Load(modelTypeStr); ok {
		if cachedColumnNameMap, ok := columnNameMapCache.Load(modelTypeStr); ok {
			q.ColumnNameMap = cachedColumnNameMap.(map[uintptr]string)
			return model.(*T)
		}
	}
	q.ColumnNameMap = make(map[uintptr]string)
	model := new(T)
	valueOf := reflect.ValueOf(model)
	typeOf := reflect.TypeOf(model)

	for i := 0; i < valueOf.Elem().NumField(); i++ {
		field := typeOf.Elem().Field(i)
		if field.Anonymous {
			modelType := field.Type
			if modelType.Kind() == reflect.Ptr {
				modelType = modelType.Elem()
			}
			for j := 0; j < modelType.NumField(); j++ {
				pointer := valueOf.Elem().FieldByName(modelType.Field(j).Name).Addr().Pointer()
				name := parseColumnName(modelType.Field(j))
				q.ColumnNameMap[pointer] = name
			}
		} else {
			pointer := valueOf.Elem().Field(i).Addr().Pointer()
			name := parseColumnName(field)
			q.ColumnNameMap[pointer] = name
		}
	}

	// store to cache
	modelInstanceCache.Store(modelTypeStr, model)
	columnNameMapCache.Store(modelTypeStr, q.ColumnNameMap)

	return model
}

func parseColumnName(field reflect.StructField) string {
	tagSetting := schema.ParseTagSetting(field.Tag.Get("gorm"), ";")
	name, ok := tagSetting["COLUMN"]
	if ok {
		return name
	}
	namingStrategy := schema.NamingStrategy{}
	return namingStrategy.ColumnName("", field.Name)
}

func (q *Query[T]) getColumnName(v any) string {
	var columnName string
	valueOf := reflect.ValueOf(v)
	switch valueOf.Kind() {
	case reflect.String:
		columnName = v.(string)
	case reflect.Pointer:
		columnName = q.ColumnNameMap[valueOf.Pointer()]
	}
	return columnName
}
