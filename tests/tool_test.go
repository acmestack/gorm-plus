package tests

import (
	"github.com/acmestack/gorm-plus/gplus"
	"net/url"
	"testing"
)

func TestQueryById(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"id=1"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE id = 1"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByIdSelect(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"id=1"}
	values["select"] = []string{"username,age"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT `username`,`age` FROM `Users` WHERE id = 1"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByIdOmit(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"id=1"}
	values["omit"] = []string{"username,age"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT `Users`.`id`,`Users`.`password`,`Users`.`address`,`Users`.`phone`,`Users`.`score`,`Users`.`dept`,`Users`.`created_at`,`Users`.`updated_at` FROM `Users` WHERE id = 1"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByIdsIn(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"id?=1,2"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE id IN (1,2)"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByEqUsername(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"username=afumu"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE username = 'afumu'"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByNeUsername(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"username!=afumu"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE username <> 'afumu'"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}

func TestQueryByGtAge(t *testing.T) {
	values := url.Values{}
	values["q"] = []string{"username!=afumu"}
	query := gplus.BuildQuery[User](values)
	var expectSql = "SELECT * FROM `Users` WHERE username <> 'afumu'"
	sessionDb := checkSelectSql(t, expectSql)
	gplus.SelectList[User](query, gplus.Db(sessionDb))
}
