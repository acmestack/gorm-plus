package common

import (
	"encoding/json"
	"fmt"
	"github.com/gorm-plus/gorm-plus/gplus"
	"testing"
)

func TestGetById(t *testing.T) {
	user, resultDb := userDao.GetById(2)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(user)
	fmt.Println("user1:", string(marshal))
}

func TestGetByOne(t *testing.T) {
	query := gplus.NewQuery[User]()
	query.Eq(UserColumn.Username, "zhangsan1")
	user, resultDb := userDao.GetOne(query)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(user)
	fmt.Println("user1:", string(marshal))
}

func TestListAll(t *testing.T) {
	users, resultDb := userDao.ListAll()
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, v := range users {
		marshal, _ := json.Marshal(v)
		fmt.Println("u:", string(marshal))
	}
}

func TestList(t *testing.T) {
	query := gplus.NewQuery[User]()
	query.Eq(UserColumn.Username, "zhangsan1")
	users, resultDb := userDao.List(query)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, v := range users {
		marshal, _ := json.Marshal(v)
		fmt.Println("u:", string(marshal))
	}
}

func TestPageAll(t *testing.T) {
	page := gplus.NewPage[User](1, 2)
	page, resultDb := userDao.PageAll(page)
	fmt.Println("page total:", page.Total)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, v := range page.Records {
		marshal, _ := json.Marshal(v)
		fmt.Println("u:", string(marshal))
	}
}

func TestPage(t *testing.T) {
	page := gplus.NewPage[User](1, 2)
	query := gplus.NewQuery[User]()
	query.Eq(UserColumn.Username, "zhangsan1")
	page, resultDb := userDao.Page(page, query)
	fmt.Println("page total:", page.Total)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, v := range page.Records {
		marshal, _ := json.Marshal(v)
		fmt.Println("u:", string(marshal))
	}
}
