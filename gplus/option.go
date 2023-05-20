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
