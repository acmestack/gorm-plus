package gormplus

import (
	"fmt"
	"github.com/gorm-plus/gorm-plus/constants"
	"strings"
)

type Query[T any] struct {
	SelectColumns   []string
	DistinctColumns []string
	QueryBuilder    strings.Builder
	QueryArgs       []any
	OrderBuilder    strings.Builder
	GroupBuilder    strings.Builder
	HavingBuilder   strings.Builder
	HavingArgs      []any
	LastCond        string
}

func (q *Query[T]) Eq(column string, val any) *Query[T] {
	q.addCond(column, val, constants.Eq)
	return q
}

func (q *Query[T]) Ne(column string, val any) *Query[T] {
	q.addCond(column, val, constants.Ne)
	return q
}

func (q *Query[T]) Gt(column string, val any) *Query[T] {
	q.addCond(column, val, constants.Gt)
	return q
}

func (q *Query[T]) Ge(column string, val any) *Query[T] {
	q.addCond(column, val, constants.Ge)
	return q
}

func (q *Query[T]) Lt(column string, val any) *Query[T] {
	q.addCond(column, val, constants.Lt)
	return q
}

func (q *Query[T]) Le(column string, val any) *Query[T] {
	q.addCond(column, val, constants.Le)
	return q
}

func (q *Query[T]) Like(column string, val any) *Query[T] {
	s := val.(string)
	q.addCond(column, "%"+s+"%", constants.Like)
	return q
}

func (q *Query[T]) NotLike(column string, val any) *Query[T] {
	s := val.(string)
	q.addCond(column, "%"+s+"%", constants.Not+constants.Like)
	return q
}

func (q *Query[T]) LikeLeft(column string, val any) *Query[T] {
	s := val.(string)
	q.addCond(column, "%"+s, constants.Like)
	return q
}

func (q *Query[T]) LikeRight(column string, val any) *Query[T] {
	s := val.(string)
	q.addCond(column, s+"%", constants.Like)
	return q
}

func (q *Query[T]) In(column string, val ...any) *Query[T] {
	q.addCond(column, val, constants.In)
	return q
}

func (q *Query[T]) Distinct(column ...string) *Query[T] {
	q.DistinctColumns = column
	return q
}

func (q *Query[T]) And() *Query[T] {
	q.QueryBuilder.WriteString(constants.And)
	q.QueryBuilder.WriteString(" ")
	q.LastCond = constants.And
	return q
}

func (q *Query[T]) Or() *Query[T] {
	q.QueryBuilder.WriteString(constants.Or)
	q.QueryBuilder.WriteString(" ")
	q.LastCond = constants.Or
	return q
}

func (q *Query[T]) Select(columns ...string) *Query[T] {
	q.SelectColumns = append(q.SelectColumns, columns...)
	return q
}

func (q *Query[T]) OrderByDesc(columns ...string) *Query[T] {
	q.buildOrder(constants.Desc, columns...)
	return q
}

func (q *Query[T]) OrderByAsc(columns ...string) *Query[T] {
	q.buildOrder(constants.Asc, columns...)
	return q
}

func (q *Query[T]) Group(columns ...string) *Query[T] {
	for _, v := range columns {
		if q.GroupBuilder.Len() > 0 {
			q.GroupBuilder.WriteString(constants.COMMA)
		}
		q.GroupBuilder.WriteString(v)
	}
	return q
}

func (q *Query[T]) Having(having string, args ...any) *Query[T] {
	q.HavingBuilder.WriteString(having)
	q.HavingArgs = append(q.HavingArgs, args)
	return q
}

func (q *Query[T]) addCond(column string, val any, condType string) {
	if q.LastCond != constants.And && q.LastCond != constants.Or && q.QueryBuilder.Len() > 0 {
		q.QueryBuilder.WriteString(constants.And)
		q.QueryBuilder.WriteString(" ")
	}
	cond := fmt.Sprintf("%s %s ?", column, condType)
	q.QueryBuilder.WriteString(cond)
	q.QueryBuilder.WriteString(" ")
	q.LastCond = ""
	q.QueryArgs = append(q.QueryArgs, val)
}

func (q *Query[T]) buildOrder(orderType string, columns ...string) {
	for _, v := range columns {
		if q.OrderBuilder.Len() > 0 {
			q.OrderBuilder.WriteString(constants.COMMA)
		}
		q.OrderBuilder.WriteString(v)
		q.OrderBuilder.WriteString(" ")
		q.OrderBuilder.WriteString(orderType)
	}
}
