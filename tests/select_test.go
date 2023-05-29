package tests

import (
	"github.com/acmestack/gorm-plus/gplus"
	"gorm.io/gorm"
	"strings"
	"testing"
)

func TestSelectById1Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE id = '1'  ORDER BY `Users`.`id` LIMIT 1"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectById[User](1, gplus.Db(sessionDb))
}

func TestSelectById2Name(t *testing.T) {
	var expectSql = "SELECT `username`,`age` FROM `Users` WHERE id = '1'  ORDER BY `Users`.`id` LIMIT 1"
	sessionDb := checkSelectSql(t, expectSql)
	u := gplus.GetModel[User]()
	gplus.SelectById[User](1, gplus.Db(sessionDb), gplus.Select(&u.Username, &u.Age))
}

func TestSelectById3Name(t *testing.T) {
	var expectSql = "SELECT `Users`.`id`,`Users`.`password`,`Users`.`address`,`Users`.`phone`,`Users`.`score`,`Users`.`dept`,`Users`.`created_at`,`Users`.`updated_at` FROM `Users` WHERE id = '1'  ORDER BY `Users`.`id` LIMIT 1"
	sessionDb := checkSelectSql(t, expectSql)
	u := gplus.GetModel[User]()
	gplus.SelectById[User](1, gplus.Db(sessionDb), gplus.Omit(&u.Username, &u.Age))
}

func TestSelectByIdsName(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE id = '1'  ORDER BY `Users`.`id` LIMIT 1"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectByIds[User]([]int{1, 2}, gplus.Db(sessionDb))
}

func checkSelectSql(t *testing.T, expect string) *gorm.DB {
	expect = strings.TrimSpace(expect)
	sessionDb := gormDb.Session(&gorm.Session{DryRun: true})
	callback := sessionDb.Callback().Query().After("gorm:query")
	callback.Register("print_sql", func(db *gorm.DB) {
		sql := buildSql(db)
		sql = strings.TrimSpace(sql)
		if sql != expect {
			t.Errorf("errors happened  when select expect: %v, got %v", expect, sql)
		}
		callback.Remove("print_sql")
	})

	return sessionDb
}
