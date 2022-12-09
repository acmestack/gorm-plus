package gplus

import (
	"gorm.io/gorm"
)

type CommonDao[T any] struct {
}

func (service CommonDao[T]) Save(entity *T) *gorm.DB {
	return Insert[T](entity)
}

func (service CommonDao[T]) SaveBatch(entities any) *gorm.DB {
	return InsertBatch[T](entities)
}

func (service CommonDao[T]) SaveBatchSize(entities any, batchSize int) *gorm.DB {
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

func (service CommonDao[T]) UpdateById(entity *T) *gorm.DB {
	return UpdateById[T](entity)
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
