package gplux

import (
	"github.com/acmestack/gorm-plus/constants"
	"gorm.io/gorm"
	"strings"
)

var globalDb *gorm.DB
var defaultBatchSize = 1000

func Init(db *gorm.DB) {
	globalDb = db
}

// SelectList 根据条件查询多条记录
func SelectList[T any](q *QueryCond[T]) ([]*T, *gorm.DB) {
	var results []*T
	resultDb := buildCondition(q)
	return results, resultDb.Find(&results)
}

func buildCondition[T any](q *QueryCond[T]) *gorm.DB {
	resultDb := globalDb.Model(new(T))
	if q != nil {
		expressions := q.queryExpressions
		var sqlBuilder strings.Builder
		if len(expressions) > 0 {
			q.queryArgs = buildSqlAndArgs[T](expressions, &sqlBuilder, q.queryArgs)
			resultDb.Where(sqlBuilder.String(), q.queryArgs...)
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
			sqlBuilder.WriteString(constants.LeftBracket)
			args := buildSqlAndArgs[T](segment.queryExpressions, sqlBuilder, queryArgs)
			sqlBuilder.WriteString(constants.RightBracket)
			return args
		}
	}
	return queryArgs
}
