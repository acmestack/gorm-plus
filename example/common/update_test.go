package common

import (
	"github.com/gorm-plus/gorm-plus/gplus"
	"testing"
)

func TestUpdateById(t *testing.T) {
	user4 := &User{ID: 4, Username: "zhangsan666", Password: "123456", Age: 13, Score: 12, Dept: "导弹部门"}
	userDao.UpdateById(user4)
}

func TestUpdate(t *testing.T) {
	query := gplus.NewQuery[User]()
	query.Eq(UserColumn.Username, "zhangsan1").Set(UserColumn.Age, 50)
	userDao.Update(query)
}
