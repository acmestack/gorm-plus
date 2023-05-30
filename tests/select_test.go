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
	var expectSql = "SELECT * FROM `Users` WHERE id IN ('1','2')"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectByIds[User]([]int{1, 2}, gplus.Db(sessionDb))
}

func TestSelectOneName(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username = 'afumu'  ORDER BY `Users`.`id` LIMIT 1"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu")
	gplus.SelectOne[User](query, gplus.Db(sessionDb))
}

func TestSelectList1Name(t *testing.T) {
	var expectSql = " SELECT * FROM `Users` WHERE username = 'afumu'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu")
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList2Name(t *testing.T) {
	var expectSql = " SELECT * FROM `Users` WHERE username <> 'afumu'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Ne(&u.Username, "afumu")
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList3Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE age > '20'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Gt(&u.Age, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList4Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE age >= '20'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Ge(&u.Age, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList5Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE age < '20'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Lt(&u.Age, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList6Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE age <= '20'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Le(&u.Age, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList7Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username LIKE '%zhang%'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Like(&u.Username, "zhang")
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList8Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username LIKE '%zhang'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.LikeLeft(&u.Username, "zhang")
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList9Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username LIKE 'zhang%'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.LikeRight(&u.Username, "zhang")
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList10Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username is null"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.IsNull(&u.Username)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList11Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username is not null"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.IsNotNull(&u.Username)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList12Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username IN ('afumu','zhangsan')"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.In(&u.Username, []string{"afumu", "zhangsan"})
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList13Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username NOT IN ('afumu','zhangsan')"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.NotIn(&u.Username, []string{"afumu", "zhangsan"})
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList14Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE age BETWEEN '18' and '20'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Between(&u.Age, 18, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList17Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE age NOT BETWEEN '18' and '20'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.NotBetween(&u.Age, 18, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList20Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username = 'afumu' AND age = '20'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").And().Eq(&u.Age, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList21Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username = 'afumu' OR age = '20'"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").Or().Eq(&u.Age, 20)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList22Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username = 'afumu' OR (username = 'zhangsan' AND age = '30' )"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").Or(func(q *gplus.QueryCond[User]) {
		q.Eq(&u.Username, "zhangsan").And().Eq(&u.Age, 30)
	})
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestSelectList23Name(t *testing.T) {
	var expectSql = "SELECT * FROM `Users` WHERE username = 'afumu' AND (username = 'zhangsan' OR age = '30' )"
	sessionDb := checkSelectSql(t, expectSql)
	query, u := gplus.NewQuery[User]()
	query.Eq(&u.Username, "afumu").And(func(q *gplus.QueryCond[User]) {
		q.Eq(&u.Username, "zhangsan").Or().Eq(&u.Age, 30)
	})
	gplus.SelectList[User](query, gplus.Db(sessionDb))
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
