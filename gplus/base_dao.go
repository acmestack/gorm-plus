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
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils"
	"reflect"
)

var gormDb *gorm.DB
var defaultBatchSize = 1000

func Init(db *gorm.DB) {
	gormDb = db
}

type Page[T any] struct {
	Current int
	Size    int
	Total   int64
	Records []*T
}

func NewPage[T any](current, size int) *Page[T] {
	return &Page[T]{Current: current, Size: size}
}

func Insert[T any](entity *T) *gorm.DB {
	resultDb := gormDb.Create(entity)
	return resultDb
}

func InsertBatch[T any](entities []*T) *gorm.DB {
	if len(entities) == 0 {
		return gormDb
	}
	resultDb := gormDb.CreateInBatches(entities, defaultBatchSize)
	return resultDb
}

func InsertBatchSize[T any](entities []*T, batchSize int) *gorm.DB {
	if len(entities) == 0 {
		return gormDb
	}
	if batchSize <= 0 {
		batchSize = defaultBatchSize
	}
	resultDb := gormDb.CreateInBatches(entities, batchSize)
	return resultDb
}

func DeleteById[T any](id any) *gorm.DB {
	var entity T
	resultDb := gormDb.Where(getPkColumnName[T](), id).Delete(&entity)
	return resultDb
}

func DeleteByIds[T any](ids any) *gorm.DB {
	q, _ := NewQuery[T]()
	q.In(getPkColumnName[T](), ids)
	resultDb := Delete[T](q)
	return resultDb
}

func Delete[T any](q *Query[T]) *gorm.DB {
	var entity T
	resultDb := gormDb.Where(q.QueryBuilder.String(), q.QueryArgs...).Delete(&entity)
	return resultDb
}

func UpdateById[T any](entity *T) *gorm.DB {
	resultDb := gormDb.Model(entity).Updates(entity)
	return resultDb
}

func Update[T any](q *Query[T]) *gorm.DB {
	resultDb := gormDb.Model(new(T)).Where(q.QueryBuilder.String(), q.QueryArgs...).Updates(&q.UpdateMap)
	return resultDb
}

func SelectById[T any](id any) (*T, *gorm.DB) {
	q, _ := NewQuery[T]()
	q.Eq(getPkColumnName[T](), id)
	var entity T
	resultDb := buildCondition(q)
	return &entity, resultDb.Limit(1).Find(&entity)
}

func SelectByIds[T any](ids any) ([]*T, *gorm.DB) {
	q, _ := NewQuery[T]()
	q.In(getPkColumnName[T](), ids)
	return SelectList[T](q)
}

func SelectOne[T any](q *Query[T]) (*T, *gorm.DB) {
	var entity T
	resultDb := buildCondition(q)
	return &entity, resultDb.Limit(1).Find(&entity)
}

func SelectList[T any](q *Query[T]) ([]*T, *gorm.DB) {
	resultDb := buildCondition(q)
	var results []*T
	resultDb.Find(&results)
	return results, resultDb
}

func SelectListModel[T any, R any](q *Query[T]) ([]*R, *gorm.DB) {
	resultDb := buildCondition(q)
	var results []*R
	resultDb.Scan(&results)
	return results, resultDb
}

func SelectListByMap[T any](q *Query[T]) ([]*T, *gorm.DB) {
	resultDb := buildCondition(q)
	var results []*T
	resultDb.Find(&results)
	return results, resultDb
}

func SelectListMaps[T any](q *Query[T]) ([]map[string]any, *gorm.DB) {
	resultDb := buildCondition(q)
	var results []map[string]any
	resultDb.Find(&results)
	return results, resultDb
}

func SelectPage[T any](page *Page[T], q *Query[T]) (*Page[T], *gorm.DB) {
	total, countDb := SelectCount[T](q)
	if countDb.Error != nil {
		return page, countDb
	}
	page.Total = total
	resultDb := buildCondition(q)
	var results []*T
	resultDb.Scopes(paginate(page)).Find(&results)
	page.Records = results
	return page, resultDb
}

func SelectPageModel[T any, R any](page *Page[R], q *Query[T]) (*Page[R], *gorm.DB) {
	total, countDb := SelectCount[T](q)
	if countDb.Error != nil {
		return page, countDb
	}
	page.Total = total
	resultDb := buildCondition(q)
	var results []*R
	resultDb.Scopes(paginate(page)).Scan(&results)
	page.Records = results
	return page, resultDb
}

func SelectCount[T any](q *Query[T]) (int64, *gorm.DB) {
	var count int64
	resultDb := buildCondition(q)
	resultDb.Count(&count)
	return count, resultDb
}

func paginate[T any](p *Page[T]) func(db *gorm.DB) *gorm.DB {
	page := p.Current
	pageSize := p.Size
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func buildCondition[T any](q *Query[T]) *gorm.DB {
	resultDb := gormDb.Model(new(T))
	if q != nil {
		if len(q.DistinctColumns) > 0 {
			resultDb.Distinct(q.DistinctColumns)
		}

		if len(q.SelectColumns) > 0 {
			resultDb.Select(q.SelectColumns)
		}

		if q.QueryBuilder.Len() > 0 {

			if q.AndBracketBuilder.Len() > 0 {
				q.QueryArgs = append(q.QueryArgs, q.AndBracketArgs...)
				q.QueryBuilder.WriteString(q.AndBracketBuilder.String())
			}

			if q.OrBracketBuilder.Len() > 0 {
				q.QueryArgs = append(q.QueryArgs, q.OrBracketArgs...)
				q.QueryBuilder.WriteString(q.OrBracketBuilder.String())
			}

			resultDb.Where(q.QueryBuilder.String(), q.QueryArgs...)
		}

		if len(q.ConditionMap) > 0 {
			var condMap = make(map[string]any)
			for k, v := range q.ConditionMap {
				columnName := q.getColumnName(k)
				condMap[columnName] = v
			}
			resultDb.Where(condMap)
		}

		if q.OrderBuilder.Len() > 0 {
			resultDb.Order(q.OrderBuilder.String())
		}

		if q.GroupBuilder.Len() > 0 {
			resultDb.Group(q.GroupBuilder.String())
		}

		if q.HavingBuilder.Len() > 0 {
			resultDb.Having(q.HavingBuilder.String(), q.HavingArgs...)
		}
	}
	return resultDb
}

func getPkColumnName[T any]() string {
	var entity T
	entityType := reflect.TypeOf(entity)
	numField := entityType.NumField()
	var columnName string
	for i := 0; i < numField; i++ {
		field := entityType.Field(i)
		tagSetting := schema.ParseTagSetting(field.Tag.Get("gorm"), ";")
		isPrimaryKey := utils.CheckTruth(tagSetting["PRIMARYKEY"], tagSetting["PRIMARY_KEY"])
		if isPrimaryKey {
			name, ok := tagSetting["COLUMN"]
			if !ok {
				namingStrategy := schema.NamingStrategy{}
				name = namingStrategy.ColumnName("", field.Name)
			}
			columnName = name
			break
		}
	}
	if columnName == "" {
		return constants.DefaultPrimaryName
	}
	return columnName
}
