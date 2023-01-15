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
	"gorm.io/gorm"
)

type CommonDao[T any, K PrimaryKey] struct {
	pkColumn string
}

func NewCommonDao[T any, K PrimaryKey](pk string) *CommonDao[T, K] {
	return &CommonDao[T, K]{pkColumn: pk}
}

func (service CommonDao[T, K]) Db() *gorm.DB {
	return gormDb
}

func (service CommonDao[T, K]) Save(entity *T) *gorm.DB {
	return Insert[T](entity)
}

func (service CommonDao[T, K]) SaveBatch(entities []*T) *gorm.DB {
	return InsertBatch[T](entities)
}

func (service CommonDao[T, K]) SaveBatchSize(entities []*T, batchSize int) *gorm.DB {
	return InsertBatchSize[T](entities, batchSize)
}

func (service CommonDao[T, K]) RemoveById(id K) *gorm.DB {
	return DeleteById[T, K](id, service.pkColumn)
}

func (service CommonDao[T, K]) RemoveByIds(ids []K) *gorm.DB {
	return DeleteByIds[T, K](ids)
}

func (service CommonDao[T, K]) Remove(q *Query[T]) *gorm.DB {
	return Delete[T](q)
}

func (service CommonDao[T, K]) UpdateById(entity *T, id K) *gorm.DB {
	return UpdateById[T, K](entity, id, service.pkColumn)
}

func (service CommonDao[T, K]) Update(q *Query[T]) *gorm.DB {
	return Update[T](q)
}

func (service CommonDao[T, K]) GetById(id K) (*T, *gorm.DB) {
	return SelectById[T, K](id)
}

func (service CommonDao[T, K]) GetOne(q *Query[T]) (*T, *gorm.DB) {
	return SelectOne[T](q)
}

func (service CommonDao[T, K]) ListAll() ([]*T, *gorm.DB) {
	return SelectList[T](nil)
}

func (service CommonDao[T, K]) List(q *Query[T]) ([]*T, *gorm.DB) {
	return SelectList[T](q)
}

func (service CommonDao[T, K]) ListByIds(ids []K) ([]*T, *gorm.DB) {
	return SelectByIds[T, K](ids)
}

func (service CommonDao[T, K]) PageAll(page *Page[T]) (*Page[T], *gorm.DB) {
	return SelectPage[T](page, nil)
}

func (service CommonDao[T, K]) Page(page *Page[T], q *Query[T]) (*Page[T], *gorm.DB) {
	return SelectPage[T](page, q)
}

func (service CommonDao[T, K]) CountAll() (int64, *gorm.DB) {
	return SelectCount[T](nil)
}

func (service CommonDao[T, K]) Count(q *Query[T]) (int64, *gorm.DB) {
	return SelectCount[T](q)
}
