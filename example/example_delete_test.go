package example

import (
	"fmt"
	"github.com/gorm-plus/gorm-plus/gormplus"
	"testing"
)

func TestDeleteById(t *testing.T) {
	result := gormplus.DeleteById[User](13)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
}

func TestDelete(t *testing.T) {
	q := gormplus.NewQuery[User]()
	q.Ge(UserColumn.Age, 50)
	result := gormplus.Delete(q)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
}
