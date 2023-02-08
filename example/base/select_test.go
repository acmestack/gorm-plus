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
	"errors"
	"github.com/acmestack/gorm-plus/gplus"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

type Test2 struct {
	TestId    string `gorm:"primaryKey"`
	Code      string
	Price     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func TestSelectById(t *testing.T) {
	user, resultDb := gplus.SelectById[User]("or 1=1")
	if resultDb.Error != nil {
		if errors.Is(resultDb.Error, gorm.ErrRecordNotFound) {
			log.Fatalln("SelectById Data not found:", resultDb.Error)
		}
		log.Fatalln("SelectById error:", resultDb.Error)
	}
	log.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(user)
	log.Println(string(marshal))
}

func TestSelectByStrId(t *testing.T) {
	test, resultDb := gplus.SelectById[Test2]("a = 1 or 1=1")
	if resultDb.Error != nil {
		if errors.Is(resultDb.Error, gorm.ErrRecordNotFound) {
			log.Fatalln("SelectById Data not found:", resultDb.Error)
		}
		log.Fatalln("SelectById error:", resultDb.Error)
	}
	log.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(test)
	log.Println(string(marshal))
}

func TestSelectByIds(t *testing.T) {
	var ids []int
	ids = append(ids, 1)
	ids = append(ids, 2)
	users, resultDb := gplus.SelectByIds[User](ids)
	if resultDb.Error != nil {
		log.Fatalln(resultDb.Error)
	}
	log.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(users)
	log.Println(string(marshal))
}

func TestSelectOne1(t *testing.T) {
	q, model := gplus.NewQuery[User]()
	q.Eq(&model.Username, "zhangsan1")
	user, resultDb := gplus.SelectOne(q)
	if resultDb.Error != nil {
		if errors.Is(resultDb.Error, gorm.ErrRecordNotFound) {
			log.Fatalln("SelectOne Data not found:", resultDb.Error)
		}
		log.Fatalln("SelectOne error:", resultDb.Error)
	}
	log.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(user)
	log.Println(string(marshal))
}

func TestSelectOne2(t *testing.T) {
	q, model := gplus.NewQuery[User]()
	q.Eq(&model.Username, "zhangsan").
		Select(UserColumn.Username, UserColumn.Password)
	user, resultDb := gplus.SelectOne(q)

	if resultDb.Error != nil {
		if errors.Is(resultDb.Error, gorm.ErrRecordNotFound) {
			log.Fatalln("SelectOne Data not found:", resultDb.Error)
		}
		log.Fatalln("SelectOne error:", resultDb.Error)
	}

	log.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(user)
	log.Println(string(marshal))
}

func TestSelectList(t *testing.T) {
	q, model := gplus.NewQuery[User]()
	q.Eq(&model.Username, "zhangsan")
	users, resultDb := gplus.SelectList(q)
	if resultDb.Error != nil {
		log.Fatalln("error:", resultDb.Error)
	}
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		log.Println(string(marshal))
	}
}

func TestSelectBracketList(t *testing.T) {
	q, model := gplus.NewQuery[User]()
	bracketQuery, bracketModel := gplus.NewQuery[User]()
	bracketQuery.Eq(&bracketModel.Address, "上海").Or().Eq(&bracketModel.Address, "北京")

	q.Eq(&model.Username, "zhangsan").AndBracket(bracketQuery)
	users, resultDb := gplus.SelectList(q)
	if resultDb.Error != nil {
		log.Fatalln("error:", resultDb.Error)
	}
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		log.Println(string(marshal))
	}
}

func TestSelectTableList(t *testing.T) {
	type deptCount struct {
		Dept  string
		Count string
	}
	q, model := gplus.NewQuery[User]()
	q.Group(&model.Dept).Select(&model.Dept, "count(*) as count")
	users, resultDb := gplus.SelectListModel[User, deptCount](q)
	if resultDb.Error != nil {
		log.Fatalln("error:", resultDb.Error)
	}
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		log.Println(string(marshal))
	}
}

func TestSelectPage(t *testing.T) {
	q, model := gplus.NewQuery[User]()
	q.Eq(&model.Age, 18)
	page := gplus.NewPage[User](1, 10)
	pageResult, resultDb := gplus.SelectPage(page, q)
	if resultDb.Error != nil {
		log.Fatalln("error:", resultDb.Error)
	}
	log.Println("total", pageResult.Total)
	for _, u := range pageResult.Records {
		marshal, _ := json.Marshal(u)
		log.Println(string(marshal))
	}
}

func TestSelectTablePage(t *testing.T) {
	type deptCount struct {
		Dept  string
		Count string
	}
	q, model := gplus.NewQuery[User]()
	q.Group(&model.Dept).Select(UserColumn.Dept, "count(*) as count")
	page := gplus.NewPage[deptCount](1, 2)
	pageResult, resultDb := gplus.SelectPageModel[User, deptCount](page, q)
	if resultDb.Error != nil {
		log.Fatalln("error:", resultDb.Error)
	}
	log.Println("total:", pageResult.Total)
	for _, u := range pageResult.Records {
		marshal, _ := json.Marshal(u)
		log.Println(string(marshal))
	}
}

func TestSelectCount(t *testing.T) {
	q, model := gplus.NewQuery[User]()
	q.Eq(&model.Age, 18)
	count, resultDb := gplus.SelectCount(q)
	if resultDb.Error != nil {
		log.Fatalln("error:", resultDb.Error)
	}
	log.Println("count:", count)
}
