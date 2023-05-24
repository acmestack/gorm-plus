package tests

import (
	"fmt"
	"github.com/acmestack/gorm-plus/gplus"
	"gorm.io/gorm"
	"strings"
	"testing"
)

func TestInsert1Name(t *testing.T) {
	var expect = "INSERT INTO `Users` (`username`,`password`,`address`,`age`,`phone`,`score`,`dept`) VALUES ('afumu','123456','','18','','12','研发部门')"
	user := &User{Username: "afumu", Password: "123456", Age: 18, Score: 12, Dept: "研发部门"}
	u := gplus.GetModel[User]()

	db := gplus.Insert(&user, gplus.Session(&gorm.Session{DryRun: true}), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))
	stmt := db.Statement
	sql := stmt.SQL.String()
	for _, value := range stmt.Vars {
		sql = strings.Replace(sql, "?", fmt.Sprintf("'%v'", value), 1)
	}
	if sql != expect {
		t.Errorf("errors happened when Insert:expect: %v, got %v", expect, sql)
	}
}

func TestInsert2Name(t *testing.T) {
	var expect = "INSERT INTO `Users` (`username`,`password`,`address`,`age`,`phone`,`score`) VALUES ('afumu','123456','','18','','12')"
	user := &User{Username: "afumu", Password: "123456", Age: 18, Score: 12, Dept: "研发部门"}
	u := gplus.GetModel[User]()
	db := gplus.Insert(&user, gplus.Session(&gorm.Session{DryRun: true}), gplus.Omit(&u.CreatedAt, &u.UpdatedAt, &u.Dept))
	stmt := db.Statement
	sql := stmt.SQL.String()
	for _, value := range stmt.Vars {
		sql = strings.Replace(sql, "?", fmt.Sprintf("'%v'", value), 1)
	}
	if sql != expect {
		t.Errorf("errors happened when Insert:expect: %v, got %v", expect, sql)
	}
}

func TestInsert3Name(t *testing.T) {
	var expect = "INSERT INTO `Users` (`username`,`password`) VALUES ('afumu','123456')"
	user := &User{Username: "afumu", Password: "123456", Age: 18, Score: 12, Dept: "研发部门"}
	u := gplus.GetModel[User]()
	db := gplus.Insert(&user, gplus.Session(&gorm.Session{DryRun: true}), gplus.Select(&u.Username, &u.Password), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))
	stmt := db.Statement
	sql := stmt.SQL.String()
	for _, value := range stmt.Vars {
		sql = strings.Replace(sql, "?", fmt.Sprintf("'%v'", value), 1)
	}
	if sql != expect {
		t.Errorf("errors happened when Insert:expect: %v, got %v", expect, sql)
	}
}
