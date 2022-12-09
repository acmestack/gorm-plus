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
	"testing"
)

func TestSave(t *testing.T) {
	user1 := &User{Username: "zhangsan1", Password: "123456", Age: 18, Score: 12, Dept: "导弹部门"}
	resultDb := userDao.Save(user1)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(user1)
	fmt.Println("user1:", string(marshal))
}

func TestSaveBatch(t *testing.T) {
	user1 := &User{Username: "zhangsan1", Password: "123456", Age: 11, Score: 12, Dept: "导弹部门"}
	user2 := &User{Username: "zhangsan1", Password: "123456", Age: 13, Score: 12, Dept: "导弹部门"}
	var users []*User
	users = append(users, user1)
	users = append(users, user2)
	resultDb := userDao.SaveBatch(users)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		fmt.Println("user:", string(marshal))
	}
}

func TestSaveBatchSize(t *testing.T) {
	user1 := &User{Username: "zhangsan1", Password: "123456", Age: 11, Score: 12, Dept: "导弹部门"}
	user2 := &User{Username: "zhangsan1", Password: "123456", Age: 13, Score: 12, Dept: "导弹部门"}
	user3 := &User{Username: "zhangsan1", Password: "123456", Age: 13, Score: 12, Dept: "导弹部门"}
	user4 := &User{Username: "zhangsan1", Password: "123456", Age: 13, Score: 12, Dept: "导弹部门"}
	var users []*User
	users = append(users, user1)
	users = append(users, user2)
	users = append(users, user3)
	users = append(users, user4)
	resultDb := userDao.SaveBatchSize(users, 2)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		fmt.Println("user:", string(marshal))
	}
}
