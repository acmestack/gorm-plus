package gormplus

import (
	"gorm.io/gorm"
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
	resultDb := gormDb.Create(&entity)
	return resultDb
}

func InsertBatch[T any](entities any) *gorm.DB {
	resultDb := gormDb.CreateInBatches(entities, defaultBatchSize)
	return resultDb
}

func InsertBatchSize[T any](entities any, batchSize int) *gorm.DB {
	if batchSize <= 0 {
		batchSize = defaultBatchSize
	}
	resultDb := gormDb.CreateInBatches(entities, batchSize)
	return resultDb
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

func UpdateById[T any](entity *T) *gorm.DB {
	resultDb := gormDb.Model(&entity).Updates(&entity)
	return resultDb
}

func Update[T any](q *Query[T]) *gorm.DB {
	resultDb := gormDb.Model(new(T)).Where(q.QueryBuilder.String(), q.QueryArgs...).Updates(&q.UpdateMap)
	return resultDb
}

func SelectById[T any](id any) (*T, *gorm.DB) {
	var entity *T
	resultDb := gormDb.Take(&entity, id)
	if resultDb.RowsAffected == 0 {
		return nil, resultDb
	}
	return entity, resultDb
}

func SelectByIds[T any](ids any) ([]*T, *gorm.DB) {
	var results []*T
	resultDb := gormDb.Find(&results, ids)
	return results, resultDb
}

func SelectOne[T any](q *Query[T]) (*T, *gorm.DB) {
	var entity *T
	resultDb := buildCondition(q)
	resultDb.Take(&entity)
	if resultDb.RowsAffected == 0 {
		return nil, resultDb
	}
	return entity, resultDb
}

func SelectList[T any](q *Query[T]) ([]*T, *gorm.DB) {
	resultDb := buildCondition(q)
	var results []*T
	resultDb.Find(&results)
	return results, resultDb
}

func SelectModelList[T any, R any](q *Query[T]) ([]*R, *gorm.DB) {
	resultDb := buildCondition(q)
	var results []*R
	resultDb.Scan(&results)
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

func SelectModelPage[T any, R any](page *Page[R], q *Query[T]) (*Page[R], *gorm.DB) {
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
