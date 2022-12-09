package example

import (
	"encoding/json"
	"fmt"
	"testing"
)

var saveUserDao *UserDao[User]

func init() {
	saveUserDao = NewUserDao[User]()
}

func TestSave(t *testing.T) {
	user1 := &User{Username: "zhangsan1", Password: "123456", Age: 18, Score: 12, Dept: "导弹部门"}
	resultDb := saveUserDao.Save(user1)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(user1)
	fmt.Println("user1:", string(marshal))
}

func TestSaveBatch(t *testing.T) {
	user1 := &User{Username: "zhangsan1", Password: "123456", Age: 11, Score: 12, Dept: "导弹部门"}
	user2 := &User{Username: "zhangsan1", Password: "123456", Age: 13, Score: 12, Dept: "导弹部门"}
	var users []*User
	users = append(users, user1)
	users = append(users, user2)
	resultDb := saveUserDao.SaveBatch(users)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		fmt.Println("user:", string(marshal))
	}
}

func TestSaveBatchSize(t *testing.T) {
	user1 := &User{Username: "zhangsan1", Password: "123456", Age: 11, Score: 12, Dept: "导弹部门"}
	user2 := &User{Username: "zhangsan1", Password: "123456", Age: 13, Score: 12, Dept: "导弹部门"}
	user3 := &User{Username: "zhangsan1", Password: "123456", Age: 13, Score: 12, Dept: "导弹部门"}
	user4 := &User{Username: "zhangsan1", Password: "123456", Age: 13, Score: 12, Dept: "导弹部门"}
	var users []*User
	users = append(users, user1)
	users = append(users, user2)
	users = append(users, user3)
	users = append(users, user4)
	resultDb := saveUserDao.SaveBatchSize(users, 2)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		fmt.Println("user:", string(marshal))
	}
}
