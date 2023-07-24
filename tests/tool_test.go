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
	"github.com/aixj1984/gorm-plus/gplus"
	"net/url"
	"testing"
)

func TestQueryById(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"id=1"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE id = 1"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByIdSelect(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"id=1"}
	values["select"] = []string{"username,age"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT `username`,`age` FROM `Users` WHERE id = 1"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByIdOmit(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"id=1"}
	values["omit"] = []string{"username,age"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT `Users`.`id`,`Users`.`created_at`,`Users`.`updated_at`,`Users`.`password`,`Users`.`address`,`Users`.`phone`,`Users`.`score`,`Users`.`dept` FROM `Users` WHERE id = 1"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByIdSortAsc(t *testing.T) {
	values := url.Values{}
	values["sort"] = []string{"age"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` ORDER BY age ASC"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByIdSortDesc(t *testing.T) {
	values := url.Values{}
	values["sort"] = []string{"-age"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` ORDER BY age DESC"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByIdsIn(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"id?=1,2"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE id IN (1,2)"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByEqUsername(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"username=afumu"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE username = 'afumu'"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByNeUsername(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"username!=afumu"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE username <> 'afumu'"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByGtAge(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"age>20"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE age > 20"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByGeAge(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"age>=20"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE age >= 20"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByLtAge(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"age<20"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE age < 20"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByLeAge(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"age<=20"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE age <= 20"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByLike(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"username~=afumu"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE username LIKE '%afumu%'"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByLeftLike(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"username~<=afumu"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE username LIKE '%afumu'"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByNotLeftLike(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"username!~<=afumu"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE username NOT LIKE '%afumu'"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByRightLike(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"username~>=afumu"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE username LIKE 'afumu%'"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByNotRightLike(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"username!~>=afumu"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE username NOT LIKE 'afumu%'"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByIsNull(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"username=null"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE username IS NULL"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByIsNotNull(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"username!=null"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE username IS NOT NULL"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByIn(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"username?=afumu,zhangsan"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE username IN ('afumu','zhangsan')"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByNotIn(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"username!?=afumu,zhangsan"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE username NOT IN ('afumu','zhangsan')"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByBetween(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"age^=20,30"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE age BETWEEN 20 AND 30"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByNotBetween(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"age!^=20,30"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE age NOT BETWEEN 20 AND 30"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByAnd(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"useranme=afumu", "age=20"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE useranme = 'afumu' AND age = 20"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByGroupOnlyOne(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"A.useranme=afumu", "A.age=20"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE useranme = 'afumu' AND age = 20"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByGroupAAndB(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"A.useranme=afumu", "A.password=123456", "B.age=20", "B.score=90"}
	values["gcond"] = []string{"A*B"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE useranme = 'afumu' AND password = '123456' AND age = 20 AND score = 90"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByGroupAOrB(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"A.useranme=afumu", "A.password=123456", "B.age=20", "B.score=90"}
	values["gcond"] = []string{"A|B"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE useranme = 'afumu' AND password = '123456' OR age = 20 AND score = 90"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByGroupNest(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{
		"A.useranme=afumu", "B.password=12345", "C.score=60",
		"D.dept=开发", "F.address=北京",
	}
	values["gcond"] = []string{"(A*(B|C)|D)*F"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE ( useranme = 'afumu' AND ( password = '12345' OR score = 60 ) OR dept = '开发' ) AND address = '北京'"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}
