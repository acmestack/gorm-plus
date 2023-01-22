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

type CommonDao[T any] struct{}

func NewCommonDao[T any]() *CommonDao[T] {
	return &CommonDao[T]{}
}

func (service CommonDao[T]) Db() *gorm.DB {
	return gormDb
}

func (service CommonDao[T]) Save(entity *T) *gorm.DB {
	return Insert[T](entity)
}

func (service CommonDao[T]) SaveBatch(entities []*T) *gorm.DB {
	return InsertBatch[T](entities)
}

func (service CommonDao[T]) SaveBatchSize(entities []*T, batchSize int) *gorm.DB {
	return InsertBatchSize[T](entities, batchSize)
}

func (service CommonDao[T]) RemoveById(id any) *gorm.DB {
	return DeleteById[T](id)
}

func (service CommonDao[T]) RemoveByIds(ids any) *gorm.DB {
	return DeleteByIds[T](ids)
}

func (service CommonDao[T]) Remove(q *Query[T]) *gorm.DB {
	return Delete[T](q)
}

func (service CommonDao[T]) UpdateById(entity *T, id any) *gorm.DB {
	return UpdateById[T](entity, id)
}

func (service CommonDao[T]) Update(q *Query[T]) *gorm.DB {
	return Update[T](q)
}

func (service CommonDao[T]) GetById(id any) (*T, *gorm.DB) {
	return SelectById[T](id)
}

func (service CommonDao[T]) GetOne(q *Query[T]) (*T, *gorm.DB) {
	return SelectOne[T](q)
}

func (service CommonDao[T]) ListAll() ([]*T, *gorm.DB) {
	return SelectList[T](nil)
}

func (service CommonDao[T]) List(q *Query[T]) ([]*T, *gorm.DB) {
	return SelectList[T](q)
}

func (service CommonDao[T]) ListByIds(ids any) ([]*T, *gorm.DB) {
	return SelectByIds[T](ids)
}

func (service CommonDao[T]) PageAll(page *Page[T]) (*Page[T], *gorm.DB) {
	return SelectPage[T](page, nil)
}

func (service CommonDao[T]) Page(page *Page[T], q *Query[T]) (*Page[T], *gorm.DB) {
	return SelectPage[T](page, q)
}

func (service CommonDao[T]) CountAll() (int64, *gorm.DB) {
	return SelectCount[T](nil)
}

func (service CommonDao[T]) Count(q *Query[T]) (int64, *gorm.DB) {
	return SelectCount[T](q)
}
