package gplux

import (
	"github.com/acmestack/gorm-plus/constants"
	"github.com/acmestack/gorm-plus/gplus"
)

type QueryCond[T any] struct {
	normalExpression []any
	last             any
}

type SubQueryCond[T any] struct {
	normalExpression []any
}

func (q *SubQueryCond[T]) getSqlSegment() string {
	return ""
}

// NewQuery 构建查询条件
func NewQuery[T any]() (*QueryCond[T], *T) {
	q := &QueryCond[T]{}
	m := new(T)
	gplus.Cache(m)
	return q, m
}

// Eq 等于 =
func (q *QueryCond[T]) Eq(column any, val any) *QueryCond[T] {

	return q
}

func (q *QueryCond[T]) handelExpression(sqlSegments ...SqlSegment) {
	if len(sqlSegments) == 1 {
		// 只有and(),or(),not() 会进入这里
		if len(q.normalExpression) == 0 {
			return
		}
		sqlSegment := sqlSegments[0]
		sk, ok := sqlSegment.(*sqlKeyword)
		lastKeyword, isKeyword := q.last.(*sqlKeyword)
		if ok && isKeyword {
			// 如果上一次是and，这一次也是and，那么就不需要重复设置了
			isAnd := lastKeyword.keyword == constants.And
			if isAnd && sk.keyword == constants.And {
				return
			}
			// 如果上一次是or，这一次也是or，那么就不需要重复设置了
			isOr := lastKeyword.keyword == constants.Or
			if isOr && sk.keyword == constants.Or {
				return
			}
			// 如果上一次是and，这次是or，那么删除上一次的值，使用最新的值
			// 如果上一次是or，这次是and，那么删除上一次的值，使用最新的值
			q.normalExpression = append(q.normalExpression, q.normalExpression[:len(q.normalExpression)-1]...)
			q.normalExpression = append(q.normalExpression, sqlSegment)
		}
	} else {
		lastKeyword, isKeyword := q.last.(*sqlKeyword)
		// 如果上一个值不是and或or,并且当前列表不为空,则添加and。
		if isKeyword && lastKeyword.keyword != constants.And && lastKeyword.keyword != constants.Or {
			sk := sqlKeyword{keyword: constants.And}
			q.normalExpression = append(q.normalExpression, &sk)
		}
	}

	for _, sqlSegment := range sqlSegments {
		q.normalExpression = append(q.normalExpression, sqlSegment)
	}
}
