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
	return globalDb
}

func (service CommonDao[T]) Save(entity *T, dbs ...*gorm.DB) *gorm.DB {
	return Insert[T](entity, dbs...)
}

func (service CommonDao[T]) SaveBatch(entities []*T, dbs ...*gorm.DB) *gorm.DB {
	return InsertBatch[T](entities, dbs...)
}

func (service CommonDao[T]) SaveBatchSize(entities []*T, batchSize int, dbs ...*gorm.DB) *gorm.DB {
	return InsertBatchSize[T](entities, batchSize, dbs...)
}

func (service CommonDao[T]) RemoveById(id any, dbs ...*gorm.DB) *gorm.DB {
	return DeleteById[T](id, dbs...)
}

func (service CommonDao[T]) RemoveByIds(ids any, dbs ...*gorm.DB) *gorm.DB {
	return DeleteByIds[T](ids, dbs...)
}

func (service CommonDao[T]) Remove(q *Query[T], dbs ...*gorm.DB) *gorm.DB {
	return Delete[T](q, dbs...)
}

func (service CommonDao[T]) UpdateById(entity *T, dbs ...*gorm.DB) *gorm.DB {
	return UpdateById[T](entity, dbs...)
}

func (service CommonDao[T]) Update(q *Query[T], dbs ...*gorm.DB) *gorm.DB {
	return Update[T](q, dbs...)
}

func (service CommonDao[T]) GetById(id any, dbs ...*gorm.DB) (*T, *gorm.DB) {
	return SelectById[T](id, dbs...)
}

func (service CommonDao[T]) GetOne(q *Query[T], dbs ...*gorm.DB) (*T, *gorm.DB) {
	return SelectOne[T](q, dbs...)
}

func (service CommonDao[T]) ListAll(dbs ...*gorm.DB) ([]*T, *gorm.DB) {
	return SelectList[T](nil, dbs...)
}

func (service CommonDao[T]) List(q *Query[T], dbs ...*gorm.DB) ([]*T, *gorm.DB) {
	return SelectList[T](q, dbs...)
}

func (service CommonDao[T]) ListByIds(ids any, dbs ...*gorm.DB) ([]*T, *gorm.DB) {
	return SelectByIds[T](ids, dbs...)
}

func (service CommonDao[T]) PageAll(page *Page[T], dbs ...*gorm.DB) (*Page[T], *gorm.DB) {
	return SelectPage[T](page, nil, dbs...)
}

func (service CommonDao[T]) Page(page *Page[T], q *Query[T], dbs ...*gorm.DB) (*Page[T], *gorm.DB) {
	return SelectPage[T](page, q, dbs...)
}

func (service CommonDao[T]) CountAll(dbs ...*gorm.DB) (int64, *gorm.DB) {
	return SelectCount[T](nil, dbs...)
}

func (service CommonDao[T]) Count(q *Query[T], dbs ...*gorm.DB) (int64, *gorm.DB) {
	return SelectCount[T](q, dbs...)
}
