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
	"strings"
	"testing"

	"github.com/aixj1984/gorm-plus/gplus"
	"gorm.io/gorm"
)

func TestSelectByIdName(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE id = 1  LIMIT 1"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectById[User](1, gplus.Db(sessionDb))
}

func TestSelectByIdSelect(t *testing.T) {
	var expectSql = "SELECT `username`,`age` FROM `Users` WHERE id = 1  LIMIT 1"
	sessionDb := checkSelectSql(t, expectSql)
	u := gplus.GetModel[User]()
	gplus.SelectById[User](1, gplus.Db(sessionDb), gplus.Select(&u.Username, &u.Age))
}

func TestSelectByIdOmit(t *testing.T) {
	var expectSql = "SELECT `Users`.`id`,`Users`.`created_at`,`Users`.`updated_at`,`Users`.`password`,`Users`.`address`,`Users`.`phone`,`Users`.`score`,`Users`.`dept` FROM `Users` WHERE id = 1  LIMIT 1"
	sessionDb := checkSelectSql(t, expectSql)
	u := gplus.GetModel[User]()
	gplus.SelectById[User](1, gplus.Db(sessionDb), gplus.Omit(&u.Username, &u.Age))
}

func TestSelectByIdsIn(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE id IN (1,2)"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectByIds[User]([]int{1, 2}, gplus.Db(sessionDb))
}

func TestSelectOneName(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username = 'afumu'  LIMIT 1"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu")
	gplus.SelectOne[User](query, gplus.Db(sessionDb))
}

func TestSelectListEq(t *testing.T) {
	var expectSql = " SELECT * FROM `Users` WHERE username = 'afumu'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu")
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListNe(t *testing.T) {
	var expectSql = " SELECT * FROM `Users` WHERE username <> 'afumu'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Ne(&u.Username, "afumu")
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListGt(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE age > 20"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Gt(&u.Age, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListGe(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE age >= 20"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Ge(&u.Age, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListLt(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE age < 20"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Lt(&u.Age, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListLe(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE age <= 20"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Le(&u.Age, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListLike(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username LIKE '%zhang%'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Like(&u.Username, "zhang")
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListLeftLike(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username LIKE '%zhang'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.LikeLeft(&u.Username, "zhang")
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListNotLeftLike(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username NOT LIKE '%zhang'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.NotLikeLeft(&u.Username, "zhang")
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListRightLike(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username LIKE 'zhang%'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.LikeRight(&u.Username, "zhang")
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListNotRightLike(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username NOT LIKE 'zhang%'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.NotLikeRight(&u.Username, "zhang")
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListIsNull(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username IS NULL"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.IsNull(&u.Username)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListIsNotNull(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username IS NOT NULL"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.IsNotNull(&u.Username)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListIn(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username IN ('afumu','zhangsan')"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.In(&u.Username, []string{"afumu", "zhangsan"})
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListNotIn(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username NOT IN ('afumu','zhangsan')"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.NotIn(&u.Username, []string{"afumu", "zhangsan"})
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListBetween(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE age BETWEEN 18 AND 20"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Between(&u.Age, 18, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListNotBetween(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE age NOT BETWEEN 18 AND 20"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.NotBetween(&u.Age, 18, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListAnd(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username = 'afumu' AND age = 20"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").And().Eq(&u.Age, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListOr(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username = 'afumu' OR age = 20"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").Or().Eq(&u.Age, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListOrNest(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username = 'afumu' OR ( username = 'zhangsan' AND age = 30 )"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").Or(func(q *gplus.QueryCond[User]) {
		q.Eq(&u.Username, "zhangsan").And().Eq(&u.Age, 30)
	})
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListAndNest(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username = 'afumu' AND ( username = 'zhangsan' OR age = 30 )"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").And(func(q *gplus.QueryCond[User]) {
		q.Eq(&u.Username, "zhangsan").Or().Eq(&u.Age, 30)
	})
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListAndOrNest(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE ( username = 'afumu' AND ( password = '123456' OR score = 60 ) OR dept = '开发' ) AND address = '北京' "
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.And(func(q *gplus.QueryCond[User]) {
		q.Eq(&u.Username, "afumu").And(func(q *gplus.QueryCond[User]) {
			q.Eq(&u.Password, "123456").Or().Eq(&u.Score, 60)
		}).Or().Eq(&u.Dept, "开发")
	}).Eq(&u.Address, "北京")
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListOrder(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` ORDER BY username DESC,age ASC"
	sessionDb := checkSelectSql(t, expectSql)
	query, user := gplus.NewQuery[User]()
	query.OrderByDesc(&user.Username).OrderByAsc(&user.Age)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectListQueryModel(t *testing.T) {
	var expectSql = "SELECT username AS name,`age` FROM `Users` WHERE username = 'afumu' AND ( address = '北京' OR age = 20 ) "
	sessionDb := checkSelectSql(t, expectSql)
	type UserVo struct {
		Name string
		Age  int64
	}
	query, user, userVo := gplus.NewQueryModel[User, UserVo]()
	query.Eq(&user.Username, "afumu").And(func(q *gplus.QueryCond[User]) {
		q.Eq(&user.Address, "北京").Or().Eq(&user.Age, 20)
	}).Select(gplus.As(&user.Username, &userVo.Name), &user.Age)
	gplus.SelectGeneric[User, []UserVo](query, gplus.Db(sessionDb))
}

func TestSelectListQueryModelSum(t *testing.T) {
	var expectSql = "SELECT `username`,SUM(age) AS total FROM `Users` GROUP BY `username` HAVING SUM(age) NOT BETWEEN 333 AND 1000"
	sessionDb := checkSelectSql(t, expectSql)
	type UserVo struct {
		Username string
		Total    int64
	}
	query, user, userVo := gplus.NewQueryModel[User, UserVo]()
	query.Group(&user.Username).
		Select(&user.Username, gplus.Sum(&user.Age).As(&userVo.Total)).
		Having(gplus.Sum(&user.Age).NotBetween(333, 1000))
	gplus.SelectGeneric[User, []UserVo](query, gplus.Db(sessionDb))
}

func TestSelectListQueryModelCount(t *testing.T) {
	var expectSql = "SELECT `username`,COUNT(age) AS total FROM `Users` GROUP BY `username`"
	sessionDb := checkSelectSql(t, expectSql)
	type UserVo struct {
		Username string
		Total    int64
	}
	query, user, userVo := gplus.NewQueryModel[User, UserVo]()
	query.Group(&user.Username).
		Select(&user.Username, gplus.Count(&user.Age).As(&userVo.Total))
	gplus.SelectGeneric[User, []UserVo](query, gplus.Db(sessionDb))
}

func checkSelectSql(t *testing.T, expect string) *gorm.DB {
	expect = strings.TrimSpace(expect)
	sessionDb := gormDb.Session(&gorm.Session{DryRun: true})
	callback := sessionDb.Callback().Query().After("gorm:query")
	callback.Register("print_sql", func(db *gorm.DB) {
		sql := buildSql(db)
		sql = strings.TrimSpace(sql)
		if sql != expect {
			t.Errorf("errors happened  when select expect: %v, got %v", expect, sql)
		}
		callback.Remove("print_sql")
	})

	return sessionDb
}
