package gplus

import "gorm.io/gorm"

type Option struct {
	Omits   []any
	Selects []any
	Db      *gorm.DB
}

type OptionFunc func(*Option)

func Omit(columns ...any) OptionFunc {
	return func(o *Option) {
		o.Omits = append(o.Omits, columns...)
	}
}

func Select(columns ...any) OptionFunc {
	return func(o *Option) {
		o.Selects = append(o.Selects, columns...)
	}
}

func Db(db *gorm.DB) OptionFunc {
	return func(o *Option) {
		o.Db = db
	}
}
