# gorm-plus
这是一个gorm的增强版，类似mybatis-plus语法

下载使用：

go get github.com/zouchangfu/gorm-plus

使用demo：

```

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
	resultDb := mapper.Insert(&test1)
	fmt.Println(resultDb)
	fmt.Println(test1)
}

func TestSaveMigrate(t *testing.T) {
	test1 := Test1{Code: "D2", Price: 100}
	resultDb, err := mapper.InsertMigrate(&test1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultDb)
	fmt.Println(test1)
}

func TestBatchSave(t *testing.T) {
	test1 := Test1{Code: "D3", Price: 100}
	test2 := Test1{Code: "D4", Price: 100}
	resultDb := mapper.InsertBatch(&test1, &test2)
	fmt.Println(resultDb)
	fmt.Println(test1)
	fmt.Println(test2)
}

func TestSaveBatchMigrate(t *testing.T) {
	test1 := Test1{Code: "D5", Price: 100}
	test2 := Test1{Code: "D6", Price: 100}
	resultDb, err := mapper.InsertBatchMigrate(&test1, &test2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultDb)
	fmt.Println(test1)
	fmt.Println(test2)
}

func TestDeleteById(t *testing.T) {
	resultDb := mapper.DeleteById[Test1](1)
	fmt.Println(resultDb)
}

func TestDeleteByIds(t *testing.T) {
	resultDb := mapper.DeleteByIds[Test1](4, 5)
	fmt.Println(resultDb)
}

func TestDelete(t *testing.T) {
	q := mapper.Query[Test1]{}
	q.Eq("code", "D1").Eq("price", 100)
	resultDb := mapper.Delete(&q)
	fmt.Println(resultDb)
}

func TestUpdateById(t *testing.T) {
	test1 := Test1{Code: "777"}
	resultDb := mapper.UpdateById(6, &test1)
	fmt.Println(resultDb)
}

func TestUpdate(t *testing.T) {
	q := mapper.Query[Test1]{}
	q.Eq("code", "D42").Eq("price", 100)
	test1 := Test1{Code: "888"}
	resultDb := mapper.Update(&q, &test1)
	fmt.Println(resultDb)
}

func TestSelectById(t *testing.T) {
	db, result := mapper.SelectById[Test1](1)
	fmt.Println(db)
	fmt.Println(result)
}

func TestSelectByIds(t *testing.T) {
	db, result := mapper.SelectByIds[Test1](1, 2)
	fmt.Println(db)
	fmt.Println(result)
}

func TestSelectOne(t *testing.T) {
	q := mapper.Query[Test1]{}
	q.Eq("code", "D42").Eq("price", 100)
	db, result := mapper.SelectOne(&q)
	fmt.Println(db)
	fmt.Println(result)
}

func TestSelectList(t *testing.T) {
	q := mapper.Query[Test1]{}
	q.Eq("price", 100)
	db, result := mapper.SelectList(&q)
	fmt.Println(db.RowsAffected)
	fmt.Println(result)
}

func TestSelectCount(t *testing.T) {
	q := mapper.Query[Test1]{}
	q.Eq("price", 100)
	db, count := mapper.SelectCount(&q)
	fmt.Println(db)
	fmt.Println(count)
}

```
