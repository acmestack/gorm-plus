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

package tests

import (
	"github.com/acmestack/gorm-plus/gplus"
	"gorm.io/gorm"
	"strings"
	"testing"
)

func TestInsert1Name(t *testing.T) {
	var expectSql = "INSERT INTO `Users` (`username`,`password`,`address`,`age`,`phone`,`score`,`dept`) VALUES ('afumu','123456','','18','','12','研发部门')"
	user := &User{Username: "afumu", Password: "123456", Age: 18, Score: 12, Dept: "研发部门"}
	u := gplus.GetModel[User]()
	sessionDb := checkInsertSql(t, expectSql)
	gplus.Insert(&user, gplus.Db(sessionDb), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))
}

func TestInsert2Name(t *testing.T) {
	var expectSql = "INSERT INTO `Users` (`username`,`password`,`address`,`age`,`phone`,`score`) VALUES ('afumu','123456','','18','','12')"
	user := &User{Username: "afumu", Password: "123456", Age: 18, Score: 12, Dept: "研发部门"}
	u := gplus.GetModel[User]()
	sessionDb := checkInsertSql(t, expectSql)
	gplus.Insert(&user, gplus.Db(sessionDb), gplus.Omit(&u.CreatedAt, &u.UpdatedAt, &u.Dept))
}

func TestInsert3Name(t *testing.T) {
	var expectSql = "INSERT INTO `Users` (`username`,`password`) VALUES ('afumu','123456')"
	user := &User{Username: "afumu", Password: "123456", Age: 18, Score: 12, Dept: "研发部门"}
	u := gplus.GetModel[User]()
	sessionDb := checkInsertSql(t, expectSql)
	gplus.Insert(&user, gplus.Db(sessionDb), gplus.Select(&u.Username, &u.Password), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))
}

func TestInsertBatchName(t *testing.T) {
	var expectSql = "INSERT INTO `Users` (`username`,`password`) VALUES ('afumu','123456'),('afumu','123456')"
	user := &User{Username: "afumu", Password: "123456", Age: 18, Score: 12, Dept: "研发部门"}
	user2 := &User{Username: "afumu", Password: "123456", Age: 18, Score: 12, Dept: "研发部门"}
	sessionDb := checkInsertSql(t, expectSql)
	u := gplus.GetModel[User]()
	gplus.InsertBatch([]*User{user, user2}, gplus.Db(sessionDb), gplus.Select(&u.Username, &u.Password), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))
}

func TestInsertBatchSizeName(t *testing.T) {
	var expectSql = "INSERT INTO `Users` (`username`,`password`) VALUES ('afumu','123456'),('afumu','123456')"
	user := &User{Username: "afumu", Password: "123456", Age: 18, Score: 12, Dept: "研发部门"}
	user2 := &User{Username: "afumu", Password: "123456", Age: 18, Score: 12, Dept: "研发部门"}
	user3 := &User{Username: "afumu", Password: "123456", Age: 18, Score: 12, Dept: "研发部门"}
	user4 := &User{Username: "afumu", Password: "123456", Age: 18, Score: 12, Dept: "研发部门"}
	sessionDb := checkInsertSql(t, expectSql)
	u := gplus.GetModel[User]()
	gplus.InsertBatchSize([]*User{user, user2, user3, user4}, 2, gplus.Db(sessionDb), gplus.Select(&u.Username, &u.Password), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))
}

func checkInsertSql(t *testing.T, expect string) *gorm.DB {
	expect = strings.TrimSpace(expect)
	sessionDb := gormDb.Session(&gorm.Session{DryRun: true})
	callback := sessionDb.Callback().Create().Before("gorm:insert")
	callback.Register("print_sql", func(db *gorm.DB) {
		sql := buildSql(db)
		sql = strings.TrimSpace(sql)
		if sql != expect {
			t.Errorf("errors happened  when insert expect: %v, got %v", expect, sql)
		}
		callback.Remove("print_sql")
	})
	return sessionDb
}
