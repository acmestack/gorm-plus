package example

import (
	"fmt"
	"testing"
)

func TestSave(t *testing.T) {
	userDao := NewUserDao[User]()
	user1 := &User{Username: "zhangsan1", Password: "123456", Age: 18, Score: 12, Dept: "导弹部门"}
	resultDb := userDao.Save(user1)
	fmt.Println(resultDb.RowsAffected)
}
