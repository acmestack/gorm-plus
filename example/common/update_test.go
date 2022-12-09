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
	"github.com/acmestack/gorm-plus/gplus"
	"testing"
)

func TestUpdateById(t *testing.T) {
	user4 := &User{ID: 4, Username: "zhangsan666", Password: "123456", Age: 13, Score: 12, Dept: "导弹部门"}
	userDao.UpdateById(user4)
}

func TestUpdate(t *testing.T) {
	query := gplus.NewQuery[User]()
	query.Eq(UserColumn.Username, "zhangsan1").Set(UserColumn.Age, 50)
	userDao.Update(query)
}
