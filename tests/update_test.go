package tests

import (
	"github.com/acmestack/gorm-plus/gplus"
	"gorm.io/gorm"
	"strings"
	"testing"
)

func TestUpdateByIdName(t *testing.T) {
	var expectSql = "INSERT INTO `Users` (`username`,`password`,`address`,`age`,`phone`,`score`,`dept`) VALUES ('afumu','123456','','18','','12','研发部门')"
	sessionDb := checkUpdateSql(t, expectSql)
	var u = &User{ID: 1}
	gplus.UpdateById(u, gplus.Db(sessionDb), gplus.Omit(&u.CreatedAt, &u.UpdatedAt))

}

func checkUpdateSql(t *testing.T, expect string) *gorm.DB {
	expect = strings.TrimSpace(expect)
	sessionDb := gormDb.Session(&gorm.Session{DryRun: true})
	callback := sessionDb.Callback().Update().Before("gorm:update")
	callback.Register("print_sql", func(db *gorm.DB) {
		sql := buildSql(db)
		sql = strings.TrimSpace(sql)
		if sql != expect {
			t.Errorf("errors happened  when insert expect: %v, got %v", expect, sql)
		}
		callback.Remove("print_sql")
	})

	return sessionDb
}
