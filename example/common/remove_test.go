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
	"fmt"
	"github.com/gorm-plus/gorm-plus/gplus"
	"testing"
)

func TestRemoveById(t *testing.T) {
	resultDb := userDao.RemoveById(7)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
}

func TestRemoveByIds(t *testing.T) {
	var ids []int
	ids = append(ids, 5)
	ids = append(ids, 6)
	resultDb := userDao.RemoveByIds(ids)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
}

func TestRemove(t *testing.T) {
	query := gplus.NewQuery[User]()
	query.Eq(UserColumn.Username, "lisi")
	resultDb := userDao.Remove(query)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
}
