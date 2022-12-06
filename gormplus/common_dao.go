package gormplus

import "gorm.io/gorm"

type CommonDao[T any] struct {
}

func (service CommonDao[T]) Save(entity *T) *gorm.DB {
	return Insert[T](entity)
}

func (service CommonDao[T]) List(q *Query[T]) ([]*T, *gorm.DB) {
	return SelectList(q)
}
