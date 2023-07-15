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

func TestDeleteByIdName(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE `id` = 1"
	sessionDb := checkDeleteSql(t, expectSql)
	gplus.DeleteById[User](1, gplus.Db(sessionDb))
}

func TestDeleteByIdsName(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE `id` IN (1,2)"
	sessionDb := checkDeleteSql(t, expectSql)
	gplus.DeleteById[User]([]int{1, 2}, gplus.Db(sessionDb))
}

func TestDeleteEq(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE username = 'afumu'"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu")
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete2Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE username = 'afumu' OR username = 'afumu2'"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").Or().Eq(&u.Username, "afumu2")
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete3Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE username = 'afumu' OR ( username = 'afumu2' AND score = 12 )"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").Or(func(q *gplus.QueryCond[User]) {
		q.Eq(&u.Username, "afumu2").Eq(&u.Score, 12)
	})
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete4Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE username = 'afumu' AND ( username = 'afumu2' AND score = 12 )"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").And(func(q *gplus.QueryCond[User]) {
		q.Eq(&u.Username, "afumu2").Eq(&u.Score, 12)
	})
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete5Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE username = 'afumu' OR username = 'afumu2'"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").Or().Eq(&u.Username, "afumu2")
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete6Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE username = 'afumu' AND score = 60"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").And().Eq(&u.Score, 60)
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete7Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE score > 60"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Gt(&u.Score, 60)
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete8Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE score > 60"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Gt(&u.Score, 60)
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete9Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE score >= 60"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Ge(&u.Score, 60)
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete10Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE score < 60"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Lt(&u.Score, 60)
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete11Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE score <= 60"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Le(&u.Score, 60)
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete12Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE username LIKE '%zhang%'"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Like(&u.Username, "zhang")
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete13Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE username LIKE '%zhang'"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.LikeLeft(&u.Username, "zhang")
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete14Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE username LIKE 'zhang%'"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.LikeRight(&u.Username, "zhang")
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete15Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE username IS NULL"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.IsNull(&u.Username)
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete16Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE username IS NOT NULL"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.IsNotNull(&u.Username)
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete17Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE username IN ('afumu','afumu2')"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.In(&u.Username, []string{"afumu", "afumu2"})
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete18Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE username NOT IN ('afumu','afumu2')"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.NotIn(&u.Username, []string{"afumu", "afumu2"})
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete20Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE score BETWEEN 60 AND 80"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Between(&u.Score, 60, 80)
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete21Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE score NOT BETWEEN 60 AND 80"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.NotBetween(&u.Score, 60, 80)
	gplus.Delete(query, gplus.Db(sessionDb))
}

func TestDelete22Name(t *testing.T) {
	var expectSql = "DELETE FROM `Users` WHERE score NOT BETWEEN 60 AND 80"
	sessionDb := checkDeleteSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.NotBetween(&u.Score, 60, 80)
	gplus.Delete(query, gplus.Db(sessionDb))
}

func checkDeleteSql(t *testing.T, expect string) *gorm.DB {
	expect = strings.TrimSpace(expect)
	sessionDb := gormDb.Session(&gorm.Session{DryRun: true})
	callBack := sessionDb.Callback().Delete().Before("gorm:DELETE")
	callBack.Register("print_sql", func(db *gorm.DB) {
		sql := buildSql(db)
		sql = strings.TrimSpace(sql)
		if sql != expect {
			t.Errorf("errors happened  when delete expect: %v, got %v", expect, sql)
		}
		callBack.Remove("print_sql")
	})
	return sessionDb
}
