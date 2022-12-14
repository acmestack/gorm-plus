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
	"encoding/json"
	"fmt"
	"github.com/acmestack/gorm-plus/gplus"
	"testing"
)

func TestInsert(t *testing.T) {
	user := &User{Username: "zhangsan", Password: "123456", Age: 18, Score: 100, Dept: "A部门"}
	result := gplus.Insert(user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
}

func TestInsertBatch(t *testing.T) {
	user1 := &User{Username: "zhangsan1", Password: "123456", Age: 18, Score: 12, Dept: "导弹部门"}
	user2 := &User{Username: "lisi", Password: "123456", Age: 16, Score: 34, Dept: "投诉部门"}
	user3 := &User{Username: "wangwu", Password: "123456", Age: 26, Score: 33, Dept: "研发部门"}
	user4 := &User{Username: "zhangsan4", Password: "123456", Age: 30, Score: 11, Dept: "产品部门"}
	user5 := &User{Username: "zhangsan5", Password: "123456", Age: 12, Score: 34, Dept: "产品部门1"}
	user6 := &User{Username: "zhangsan6", Password: "123456", Age: 45, Score: 123, Dept: "产品部门12"}

	var users []*User
	users = append(users, user1)
	users = append(users, user2)
	users = append(users, user3)
	users = append(users, user4)
	users = append(users, user5)
	users = append(users, user6)

	result := gplus.InsertBatch[User](users)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		fmt.Println(string(marshal))
	}
}

func TestInsertBatchSize(t *testing.T) {
	user1 := &User{Username: "zhangsan1", Password: "123456", Age: 18, Score: 12, Dept: "导弹部门"}
	user2 := &User{Username: "lisi", Password: "123456", Age: 16, Score: 34, Dept: "投诉部门"}
	user3 := &User{Username: "wangwu", Password: "123456", Age: 26, Score: 33, Dept: "研发部门"}
	user4 := &User{Username: "zhangsan4", Password: "123456", Age: 30, Score: 11, Dept: "产品部门"}
	user5 := &User{Username: "zhangsan5", Password: "123456", Age: 12, Score: 34, Dept: "产品部门1"}
	user6 := &User{Username: "zhangsan6", Password: "123456", Age: 45, Score: 123, Dept: "产品部门12"}

	var users []*User
	users = append(users, user1)
	users = append(users, user2)
	users = append(users, user3)
	users = append(users, user4)
	users = append(users, user5)
	users = append(users, user6)

	result := gplus.InsertBatchSize[User](users, 3)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		fmt.Println(string(marshal))
	}
}
