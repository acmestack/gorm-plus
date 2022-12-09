package common

import (
	"fmt"
	"github.com/gorm-plus/gorm-plus/gplus"
	"testing"
)

func TestRemoveById(t *testing.T) {
	resultDb := userDao.RemoveById(7)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
}

func TestRemoveByIds(t *testing.T) {
	var ids []int
	ids = append(ids, 5)
	ids = append(ids, 6)
	resultDb := userDao.RemoveByIds(ids)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
}

func TestRemove(t *testing.T) {
	query := gplus.NewQuery[User]()
	query.Eq(UserColumn.Username, "lisi")
	resultDb := userDao.Remove(query)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
}
