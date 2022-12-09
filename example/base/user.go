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

package base

import (
	"time"
)

// +gplus:column=true

type User struct {
	ID        int64
	Username  string `gorm:"column:username"`
	Password  string
	Address   string
	Age       int
	Phone     string
	Score     int
	Dept      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "Users"
}
