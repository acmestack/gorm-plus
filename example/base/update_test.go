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
	"fmt"
	"github.com/acmestack/gorm-plus/gplus"
	"testing"
)

func TestUpdateById(t *testing.T) {
	user := &User{ID: 1, Username: "zhangsan", Password: "123456", Age: 18, Score: 100, Dept: "A部门asdfasdf"}
	result := gplus.UpdateById(user, user.ID)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
}

func TestUpdate(t *testing.T) {
	q, _ := gplus.NewQuery[User]()
	q.Eq(UserColumn.Username, "zhangsan").Set(UserColumn.Dept, "相关部门123123").
		Set(UserColumn.Phone, 12312)
	result := gplus.Update(q)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
}
