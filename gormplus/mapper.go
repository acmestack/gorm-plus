package gormplus

import (
	"gorm.io/gorm"
)

var GormDb *gorm.DB

func Insert[T any](entity *T) *gorm.DB {
	resultDb := GormDb.Create(&entity)
	return resultDb
}

func InsertMigrate[T any](entity *T) (*gorm.DB, error) {
	if err := GormDb.AutoMigrate(new(T)); err != nil {
		return nil, err
	}
	resultDb := GormDb.Create(&entity)
	return resultDb, nil
}

func InsertBatch[T any](entities ...*T) *gorm.DB {
	resultDb := GormDb.Create(&entities)
	return resultDb
}

func InsertBatchMigrate[T any](entities ...*T) (*gorm.DB, error) {
	if err := GormDb.AutoMigrate(new(T)); err != nil {
		return nil, err
	}
	resultDb := GormDb.Create(&entities)
	return resultDb, nil
}

func DeleteById[T any](id any) *gorm.DB {
	resultDb := GormDb.Delete(new(T), id)
	return resultDb
}

func DeleteByIds[T any](ids ...any) *gorm.DB {
	var entities []T
	resultDb := GormDb.Delete(&entities, ids)
	return resultDb
}

func Delete[T any](q *Query[T]) *gorm.DB {
	var entity T
	resultDb := GormDb.Where(q.QueryBuilder.String(), q.QueryArgs...).Delete(&entity)
	return resultDb
}

func UpdateById[T any](id any, entity *T) *gorm.DB {
	var e T
	GormDb.First(&e, id)
	resultDb := GormDb.Model(&e).Updates(entity)
	return resultDb
}

func Update[T any](q *Query[T], entity *T) *gorm.DB {
	resultDb := GormDb.Where(q.QueryBuilder.String(), q.QueryArgs...).Updates(entity)
	return resultDb
}

func SelectById[T any](id any) (*gorm.DB, T) {
	var entity T
	resultDb := GormDb.First(&entity, id)
	return resultDb, entity
}

func SelectByIds[T any](ids ...any) (*gorm.DB, []T) {
	var results []T
	resultDb := GormDb.Find(&results, ids)
	return resultDb, results
}

func SelectOne[T any](q *Query[T]) (*gorm.DB, T) {
	var entity T
	resultDb := GormDb.Select(q.Columns).Where(q.QueryBuilder.String(), q.QueryArgs...).First(&entity)
	return resultDb, entity
}

func SelectList[T any](q *Query[T]) (*gorm.DB, []T) {
	var results []T
	resultDb := GormDb.Select(q.Columns).Where(q.QueryBuilder.String(), q.QueryArgs...).
		Order(q.OrderBuilder.String())

	if q.GroupBuilder.Len() > 0 {
		resultDb.Group(q.GroupBuilder.String())
	}

	if q.HavingBuilder.Len() > 0 {
		resultDb.Having(q.HavingBuilder.String(), q.HavingArgs...)
	}

	resultDb.Find(&results)
	return resultDb, results
}

func SelectCount[T any](q *Query[T]) (*gorm.DB, int64) {
	var count int64
	resultDb := GormDb.Model(new(T)).Where(q.QueryBuilder.String(), q.QueryArgs...).Count(&count)
	return resultDb, count
}
