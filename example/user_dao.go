package example

import "github.com/gorm-plus/gorm-plus/gormplus"

type UserDao[T any] struct {
	gormplus.CommonDao[T]
}

func NewUserDao[T any]() *UserDao[T] {
	return &UserDao[T]{}
}
