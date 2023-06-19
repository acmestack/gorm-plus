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

import "gorm.io/gorm"

type Option struct {
	Db          *gorm.DB
	Selects     []any
	Omits       []any
	IgnoreTotal bool
}

type OptionFunc func(*Option)

// Db 使用传入的Db对象
func Db(db *gorm.DB) OptionFunc {
	return func(o *Option) {
		o.Db = db
	}
}

// Session 创建回话
func Session(session *gorm.Session) OptionFunc {
	return func(o *Option) {
		o.Db = globalDb.Session(session)
	}
}

// Select 指定需要查询的字段
func Select(columns ...any) OptionFunc {
	return func(o *Option) {
		o.Selects = append(o.Selects, columns...)
	}
}

// Omit 指定需要忽略的字段
func Omit(columns ...any) OptionFunc {
	return func(o *Option) {
		o.Omits = append(o.Omits, columns...)
	}
}

// IgnoreTotal 分页查询忽略总数 issue: https://github.com/aixj1984/gorm-plus/issues/37
func IgnoreTotal() OptionFunc {
	return func(o *Option) {
		o.IgnoreTotal = true
	}
}
