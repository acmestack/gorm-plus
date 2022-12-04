package example

import (
	"fmt"
	"github.com/gorm-plus/gorm-plus/gormplus"
	"testing"
)

func TestUpdateById(t *testing.T) {
	user := &User{ID: 1, Username: "zhangsan", Password: "123456", Age: 18, Score: 100, Dept: "A部门asdfasdf"}
	result := gormplus.UpdateById(user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
}

func TestUpdate(t *testing.T) {
	q := gormplus.Query[User]{}
	q.Eq(UserColumn.Username, "zhangsan").Set(UserColumn.Dept, "相关部门123123").
		Set(UserColumn.Phone, 12312)
	result := gormplus.Update(&q)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
}
