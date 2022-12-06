package example

import (
	"encoding/json"
	"log"
	"testing"
)

func TestList(t *testing.T) {
	userDao := NewUserDao[User]()
	list, resultDb := userDao.List(nil)
	log.Println(resultDb.RowsAffected)
	for _, v := range list {
		marshal, _ := json.Marshal(v)
		log.Println(string(marshal))
	}
}
