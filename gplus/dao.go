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
	"database/sql"
	"reflect"
	"strings"
	"time"

	"github.com/aixj1984/gorm-plus/constants"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils"
)

var globalDb *gorm.DB
var defaultBatchSize = 1000

func Init(db *gorm.DB) {
	globalDb = db
}

type Page[T any] struct {
	Current    int   `json:"page"`     // 页码
	Size       int   `json:"pageSize"` // 每页大小
	Total      int64 `json:"total"`
	Records    []*T  `json:"list"`
	CurTime    int64 `json:"curTime"` // 当前时间，毫秒
	RecordsMap []T   `json:"listMap"`
}

type Dao[T any] struct{}

func (dao Dao[T]) NewQuery() (*QueryCond[T], *T) {
	return NewQuery[T]()
}

func NewPage[T any](current, size int) *Page[T] {
	return &Page[T]{Current: current, Size: size}
}

// Insert 插入一条记录
func Insert[T any](entity *T, opts ...OptionFunc) *gorm.DB {
	db := getDb(opts...)
	resultDb := db.Create(entity)
	return resultDb
}

// InsertBatch 批量插入多条记录
func InsertBatch[T any](entities []*T, opts ...OptionFunc) *gorm.DB {
	db := getDb(opts...)
	if len(entities) == 0 {
		return db
	}
	resultDb := db.CreateInBatches(entities, defaultBatchSize)
	return resultDb
}

// InsertBatchSize 批量插入多条记录
func InsertBatchSize[T any](entities []*T, batchSize int, opts ...OptionFunc) *gorm.DB {
	db := getDb(opts...)
	if len(entities) == 0 {
		return db
	}
	if batchSize <= 0 {
		batchSize = defaultBatchSize
	}
	resultDb := db.CreateInBatches(entities, batchSize)
	return resultDb
}

// DeleteById 根据 ID 删除记录
func DeleteById[T any](id any, opts ...OptionFunc) *gorm.DB {
	db := getDb(opts...)
	var entity T
	resultDb := db.Where(getPkColumnName[T](), id).Delete(&entity)
	return resultDb
}

// DeleteByIds 根据 ID 批量删除记录
func DeleteByIds[T any](ids any, opts ...OptionFunc) *gorm.DB {
	q, _ := NewQuery[T]()
	q.In(getPkColumnName[T](), ids)
	resultDb := Delete[T](q, opts...)
	return resultDb
}

// Delete 根据条件删除记录
func Delete[T any](q *QueryCond[T], opts ...OptionFunc) *gorm.DB {
	var entity T
	resultDb := buildCondition[T](q, opts...)
	resultDb.Delete(&entity)
	return resultDb
}

// UpdateById 根据 ID 更新,默认零值不更新
func UpdateById[T any](entity *T, opts ...OptionFunc) *gorm.DB {
	db := getDb(opts...)
	resultDb := db.Model(entity).Updates(entity)
	return resultDb
}

// UpdateZeroById 根据 ID 零值更新
func UpdateZeroById[T any](entity *T, opts ...OptionFunc) *gorm.DB {
	db := getDb(opts...)

	// 如果用户没有设置选择更新的字段，默认更新所有的字段，包括零值更新
	updateAllIfNeed(entity, opts, db)

	resultDb := db.Model(entity).Updates(entity)
	return resultDb
}

func updateAllIfNeed(entity any, opts []OptionFunc, db *gorm.DB) {
	option := getOption(opts)
	if len(option.Selects) == 0 {
		columnNameMap := getColumnNameMap(entity)
		var columnNames []string
		for _, columnName := range columnNameMap {
			columnNames = append(columnNames, columnName)
		}
		db.Select(columnNames)
	}
}

// Update 根据 Map 更新
func Update[T any](q *QueryCond[T], opts ...OptionFunc) *gorm.DB {
	resultDb := buildCondition[T](q, opts...)
	resultDb.Updates(&q.updateMap)
	return resultDb
}

// SelectById 根据 ID 查询单条记录
func SelectById[T any](id any, opts ...OptionFunc) (*T, *gorm.DB) {
	q, _ := NewQuery[T]()
	q.Eq(getPkColumnName[T](), id)
	var entity T
	resultDb := buildCondition(q, opts...)
	return &entity, resultDb.Take(&entity)
}

// SelectByIds 根据 ID 查询多条记录
func SelectByIds[T any](ids any, opts ...OptionFunc) ([]*T, *gorm.DB) {
	q, _ := NewQuery[T]()
	q.In(getPkColumnName[T](), ids)
	return SelectList[T](q, opts...)
}

// SelectOne 根据条件查询单条记录
func SelectOne[T any](q *QueryCond[T], opts ...OptionFunc) (*T, *gorm.DB) {
	var entity T
	resultDb := buildCondition(q, opts...)
	return &entity, resultDb.Take(&entity)
}

// SelectList 根据条件查询多条记录
func SelectList[T any](q *QueryCond[T], opts ...OptionFunc) ([]*T, *gorm.DB) {
	resultDb := buildCondition(q, opts...)
	var results []*T
	resultDb.Find(&results)
	return results, resultDb
}

// add start

// SelectByIdGeneric 查询时，转化为其他类型
// 第一个泛型代表数据库表实体
// 第二个泛型代表返回记录实体
func SelectByIdGeneric[T any, R any](id any, opts ...OptionFunc) (*R, *gorm.DB) {
	q, _ := NewQuery[T]()
	q.Eq(getPkColumnName[T](), id)
	var entity R
	resultDb := buildCondition(q, opts...)
	return &entity, resultDb.First(&entity)
}

// Pluck 取某列值，不去重
func Pluck[T any, R any](column string, q *QueryCond[T], opts ...OptionFunc) ([]R, *gorm.DB) {
	var results []R
	resultDb := buildCondition(q, opts...)
	resultDb.Pluck(column, &results)
	return results, resultDb
}

// PluckDistinct 取某列值，去重
func PluckDistinct[T any, R any](column string, q *QueryCond[T], opts ...OptionFunc) ([]R, *gorm.DB) {
	var results []R
	resultDb := buildCondition(q, opts...)
	resultDb.Distinct(column).Pluck(column, &results)
	return results, resultDb
}

// SelectListBySql 按任意SQL执行,指定返回类型数组
func SelectListBySql[R any](querySql string, opts ...OptionFunc) ([]*R, *gorm.DB) {
	resultDb := getDb(opts...)
	var results []*R
	resultDb = resultDb.Raw(querySql).Scan(&results)
	return results, resultDb
}

// SelectOneBySql 根据原始的SQL语句，取一个
func SelectOneBySql[R any](countSql string, opts ...OptionFunc) (R, *gorm.DB) {
	resultDb := getDb(opts...)
	var result R
	resultDb = resultDb.Raw(countSql).Scan(&result)
	return result, resultDb
}

// ExcSql 按任意SQL执行,返回影响的行
func ExcSql(querySql string, opts ...OptionFunc) *gorm.DB {
	resultDb := getDb(opts...)
	resultDb = resultDb.Exec(querySql)
	return resultDb
}

// add end

// SelectListGeneric 根据条件查询多条记录
// 第一个泛型代表数据库表实体
// 第二个泛型代表返回记录实体
func SelectListGeneric[T any, R any](q *QueryCond[T], opts ...OptionFunc) ([]*R, *gorm.DB) {
	resultDb := buildCondition(q, opts...)
	var results []*R
	resultDb.Scan(&results)
	return results, resultDb
}

// SelectPage 根据条件分页查询记录
func SelectPage[T any](page *Page[T], q *QueryCond[T], opts ...OptionFunc) (*Page[T], *gorm.DB) {
	option := getOption(opts)

	// 如果需要分页忽略总数，不查询总数
	if !option.IgnoreTotal {
		total, countDb := SelectCount[T](q, opts...)
		if countDb.Error != nil {
			return page, countDb
		}
		page.Total = total
	}

	resultDb := buildCondition(q, opts...)
	var results []*T
	resultDb.Scopes(paginate(page)).Find(&results)
	page.Records = results
	page.CurTime = time.Now().UnixMilli()
	return page, resultDb
}

// SelectCount 根据条件查询记录数量
func SelectCount[T any](q *QueryCond[T], opts ...OptionFunc) (int64, *gorm.DB) {
	var count int64
	resultDb := buildCondition(q, opts...)
	resultDb.Count(&count)
	return count, resultDb
}

// Exists 根据条件判断记录是否存在
func Exists[T any](q *QueryCond[T], opts ...OptionFunc) (bool, error) {
	count, resultDb := SelectCount[T](q, opts...)
	if resultDb.Error == gorm.ErrRecordNotFound {
		return false, nil
	}
	return count > 0, resultDb.Error
}

// SelectPageGeneric 根据传入的泛型封装分页记录
// 第一个泛型代表数据库表实体
// 第二个泛型代表返回记录实体
func SelectPageGeneric[T any, R any](page *Page[R], q *QueryCond[T], opts ...OptionFunc) (*Page[R], *gorm.DB) {
	option := getOption(opts)
	// 如果需要分页忽略总数，不查询总数
	if !option.IgnoreTotal {
		total, countDb := SelectCount[T](q, opts...)
		if countDb.Error != nil {
			return page, countDb
		}
		page.Total = total
	}
	resultDb := buildCondition(q, opts...)
	var r R
	switch any(r).(type) {
	case map[string]any:
		var results []R
		resultDb.Scopes(paginate(page)).Scan(&results)
		page.RecordsMap = results
	default:
		var results []*R
		resultDb.Scopes(paginate(page)).Scan(&results)
		page.Records = results
	}
	page.CurTime = time.Now().UnixMilli()
	return page, resultDb
}

// SelectGeneric 根据传入的泛型封装记录
// 第一个泛型代表数据库表实体
// 第二个泛型代表返回记录实体
func SelectGeneric[T any, R any](q *QueryCond[T], opts ...OptionFunc) (R, *gorm.DB) {
	var entity R
	resultDb := buildCondition(q, opts...)
	return entity, resultDb.Scan(&entity)
}

func Begin(opts ...*sql.TxOptions) *gorm.DB {
	db := getDb()
	return db.Begin(opts...)
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

func buildCondition[T any](q *QueryCond[T], opts ...OptionFunc) *gorm.DB {
	db := getDb(opts...)
	resultDb := db.Model(new(T))
	if q != nil {
		// 这里清空参数，避免用户重复使用一个query条件
		q.queryArgs = make([]any, 0)

		if len(q.distinctColumns) > 0 {
			resultDb.Distinct(q.distinctColumns)
		}

		if len(q.selectColumns) > 0 {
			resultDb.Select(q.selectColumns)
		}

		if len(q.omitColumns) > 0 {
			resultDb.Omit(q.omitColumns...)
		}

		expressions := q.queryExpressions
		if len(expressions) > 0 {
			var sqlBuilder strings.Builder
			q.queryArgs = buildSqlAndArgs[T](expressions, &sqlBuilder, q.queryArgs)
			resultDb.Where(sqlBuilder.String(), q.queryArgs...)
		}

		if q.orderBuilder.Len() > 0 {
			resultDb.Order(q.orderBuilder.String())
		}

		if q.groupBuilder.Len() > 0 {
			resultDb.Group(q.groupBuilder.String())
		}

		if q.havingBuilder.Len() > 0 {
			resultDb.Having(q.havingBuilder.String(), q.havingArgs...)
		}

		if q.limit != nil {
			resultDb.Limit(*q.limit)
		}

		if q.offset != 0 {
			resultDb.Offset(q.offset)
		}
	}
	return resultDb
}

func buildSqlAndArgs[T any](expressions []any, sqlBuilder *strings.Builder, queryArgs []any) []any {
	for _, v := range expressions {
		// 判断是否是columnValue类型
		switch segment := v.(type) {
		case *columnPointer:
			sqlBuilder.WriteString(segment.getSqlSegment() + " ")
		case *sqlKeyword:
			sqlBuilder.WriteString(segment.getSqlSegment() + " ")
		case *columnValue:
			if segment.value == constants.And {
				sqlBuilder.WriteString(segment.value.(string) + " ")
				continue
			}
			if segment.value != "" {
				sqlBuilder.WriteString("? ")
				queryArgs = append(queryArgs, segment.value)
			}
		case *QueryCond[T]:
			sqlBuilder.WriteString(constants.LeftBracket + " ")
			// 递归处理条件
			queryArgs = buildSqlAndArgs[T](segment.queryExpressions, sqlBuilder, queryArgs)
			sqlBuilder.WriteString(constants.RightBracket + " ")
		}
	}
	return queryArgs
}

func getDb(opts ...OptionFunc) *gorm.DB {
	option := getOption(opts)
	// Clauses()目的是为了初始化Db，如果db已经被初始化了,会直接返回db
	var db = globalDb.Clauses()

	if option.Db != nil {
		db = option.Db.Clauses()
	}

	// 设置需要忽略的字段
	setOmitIfNeed(option, db)

	// 设置选择的字段
	setSelectIfNeed(option, db)

	return db
}

func getOption(opts []OptionFunc) Option {
	var config Option
	for _, op := range opts {
		op(&config)
	}
	return config
}

func setSelectIfNeed(option Option, db *gorm.DB) {
	if len(option.Selects) > 0 {
		var columnNames []string
		for _, column := range option.Selects {
			columnName := getColumnName(column)
			columnNames = append(columnNames, columnName)
		}
		db.Select(columnNames)
	}
}

func setOmitIfNeed(option Option, db *gorm.DB) {
	if len(option.Omits) > 0 {
		var columnNames []string
		for _, column := range option.Omits {
			columnName := getColumnName(column)
			columnNames = append(columnNames, columnName)
		}
		db.Omit(columnNames...)
	}
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
