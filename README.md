# Gorm-plus

Gorm-plus是基于Gorm的增强版，类似Mybatis-plus语法。
## 特性
- [x] 无侵入：只做增强不做改变
- [x] 强大的 CRUD 操作：内置通用查询，不需要任何配置，即可构建复杂条件查询
- [x] 支持 指针字段 形式查询，方便编写各类查询条件，无需再担心字段写错
- [x] 支持 支持主键自动生成
- [x] 内置分页插件



## 预览

```go
type Student struct {
	ID        int
	Name      string
	Age       uint8
	Email     string
	Birthday  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

var gormDb *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println(err)
	}
	gplus.Init(gormDb)
}

func main() {
	var student Student
	// 创建表
	gormDb.AutoMigrate(student)

	// 插入数据
	studentItem := Student{Name: "zhangsan", Age: 18, Email: "123@11.com", Birthday: time.Now()}
	gplus.Insert(&studentItem)

	// 根据Id查询数据
	studentResult, resultDb := gplus.SelectById[Student](studentItem.ID)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
	log.Printf("studentResult:%+v\n", studentResult)

	// 根据条件查询
	query, model := gplus.NewQuery[Student]()
	query.Eq(&model.Name, "zhangsan")
	studentResult, resultDb = gplus.SelectOne(query)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
	log.Printf("studentResult:%+v\n", studentResult)

	// 根据Id更新
	studentItem.Name = "lisi"
	resultDb = gplus.UpdateById[Student](&studentItem)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)

	// 根据条件更新
	query, model = gplus.NewQuery[Student]()
	query.Eq(&model.Name, "lisi").Set(&model.Age, 35)
	resultDb = gplus.Update[Student](query)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)

	// 根据Id删除
	resultDb = gplus.DeleteById[Student](studentItem.ID)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)

	// 根据条件删除
	query, model = gplus.NewQuery[Student]()
	query.Eq(&model.Name, "zhangsan")
	resultDb = gplus.Delete[Student](query)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
}
```

## 使用

### 下载

通过以下命令安装使用：

~~~go
go get github.com/acmestack/gorm-plus
~~~



### 定义表结构

~~~go
type Student struct {
	ID        int
	Name      string
	Age       uint8
	Email     string
	Birthday  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
~~~



### 连接数据库

~~~go
var gormDb *gorm.DB

func init() {
	dsn := "root:root-abcd-1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println(err)
	}
	gplus.Init(gormDb)
}
~~~



### 自动迁移

~~~go
	var student Student
	// 自动迁移
	gormDb.AutoMigrate(student)
~~~



### 基础增删改查

##### 插入一条数据

~~~go
	studentItem := Student{Name: "zhangsan", Age: 18, Email: "123@11.com", Birthday: time.Now()}
	gplus.Insert(&studentItem)
~~~

##### 插入多条数据

~~~go
	student1 := Student{Name: "zhangsan1", Age: 18, Email: "123@11.com", Birthday: time.Now()}
	student2 := Student{Name: "zhangsan2", Age: 18, Email: "123@11.com", Birthday: time.Now()}
	var students = []*Student{&student1, &student2}
	gplus.InsertBatch[Student](students)
~~~

##### 分批插入数据

~~~go
	student1 := Student{Name: "zhangsan1", Age: 18, Email: "123@11.com", Birthday: time.Now()}
	student2 := Student{Name: "zhangsan2", Age: 18, Email: "123@11.com", Birthday: time.Now()}
	student3 := Student{Name: "zhangsan3", Age: 18, Email: "123@11.com", Birthday: time.Now()}
	student4 := Student{Name: "zhangsan4", Age: 18, Email: "123@11.com", Birthday: time.Now()}
	var students = []*Student{&student1, &student2, &student3, &student4}
	// 每次插入2条数据
	gplus.InsertBatchSize[Student](students, 2)
~~~

##### 根据一个ID查询

~~~go
	student, resultDb := gplus.SelectById[Student](2)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
	log.Printf("student:%+v\n", student)
~~~



##### 根据多个ID查询

~~~go
	var ids = []int{2, 3}
	students, resultDb := gplus.SelectByIds[Student](ids)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
	for _, student := range students {
		log.Printf("student:%+v\n", student)
	}
~~~



##### 根据条件查询一条数据
~~~go
	query, model := gplus.NewQuery[Student]()
	query.Eq(&model.Name, "zhangsan")
	student, resultDb := gplus.SelectOne(query)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
	log.Printf("student:%v\n", student)
~~~

##### 根据条件查询多条数据
~~~go
	query, model := gplus.NewQuery[Student]()
	query.Eq(&model.Name, "zhangsan")
	students, resultDb := gplus.SelectList(query)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
	for _, student := range students {
		log.Printf("student:%v\n", student)
	}
~~~

##### 根据条件查询多条数据（泛型封装）
~~~go
	type StudentVo struct {
		Name string
		Age  int
	}
	query, model := gplus.NewQuery[Student]()
	query.Eq(&model.Name, "zhangsan")
	students, resultDb := gplus.SelectListModel[Student, StudentVo](query)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
	for _, student := range students {
		log.Printf("student:%v\n", student)
	}
~~~
##### 根据条件分页查询
~~~go
	query, model := gplus.NewQuery[Student]()
	page := gplus.NewPage[Student](1, 5)
	query.Eq(&model.Name, "zhangsan")
	page, resultDb := gplus.SelectPage(page, query)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
	log.Printf("total:%v\n", page.Total)
	log.Printf("current:%v\n", page.Current)
	log.Printf("size:%v\n", page.Size)
	for _, student := range page.Records {
		log.Printf("student:%v\n", student)
	}
~~~

##### 根据条件分页查询（泛型封装）
~~~go
	type StudentVo struct {
		Name string
		Age  int
	}
	query, model := gplus.NewQuery[Student]()
	page := gplus.NewPage[StudentVo](1, 5)
	query.Eq(&model.Name, "zhangsan")
	page, resultDb := gplus.SelectPageModel[Student, StudentVo](page, query)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
	log.Printf("total:%v\n", page.Total)
	log.Printf("current:%v\n", page.Current)
	log.Printf("size:%v\n", page.Size)
	for _, student := range page.Records {
		log.Printf("student:%v\n", student)
	}
~~~
##### 根据ID更新
~~~go
	student := Student{ID: 3, Name: "lisi"}
	student.Name = "lisi"
	resultDb := gplus.UpdateById[Student](&student)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
~~~
##### 根据条件更新
~~~go
	query, model := gplus.NewQuery[Student]()
	query.Eq(&model.Name, "zhangsan").Set(&model.Age, 30)
	resultDb := gplus.Update[Student](query)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
~~~
##### 根据ID删除
~~~go
	resultDb := gplus.DeleteById[Student](4)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
~~~
##### 根据多个ID删除
~~~go
	var ids = []int{5, 6}
	resultDb := gplus.DeleteByIds[Student](ids)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
~~~
##### 根据条件删除
~~~go
	query, model := gplus.NewQuery[Student]()
	query.Eq(&model.Name, "lisi")
	resultDb := gplus.Delete(query)
	log.Printf("error:%v\n", resultDb.Error)
	log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
~~~