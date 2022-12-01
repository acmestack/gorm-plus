package mapper

import (
	"fmt"
	"github.com/zouchangfu/gorm-plus/constants"
	"strings"
)

type Query[T any] struct {
	Columns      []string
	QueryBuilder strings.Builder
	Args         []any
	LastCond     string
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
	q.Columns = append(q.Columns, columns...)
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
	q.Args = append(q.Args, val)
}
