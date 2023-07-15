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

func TestUpdateByIdName(t *testing.T) {
	var expectSql = "UPDATE `Users` SET `score`=100 WHERE `id` = 1"
	sessionDb := checkUpdateSql(t, expectSql)
	var user = &User{ID: 1, Score: 100}
	u := gplus.GetModel[User]()
	gplus.UpdateById(user, gplus.Db(sessionDb), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))
}

func TestUpdateZeroByIdName(t *testing.T) {
	var expectSql = "UPDATE `Users` SET `username`='',`password`='',`address`='',`age`=0,`phone`='',`score`=100,`dept`='' WHERE `id` = 1"
	sessionDb := checkUpdateSql(t, expectSql)
	var user = &User{ID: 1, Score: 100}
	u := gplus.GetModel[User]()
	gplus.UpdateZeroById(user, gplus.Db(sessionDb), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))
}

func TestUpdate1Name(t *testing.T) {
	var expectSql = "UPDATE `Users` SET `score`=100 WHERE id = 1"
	sessionDb := checkUpdateSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.ID, 1).Set(&u.Score, 100)
	gplus.Update(query, gplus.Db(sessionDb), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))
}

func TestUpdate2Name(t *testing.T) {
	var expectSql = "UPDATE `Users` SET `address`='shanghai',`score`=100 WHERE username = 'afumu' AND age = 18"
	sessionDb := checkUpdateSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").Eq(&u.Age, 18).
		Set(&u.Score, 100).
		Set(&u.Address, "shanghai")
	gplus.Update(query, gplus.Db(sessionDb), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))
}

func TestUpdate3Name(t *testing.T) {
	var expectSql = "UPDATE `Users` SET `address`='shanghai',`score`=100 WHERE username = 'afumu' OR age = 18"
	sessionDb := checkUpdateSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").Or().Eq(&u.Age, 18).
		Set(&u.Score, 100).
		Set(&u.Address, "shanghai")
	gplus.Update(query, gplus.Db(sessionDb), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))
}

func TestUpdate4Name(t *testing.T) {
	var expectSql = "UPDATE `Users` SET `address`='shanghai',`score`=100 WHERE username = 'afumu' OR ( age = 18 AND score = 100 )"
	sessionDb := checkUpdateSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").Or(func(q *gplus.QueryCond[User]) {
		q.Eq(&u.Age, 18).Eq(&u.Score, 100)
	}).
		Set(&u.Score, 100).
		Set(&u.Address, "shanghai")
	gplus.Update(query, gplus.Db(sessionDb), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))
}

func TestUpdate5Name(t *testing.T) {
	var expectSql = "UPDATE `Users` SET `address`='shanghai',`score`=100 WHERE username = 'afumu' AND ( age = 18 OR score = 100 )"
	sessionDb := checkUpdateSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").
		And(func(q *gplus.QueryCond[User]) {
			q.Eq(&u.Age, 18).Or().Eq(&u.Score, 100)
		}).
		Set(&u.Score, 100).
		Set(&u.Address, "shanghai")
	gplus.Update(query, gplus.Db(sessionDb), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))
}

func TestUpdate6Name(t *testing.T) {
	var expectSql = "UPDATE `Users` SET `address`='shanghai',`score`=100 WHERE username <> 'afumu'"
	sessionDb := checkUpdateSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Ne(&u.Username, "afumu").
		Set(&u.Score, 100).
		Set(&u.Address, "shanghai")
	gplus.Update(query, gplus.Db(sessionDb), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))
}

func TestUpdate7Name(t *testing.T) {
	var expectSql = "UPDATE `Users` SET `address`='shanghai',`score`=100 WHERE username IS NULL"
	sessionDb := checkUpdateSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.IsNull(&u.Username).
		Set(&u.Score, 100).
		Set(&u.Address, "shanghai")
	gplus.Update(query, gplus.Db(sessionDb), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))
}

func checkUpdateSql(t *testing.T, expect string) *gorm.DB {
	expect = strings.TrimSpace(expect)
	sessionDb := gormDb.Session(&gorm.Session{DryRun: true})
	callback := sessionDb.Callback().Update().After("gorm:update")
	callback.Register("print_sql", func(db *gorm.DB) {
		sql := buildSql(db)
		sql = strings.TrimSpace(sql)
		if sql != expect {
			t.Errorf("errors happened  when update expect: %v, got %v", expect, sql)
		}
		callback.Remove("print_sql")
	})

	return sessionDb
}
