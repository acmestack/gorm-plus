package base

import (
	"fmt"
	"github.com/gorm-plus/gorm-plus/gplus"
	"testing"
)

func TestDeleteById(t *testing.T) {
	result := gplus.DeleteById[User](13)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
}

func TestDelete(t *testing.T) {
	q := gplus.NewQuery[User]()
	q.Ge(UserColumn.Age, 50)
	result := gplus.Delete(q)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
}
