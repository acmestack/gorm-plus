# gorm-plus
这是一个gorm的增强版，类似mybatis-plus语法。

This is an plus version of gorm, which is similar to the mybatis-plus syntax.

下载使用：

go get github.com/gorm-plus/gorm-plus

使用demo：

```go


func init() {
	dsn := "root:root-abcd-1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	gormplus.Init(gormDb)
}

type Test1 struct {
	gorm.Model
	Code  string
	Price uint
}

func TestSave(t *testing.T) {
	test1 := Test1{Code: "D1", Price: 100}
	resultDb := gormplus.Insert(&test1)
	fmt.Println(resultDb)
	fmt.Println(test1)
}

func TestSaveMigrate(t *testing.T) {
	test1 := Test1{Code: "D2", Price: 100}
	resultDb, err := gormplus.InsertMigrate(&test1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultDb)
	fmt.Println(test1)
}

func TestBatchSave(t *testing.T) {
	test1 := Test1{Code: "D3", Price: 100}
	test2 := Test1{Code: "D4", Price: 100}
	resultDb := gormplus.InsertBatch(&test1, &test2)
	fmt.Println(resultDb)
	fmt.Println(test1)
	fmt.Println(test2)
}

func TestSaveBatchMigrate(t *testing.T) {
	test1 := Test1{Code: "D5", Price: 100}
	test2 := Test1{Code: "D6", Price: 100}
	resultDb, err := gormplus.InsertBatchMigrate(&test1, &test2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultDb)
	fmt.Println(test1)
	fmt.Println(test2)
}

func TestDeleteById(t *testing.T) {
	resultDb := gormplus.DeleteById[Test1](1)
	fmt.Println(resultDb)
}

func TestDeleteByIds(t *testing.T) {
	resultDb := gormplus.DeleteByIds[Test1](4, 5)
	fmt.Println(resultDb)
}

func TestDelete(t *testing.T) {
	q := gormplus.Query[Test1]{}
	q.Eq("code", "D1").Eq("price", 100)
	resultDb := gormplus.Delete(&q)
	fmt.Println(resultDb)
}

func TestUpdateById(t *testing.T) {
	test1 := Test1{Code: "777"}
	resultDb := gormplus.UpdateById(6, &test1)
	fmt.Println(resultDb)
}

func TestUpdate(t *testing.T) {
	q := gormplus.Query[Test1]{}
	q.Eq("code", "D42").Eq("price", 100)
	test1 := Test1{Code: "888"}
	resultDb := gormplus.Update(&q, &test1)
	fmt.Println(resultDb)
}

func TestSelectById(t *testing.T) {
	db, result := gormplus.SelectById[Test1](1)
	fmt.Println(db)
	fmt.Println(result)
}

func TestSelectByIds(t *testing.T) {
	db, result := gormplus.SelectByIds[Test1](1, 2)
	fmt.Println(db)
	fmt.Println(result)
}

func TestSelectOne(t *testing.T) {
	q := gormplus.Query[Test1]{}
	q.Eq("code", "D42").Eq("price", 100)
	db, result := gormplus.SelectOne(&q)
	fmt.Println(db)
	fmt.Println(result)
}

func TestSelectList(t *testing.T) {
	q := gormplus.Query[Test1]{}
	q.Eq("price", 100)
	db, result := gormplus.SelectList(&q)
	fmt.Println(db.RowsAffected)
	fmt.Println(result)
}

func TestSelectCount(t *testing.T) {
	q := gormplus.Query[Test1]{}
	q.Eq("price", 100)
	db, count := gormplus.SelectCount(&q)
	fmt.Println(db)
	fmt.Println(count)
}

```
