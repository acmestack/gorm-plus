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

