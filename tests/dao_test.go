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
	"errors"
	"fmt"
	"github.com/acmestack/gorm-plus/gplus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

var gormDb *gorm.DB

func init() {
	dsn := "root:root-abcd-1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
	gplus.Init(gormDb)
}

func TestInsert(t *testing.T) {
	var u User
	gormDb.AutoMigrate(u)
	user := &User{Username: "user1", Password: "123456", Age: 18, Score: 100, Dept: "财务部门"}
	result := gplus.Insert(user)
	if result.Error != nil {
		t.Fatalf("errors happened when insert: %v", result.Error)
	} else if result.RowsAffected != 1 {
		t.Fatalf("rows affected expects: %v, got %v", 1, result.RowsAffected)
	}

	if user.ID == 0 {
		t.Errorf("user's primary key should has value after create, got : %v", user.ID)
	}

	if user.CreatedAt.IsZero() {
		t.Errorf("user's created at should be not zero")
	}

	if user.UpdatedAt.IsZero() {
		t.Errorf("user's updated at should be not zero")
	}

	var newUser User
	if err := gormDb.Where("id = ?", user.ID).First(&newUser).Error; err != nil {
		t.Fatalf("errors happened when query: %v", err)
	} else {
		AssertObjEqual(t, newUser, user, "ID", "Username", "Password", "Address", "Age", "Phone", "Score", "Dept", "CreatedAt", "UpdatedAt")
	}
}

func TestInsertBatch(t *testing.T) {
	user1 := &User{Username: "insert-batch-user1", Password: "123456", Age: 18, Score: 12, Dept: "财务部门"}
	user2 := &User{Username: "insert-batch-user2", Password: "123456", Age: 16, Score: 34, Dept: "行政部门"}
	user3 := &User{Username: "insert-batch-user3", Password: "123456", Age: 26, Score: 33, Dept: "研发部门"}
	user4 := &User{Username: "insert-batch-user4", Password: "123456", Age: 30, Score: 11, Dept: "产品部门"}
	user5 := &User{Username: "insert-batch-user5", Password: "123456", Age: 12, Score: 34, Dept: "生产部门"}
	user6 := &User{Username: "insert-batch-user5", Password: "123456", Age: 12, Score: 34, Dept: "生产部门"}
	user7 := &User{Username: "insert-batch-user6", Password: "123456", Age: 45, Score: 123, Dept: "销售部门"}
	users := []*User{user1, user2, user3, user4, user5, user6, user7}

	result := gplus.InsertBatch[User](users)

	if result.RowsAffected != int64(len(users)) {
		t.Errorf("affected rows should be %v, but got %v", len(users), result.RowsAffected)
	}

	for _, user := range users {
		if user.ID == 0 {
			t.Fatalf("failed to fill user's ID, got %v", user.ID)
		} else {
			var newUser User
			if err := gormDb.Where("id = ?", user.ID).Preload(clause.Associations).First(&newUser).Error; err != nil {
				t.Fatalf("errors happened when query: %v", err)
			} else {
				AssertObjEqual(t, newUser, user, "ID", "Username", "Password", "Address", "Age", "Phone", "Score", "Dept", "CreatedAt", "UpdatedAt")
			}
		}
	}
}

func TestInsertBatchSize(t *testing.T) {
	user1 := &User{Username: "insert-batch-size-user1", Password: "123456", Age: 18, Score: 12, Dept: "财务部门"}
	user2 := &User{Username: "insert-batch-size-user2", Password: "123456", Age: 16, Score: 34, Dept: "行政部门"}
	user3 := &User{Username: "insert-batch-size-user3", Password: "123456", Age: 26, Score: 33, Dept: "研发部门"}
	user4 := &User{Username: "insert-batch-size-user4", Password: "123456", Age: 30, Score: 11, Dept: "产品部门"}
	user5 := &User{Username: "insert-batch-size-user5", Password: "123456", Age: 12, Score: 34, Dept: "生产部门"}
	user6 := &User{Username: "insert-batch-size-user5", Password: "123456", Age: 12, Score: 34, Dept: "生产部门"}
	user7 := &User{Username: "insert-batch-size-user6", Password: "123456", Age: 45, Score: 123, Dept: "销售部门"}
	users := []*User{user1, user2, user3, user4, user5, user6, user7}

	result := gplus.InsertBatchSize[User](users, 2)

	if result.RowsAffected != int64(len(users)) {
		t.Errorf("affected rows should be %v, but got %v", len(users), result.RowsAffected)
	}

	for _, user := range users {
		if user.ID == 0 {
			t.Fatalf("failed to fill user's ID, got %v", user.ID)
		} else {
			var newUser User
			if err := gormDb.Where("id = ?", user.ID).Preload(clause.Associations).First(&newUser).Error; err != nil {
				t.Fatalf("errors happened when query: %v", err)
			} else {
				AssertObjEqual(t, newUser, user, "ID", "Username", "Password", "Address", "Age", "Phone", "Score", "Dept", "CreatedAt", "UpdatedAt")
			}
		}
	}
}

func TestDeleteById(t *testing.T) {
	user1 := &User{Username: "delete-user1", Password: "123456", Age: 18, Score: 12, Dept: "财务部门"}
	user2 := &User{Username: "delete-user2", Password: "123456", Age: 16, Score: 34, Dept: "行政部门"}
	user3 := &User{Username: "delete-user3", Password: "123456", Age: 26, Score: 33, Dept: "研发部门"}
	users := []*User{user1, user2, user3}

	if err := gplus.InsertBatch[User](users).Error; err != nil {
		t.Errorf("errors happened when insertBatch: %v", err)
	}

	for _, user := range users {
		if user.ID == 0 {
			t.Fatalf("user's primary key should has value after insert, got : %v", user.ID)
		}
	}

	if res := gplus.DeleteById[User](user2.ID); res.Error != nil || res.RowsAffected != 1 {
		t.Errorf("errors happened when deleteById: %v, affected: %v", res.Error, res.RowsAffected)
	}

	var result User
	if err := gormDb.Where("id = ?", user2.ID).First(&result).Error; err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("should returns record not found error, but got %v", err)
	}

	for _, user := range []*User{users[0], users[2]} {
		result = User{}
		if err := gormDb.Where("id = ?", user.ID).First(&result).Error; err != nil {
			t.Errorf("no error should returns when query %v, but got %v", user.ID, err)
		}
	}

}

func TestDeleteByIds(t *testing.T) {
	user1 := &User{Username: "delete-user1", Password: "123456", Age: 18, Score: 12, Dept: "财务部门"}
	user2 := &User{Username: "delete-user2", Password: "123456", Age: 16, Score: 34, Dept: "行政部门"}
	user3 := &User{Username: "delete-user3", Password: "123456", Age: 26, Score: 33, Dept: "研发部门"}
	user4 := &User{Username: "delete-user4", Password: "123456", Age: 26, Score: 33, Dept: "研发部门"}
	user5 := &User{Username: "delete-user5", Password: "123456", Age: 26, Score: 33, Dept: "研发部门"}
	users := []*User{user1, user2, user3, user4, user5}

	if err := gplus.InsertBatch[User](users).Error; err != nil {
		t.Errorf("errors happened when insertBatch: %v", err)
	}

	for _, user := range users {
		if user.ID == 0 {
			t.Fatalf("user's primary key should has value after insert, got : %v", user.ID)
		}
	}

	userIds := []int64{user1.ID, user2.ID, user3.ID}
	if res := gplus.DeleteByIds[User](userIds); res.Error != nil || res.RowsAffected != 3 {
		t.Errorf("errors happened when deleteByIds: %v, affected: %v", res.Error, res.RowsAffected)
	}

	var result1 User
	if err := gormDb.Where("id = ?", user1.ID).First(&result1).Error; err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("should returns record not found error, but got %v", err)
	}

	var result2 User
	if err := gormDb.Where("id = ?", user2.ID).First(&result2).Error; err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("should returns record not found error, but got %v", err)
	}

	var result3 User
	if err := gormDb.Where("id = ?", user3.ID).First(&result3).Error; err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("should returns record not found error, but got %v", err)
	}

	for _, user := range []*User{user4, user5} {
		result := User{}
		if err := gormDb.Where("id = ?", user.ID).First(&result).Error; err != nil {
			t.Errorf("no error should returns when query %v, but got %v", user.ID, err)
		}
	}
}

func TestUpdateById(t *testing.T) {
	user1suffix := fmt.Sprintf("%v", time.Now().UnixMilli())
	user1Name := "update-user" + user1suffix
	user5suffix := fmt.Sprintf("%v", time.Now().UnixMilli())
	user2Name := "update-user" + user5suffix
	user1 := &User{Username: user1Name, Password: "123456", Age: 18, Score: 12, Dept: "财务部门"}
	user2 := &User{Username: "update-user2", Password: "123456", Age: 16, Score: 34, Dept: "行政部门"}
	user3 := &User{Username: "update-user3", Password: "123456", Age: 26, Score: 33, Dept: "研发部门"}
	user4 := &User{Username: "update-user4", Password: "123456", Age: 16, Score: 33, Dept: "研发部门"}
	user5 := &User{Username: user2Name, Password: "123456", Age: 36, Score: 33, Dept: "研发部门"}
	users := []*User{user1, user2, user3, user4, user5}

	if err := gplus.InsertBatch[User](users).Error; err != nil {
		t.Errorf("errors happened when insertBatch: %v", err)
	}

	for _, user := range users {
		if user.ID == 0 {
			t.Fatalf("user's primary key should has value after insert, got : %v", user.ID)
		}
	}

	query, model := gplus.NewQuery[User]()
	// delete user1 and user5
	query.Eq(&model.Username, user1Name).Or().Eq(&model.Username, user2Name)

	if res := gplus.Delete(query); res.Error != nil || res.RowsAffected != 2 {
		t.Errorf("errors happened when deleteByIds: %v, affected: %v", res.Error, res.RowsAffected)
	}

	var result1 User
	if err := gormDb.Where("id = ?", user1.ID).First(&result1).Error; err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("should returns record not found error, but got %v", err)
	}

	var result2 User
	if err := gormDb.Where("id = ?", user5.ID).First(&result2).Error; err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("should returns record not found error, but got %v", err)
	}

	for _, user := range []*User{user2, user3, user4} {
		result := User{}
		if err := gormDb.Where("id = ?", user.ID).First(&result).Error; err != nil {
			t.Errorf("no error should returns when query %v, but got %v", user.ID, err)
		}
	}
}

func TestUpdate(t *testing.T) {
	user := &User{Username: "update-user1", Password: "123456", Age: 18, Score: 100, Dept: "财务部门"}
	result := gplus.Insert(user)
	if result.Error != nil {
		t.Fatalf("errors happened when insert: %v", result.Error)
	} else if result.RowsAffected != 1 {
		t.Fatalf("rows affected expects: %v, got %v", 1, result.RowsAffected)
	}

	q, model := gplus.NewQuery[User]()
	q.Eq(&model.ID, user.ID).Set(&model.Score, 60)
	if err := gplus.Update(q).Error; err != nil {
		t.Errorf("errors happened when update: %v", err)
	}
	user.Score = 60

	var firstUser User
	if err := gormDb.Where("id = ?", user.ID).First(&firstUser).Error; err != nil {
		t.Errorf("errors happened when query before user: %v", err)
	} else {
		AssertObjEqual(t, firstUser, user, "ID", "Username", "Password", "Address", "Age", "Phone", "Score", "Dept", "CreatedAt", "UpdatedAt")
	}
}

func TestSelectById(t *testing.T) {
	user := &User{Username: "user1", Password: "123456", Age: 18, Score: 100, Dept: "财务部门"}
	result := gplus.Insert(user)
	if result.Error != nil {
		t.Fatalf("errors happened when insert: %v", result.Error)
	} else if result.RowsAffected != 1 {
		t.Fatalf("rows affected expects: %v, got %v", 1, result.RowsAffected)
	}

	resultUser, db := gplus.SelectById[User](user.ID)
	if db.Error != nil {
		t.Errorf("errors happened when selectById : %v", db.Error)
	} else {
		AssertObjEqual(t, resultUser, user, "ID", "Username", "Password", "Address", "Age", "Phone", "Score", "Dept", "CreatedAt", "UpdatedAt")
	}

}

func TestSelectByIds(t *testing.T) {
	user1 := &User{Username: "user1", Password: "123456", Age: 18, Score: 100, Dept: "财务部门"}
	user2 := &User{Username: "user2", Password: "123456", Age: 18, Score: 100, Dept: "财务部门"}
	users := []*User{user1, user2}
	result := gplus.InsertBatch[User](users)
	if result.Error != nil {
		t.Fatalf("errors happened when insert: %v", result.Error)
	} else if result.RowsAffected != 2 {
		t.Fatalf("rows affected expects: %v, got %v", 2, result.RowsAffected)
	}
	userIds := []int64{user1.ID, user2.ID}
	resultUsers, db := gplus.SelectByIds[User](userIds)
	if db.Error != nil {
		t.Errorf("errors happened when selectByIds : %v", db.Error)
	} else {
		AssertObjEqual(t, resultUsers[0], user1, "ID", "Username", "Password", "Address", "Age", "Phone", "Score", "Dept", "CreatedAt", "UpdatedAt")
		AssertObjEqual(t, resultUsers[1], user2, "ID", "Username", "Password", "Address", "Age", "Phone", "Score", "Dept", "CreatedAt", "UpdatedAt")
	}
}

func TestSelectOne(t *testing.T) {
	user1suffix := fmt.Sprintf("%v", time.Now().UnixMilli())
	user1Name := "select-user" + user1suffix
	user := &User{Username: user1Name, Password: "123456", Age: 18, Score: 100, Dept: "财务部门"}
	result := gplus.Insert(user)
	if result.Error != nil {
		t.Fatalf("errors happened when insert: %v", result.Error)
	} else if result.RowsAffected != 1 {
		t.Fatalf("rows affected expects: %v, got %v", 1, result.RowsAffected)
	}

	query, model := gplus.NewQuery[User]()
	query.Eq(&model.Username, user1Name)
	resultUser, db := gplus.SelectOne[User](query)
	if db.Error != nil {
		t.Errorf("errors happened when selectByOne : %v", db.Error)
	} else {
		AssertObjEqual(t, resultUser, user, "ID", "Username", "Password", "Address", "Age", "Phone", "Score", "Dept", "CreatedAt", "UpdatedAt")
	}
}

func TestSelectList(t *testing.T) {
	user1suffix := fmt.Sprintf("%v", time.Now().UnixMilli())
	user1Name := "select-user" + user1suffix
	user5suffix := fmt.Sprintf("%v", time.Now().UnixMilli())
	user2Name := "select-user" + user5suffix
	user1 := &User{Username: user1Name, Password: "123456", Age: 18, Score: 12, Dept: "财务部门"}
	user2 := &User{Username: "select-user2", Password: "123456", Age: 16, Score: 34, Dept: "行政部门"}
	user3 := &User{Username: "select-user3", Password: "123456", Age: 26, Score: 33, Dept: "研发部门"}
	user4 := &User{Username: "select-user4", Password: "123456", Age: 16, Score: 33, Dept: "研发部门"}
	user5 := &User{Username: user2Name, Password: "123456", Age: 36, Score: 33, Dept: "研发部门"}
	users := []*User{user1, user2, user3, user4, user5}
	result := gplus.InsertBatch[User](users)
	if result.Error != nil {
		t.Fatalf("errors happened when insert: %v", result.Error)
	} else if result.RowsAffected != 5 {
		t.Fatalf("rows affected expects: %v, got %v", 5, result.RowsAffected)
	}

	query, model := gplus.NewQuery[User]()
	query.Eq(&model.Username, user1Name).Or().Eq(&model.Username, user2Name)
	resultUsers, db := gplus.SelectList(query)
	if db.Error != nil {
		t.Errorf("errors happened when SelectList : %v", db.Error)
	} else {
		AssertObjEqual(t, resultUsers[0], user1, "ID", "Username", "Password", "Address", "Age", "Phone", "Score", "Dept", "CreatedAt", "UpdatedAt")
		AssertObjEqual(t, resultUsers[1], user5, "ID", "Username", "Password", "Address", "Age", "Phone", "Score", "Dept", "CreatedAt", "UpdatedAt")
	}
}

func TestSelectPage(t *testing.T) {
	user1suffix := fmt.Sprintf("%v", time.Now().UnixMilli())
	user1Name := "select-user" + user1suffix
	user5suffix := fmt.Sprintf("%v", time.Now().UnixMilli())
	user2Name := "select-user" + user5suffix
	user1 := &User{Username: user1Name, Password: "123456", Age: 18, Score: 12, Dept: "财务部门"}
	user2 := &User{Username: "select-user2", Password: "123456", Age: 16, Score: 34, Dept: "行政部门"}
	user3 := &User{Username: "select-user3", Password: "123456", Age: 26, Score: 33, Dept: "研发部门"}
	user4 := &User{Username: "select-user4", Password: "123456", Age: 16, Score: 33, Dept: "研发部门"}
	user5 := &User{Username: user2Name, Password: "123456", Age: 36, Score: 33, Dept: "研发部门"}
	users := []*User{user1, user2, user3, user4, user5}
	result := gplus.InsertBatch[User](users)
	if result.Error != nil {
		t.Fatalf("errors happened when insert: %v", result.Error)
	} else if result.RowsAffected != 5 {
		t.Fatalf("rows affected expects: %v, got %v", 5, result.RowsAffected)
	}

	query, model := gplus.NewQuery[User]()
	page := gplus.NewPage[User](1, 10)
	query.Eq(&model.Username, user1Name).Or().Eq(&model.Username, user2Name)
	resultPage, db := gplus.SelectPage(page, query)
	if db.Error != nil {
		t.Errorf("errors happened when selectByIds : %v", db.Error)
	}
	if resultPage.Total != 2 {
		t.Errorf("page total expects: %v, got %v", 2, resultPage.Total)
	}

	AssertObjEqual(t, resultPage.Records[0], user1, "ID", "Username", "Password", "Address", "Age", "Phone", "Score", "Dept", "CreatedAt", "UpdatedAt")
	AssertObjEqual(t, resultPage.Records[1], user5, "ID", "Username", "Password", "Address", "Age", "Phone", "Score", "Dept", "CreatedAt", "UpdatedAt")

}

func TestSelectCount(t *testing.T) {
	user1suffix := fmt.Sprintf("%v", time.Now().UnixMilli())
	user1Name := "select-user" + user1suffix
	user5suffix := fmt.Sprintf("%v", time.Now().UnixMilli())
	user2Name := "select-user" + user5suffix
	user1 := &User{Username: user1Name, Password: "123456", Age: 18, Score: 12, Dept: "财务部门"}
	user2 := &User{Username: "select-user2", Password: "123456", Age: 16, Score: 34, Dept: "行政部门"}
	user3 := &User{Username: "select-user3", Password: "123456", Age: 26, Score: 33, Dept: "研发部门"}
	user4 := &User{Username: "select-user4", Password: "123456", Age: 16, Score: 33, Dept: "研发部门"}
	user5 := &User{Username: user2Name, Password: "123456", Age: 36, Score: 33, Dept: "研发部门"}
	users := []*User{user1, user2, user3, user4, user5}
	result := gplus.InsertBatch[User](users)
	if result.Error != nil {
		t.Fatalf("errors happened when insert: %v", result.Error)
	} else if result.RowsAffected != 5 {
		t.Fatalf("rows affected expects: %v, got %v", 5, result.RowsAffected)
	}

	query, model := gplus.NewQuery[User]()
	query.Eq(&model.Username, user1Name).Or().Eq(&model.Username, user2Name)
	count, db := gplus.SelectCount(query)
	if db.Error != nil {
		t.Errorf("errors happened when SelectCount : %v", db.Error)
	}
	if count != 2 {
		t.Errorf("count expects: %v, got %v", 2, count)
	}
}
