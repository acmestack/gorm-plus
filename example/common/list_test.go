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

package common

import (
	"encoding/json"
	"fmt"
	"github.com/acmestack/gorm-plus/gplus"
	"testing"
)

func TestGetById(t *testing.T) {
	user, resultDb := userDao.GetById(2)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(user)
	fmt.Println("user1:", string(marshal))
}

func TestGetByOne(t *testing.T) {
	query, model := gplus.NewQuery[User]()
	query.Eq(&model.Username, "zhangsan")
	user, resultDb := userDao.GetOne(query)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(user)
	fmt.Println("user1:", string(marshal))
}

func TestListAll(t *testing.T) {
	users, resultDb := userDao.ListAll()
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, v := range users {
		marshal, _ := json.Marshal(v)
		fmt.Println("u:", string(marshal))
	}
}

func TestList(t *testing.T) {
	query, model := gplus.NewQuery[User]()
	query.Eq(&model.Username, "zhangsan")
	users, resultDb := userDao.List(query)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, v := range users {
		marshal, _ := json.Marshal(v)
		fmt.Println("u:", string(marshal))
	}
}

func TestPageAll(t *testing.T) {
	page := gplus.NewPage[User](1, 2)
	page, resultDb := userDao.PageAll(page)
	fmt.Println("page total:", page.Total)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, v := range page.Records {
		marshal, _ := json.Marshal(v)
		fmt.Println("u:", string(marshal))
	}
}

func TestPage(t *testing.T) {
	page := gplus.NewPage[User](1, 2)
	query, model := gplus.NewQuery[User]()
	query.Eq(&model.Username, "zhangsan")
	page, resultDb := userDao.Page(page, query)
	fmt.Println("page total:", page.Total)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, v := range page.Records {
		marshal, _ := json.Marshal(v)
		fmt.Println("u:", string(marshal))
	}
}
