package example

import (
	"encoding/json"
	"errors"
	"github.com/gorm-plus/gorm-plus/gormplus"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestSelectById(t *testing.T) {
	user, resultDb := gormplus.SelectById[User](1)
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

func TestSelectByIds(t *testing.T) {
	var ids []int
	ids = append(ids, 1)
	ids = append(ids, 2)
	users, resultDb := gormplus.SelectByIds[User](ids)
	if resultDb.Error != nil {
		log.Fatalln(resultDb.Error)
	}
	log.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(users)
	log.Println(string(marshal))
}

func TestSelectOne1(t *testing.T) {
	q := &gormplus.Query[User]{}
	q.Eq(UserColumn.Username, "zhangsan1")
	user, resultDb := gormplus.SelectOne(q)

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
	q := &gormplus.Query[User]{}
	q.Eq(UserColumn.Username, "zhangsan").
		Select(UserColumn.Username, UserColumn.Password)
	user, resultDb := gormplus.SelectOne(q)

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

func TestSelectList1(t *testing.T) {
	q := &gormplus.Query[User]{}
	q.Eq(UserColumn.Username, "zhangsan")
	users, resultDb := gormplus.SelectList(q)
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
	q := &gormplus.Query[User]{}
	q.Group(UserColumn.Dept).Select(UserColumn.Dept, "count(*) as count")
	users, resultDb := gormplus.SelectModelList[User, deptCount](q)
	if resultDb.Error != nil {
		log.Fatalln("error:", resultDb.Error)
	}
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		log.Println(string(marshal))
	}
}

func TestSelectPage(t *testing.T) {
	q := &gormplus.Query[User]{}
	q.Eq(UserColumn.Age, 18)
	page := &gormplus.Page[User]{Current: 1, Size: 10}
	pageResult, resultDb := gormplus.SelectPage(page, q)
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
	q := gormplus.NewQuery[User]()
	q.Group(UserColumn.Dept).Select(UserColumn.Dept, "count(*) as count")
	page := gormplus.NewPage[deptCount](1, 2)
	pageResult, resultDb := gormplus.SelectModelPage[User, deptCount](page, q)
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
	q := gormplus.NewQuery[User]()
	q.Eq(UserColumn.Age, 18)
	count, resultDb := gormplus.SelectCount(q)
	if resultDb.Error != nil {
		log.Fatalln("error:", resultDb.Error)
	}
	log.Println("count:", count)
}
