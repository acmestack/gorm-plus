package gormplus

import (
	"gorm.io/gorm"
)

var gormDb *gorm.DB
var defaultBatchSize = 1000

func Init(db *gorm.DB) {
	gormDb = db
}

type Page struct {
	Page     int
	PageSize int
}

func Insert[T any](entity *T) *gorm.DB {
	resultDb := gormDb.Create(&entity)
	return resultDb
}

func InsertMigrate[T any](entity *T) (*gorm.DB, error) {
	if err := gormDb.AutoMigrate(new(T)); err != nil {
		return nil, err
	}
	resultDb := gormDb.Create(&entity)
	return resultDb, nil
}

func InsertBatch[T any](entities ...*T) *gorm.DB {
	resultDb := gormDb.CreateInBatches(&entities, defaultBatchSize)
	return resultDb
}

func InsertBatchSize[T any](batchSize int, entities ...*T) *gorm.DB {
	if batchSize <= 0 {
		batchSize = defaultBatchSize
	}
	resultDb := gormDb.CreateInBatches(&entities, batchSize)
	return resultDb
}

func InsertBatchMigrate[T any](entities ...*T) (*gorm.DB, error) {
	if err := gormDb.AutoMigrate(new(T)); err != nil {
		return nil, err
	}
	resultDb := gormDb.Create(&entities)
	return resultDb, nil
}

func DeleteById[T any](id any) *gorm.DB {
	resultDb := gormDb.Delete(new(T), id)
	return resultDb
}

func DeleteByIds[T any](ids ...any) *gorm.DB {
	var entities []T
	resultDb := gormDb.Delete(&entities, ids)
	return resultDb
}

func Delete[T any](q *Query[T]) *gorm.DB {
	var entity T
	resultDb := gormDb.Where(q.QueryBuilder.String(), q.QueryArgs...).Delete(&entity)
	return resultDb
}

func UpdateById[T any](id any, entity *T) *gorm.DB {
	var e T
	gormDb.First(&e, id)
	resultDb := gormDb.Model(&e).Updates(entity)
	return resultDb
}

func Update[T any](q *Query[T], entity *T) *gorm.DB {
	resultDb := gormDb.Where(q.QueryBuilder.String(), q.QueryArgs...).Updates(entity)
	return resultDb
}

func SelectById[T any](id any) (*gorm.DB, T) {
	var entity T
	resultDb := gormDb.Limit(1).Find(&entity, id)
	return resultDb, entity
}

func SelectByIds[T any](ids ...any) (*gorm.DB, []T) {
	var results []T
	resultDb := gormDb.Find(&results, ids)
	return resultDb, results
}

func SelectOne[T any](q *Query[T]) (*gorm.DB, T) {
	var entity T
	resultDb := gormDb.Select(q.SelectColumns).Where(q.QueryBuilder.String(), q.QueryArgs...).Limit(1).Find(&entity)
	return resultDb, entity
}

func SelectList[T any](q *Query[T]) (*gorm.DB, []T) {
	resultDb := buildCondition(q)
	var results []T
	resultDb.Find(&results)
	return resultDb, results
}

func SelectPage[T any](page Page, q *Query[T]) (*gorm.DB, []T) {
	resultDb := buildCondition(q)
	var results []T
	resultDb.Scopes(paginate(page)).Find(&results)
	return resultDb, results
}

func SelectCount[T any](q *Query[T]) (*gorm.DB, int64) {
	var count int64
	resultDb := gormDb.Model(new(T)).Where(q.QueryBuilder.String(), q.QueryArgs...).Count(&count)
	return resultDb, count
}

func paginate(p Page) func(db *gorm.DB) *gorm.DB {
	page := p.Page
	pageSize := p.PageSize
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
			resultDb.Where(q.QueryBuilder.String(), q.QueryArgs...)
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
