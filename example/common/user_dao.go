package common

import "github.com/gorm-plus/gorm-plus/gplus"

var userDao = NewUserDao[User]()

type UserDao[T any] struct {
	gplus.CommonDao[T]
}

func NewUserDao[T any]() *UserDao[T] {
	return &UserDao[T]{}
}
