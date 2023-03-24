# Gorm-plus

基于 Gorm-plus，详见[https://github.com/acmestack/gorm-plus]()

Gorm-plus是基于Gorm的增强版，类似Mybatis-plus语法。

## 特性

- [x] 无侵入，只做增强不做改变
- [x] 强大的CRUD 操作，内置通用查询，不需要任何配置，即可构建复杂条件查询
- [x] 支持指针字段形式查询，方便编写各类查询条件，无需再担心字段写错
- [x] 支持主键自动生成
- [x] 内置分页插件

## 事务
- 在需要事务的场景，可以使用gplus.Begin()开启事务，获取*gorm.DB
- 所有的dao方法，均支持传入*gorm.DB，后续的操作均以传入的为准

## 预览

```Go
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

### 定义表结构

```Go
type Student struct {
    ID        int
    Name      string
    Age       uint8
    Email     string
    Birthday  time.Time
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### 连接数据库

和gorm的连接数据库的方式是一样的，只是多了一行 `gplus.Init(gormDb)`，通过它来初始化gorm-plus的连接。

```Go
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
```

### 自动迁移

```Go
var student Student
gormDb.AutoMigrate(student)
```

### 基础增删改查

#### 插入一条数据

```Go
studentItem := Student{Name: "zhangsan", Age: 18, Email: "123@11.com", Birthday: time.Now()}
gplus.Insert(&studentItem)
```

#### 插入多条数据

```Go
student1 := Student{Name: "zhangsan1", Age: 18, Email: "123@11.com", Birthday: time.Now()}
student2 := Student{Name: "zhangsan2", Age: 18, Email: "123@11.com", Birthday: time.Now()}
var students = []*Student{&student1, &student2}
gplus.InsertBatch[Student](students)
```

#### 分批插入数据

```Go
student1 := Student{Name: "zhangsan1", Age: 18, Email: "123@11.com", Birthday: time.Now()}
student2 := Student{Name: "zhangsan2", Age: 18, Email: "123@11.com", Birthday: time.Now()}
student3 := Student{Name: "zhangsan3", Age: 18, Email: "123@11.com", Birthday: time.Now()}
student4 := Student{Name: "zhangsan4", Age: 18, Email: "123@11.com", Birthday: time.Now()}
var students = []*Student{&student1, &student2, &student3, &student4}
// 每次插入2条数据
gplus.InsertBatchSize[Student](students, 2)
```

#### 根据一个ID查询

```Go
student, resultDb := gplus.SelectById[Student](2)
log.Printf("error:%v\n", resultDb.Error)
log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
log.Printf("student:%+v\n", student)
```

#### 根据多个ID查询

```Go
var ids = []int{2, 3}
students, resultDb := gplus.SelectByIds[Student](ids)
log.Printf("error:%v\n", resultDb.Error)
log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
for _, student := range students {
    log.Printf("student:%+v\n", student)
}
```

#### 根据条件查询一条数据

```Go
query, model := gplus.NewQuery[Student]()
query.Eq(&model.Name, "zhangsan")
student, resultDb := gplus.SelectOne(query)
log.Printf("error:%v\n", resultDb.Error)
log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
log.Printf("student:%v\n", student)
```

#### 根据条件查询多条数据

```Go
query, model := gplus.NewQuery[Student]()
query.Eq(&model.Name, "zhangsan")
students, resultDb := gplus.SelectList(query)
log.Printf("error:%v\n", resultDb.Error)
log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
for _, student := range students {
    log.Printf("student:%v\n", student)
}
```

#### 根据条件查询多条数据（泛型封装）

有时候我们可能需要额外使用其他类型的接收数据，gplus也支持这种方式，通过传入第二个泛型参数来实现。

```Go
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
```

#### 根据条件分页查询

```Go
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
```

#### 根据条件分页查询（泛型封装）

```Go
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
```

#### 根据ID更新

```Go
student := Student{ID: 3, Name: "lisi"}
student.Name = "lisi"
resultDb := gplus.UpdateById[Student](&student)
log.Printf("error:%v\n", resultDb.Error)
log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
```

#### 根据条件更新

```Go
query, model := gplus.NewQuery[Student]()
query.Eq(&model.Name, "zhangsan").Set(&model.Age, 30)
resultDb := gplus.Update[Student](query)
log.Printf("error:%v\n", resultDb.Error)
log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
```

#### 根据ID删除

```Go
resultDb := gplus.DeleteById[Student](4)
log.Printf("error:%v\n", resultDb.Error)
log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
```

#### 根据多个ID删除

```Go
var ids = []int{5, 6}
resultDb := gplus.DeleteByIds[Student](ids)
log.Printf("error:%v\n", resultDb.Error)
log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
```

#### 根据条件删除

```Go
query, model := gplus.NewQuery[Student]()
query.Eq(&model.Name, "lisi")
resultDb := gplus.Delete(query)
log.Printf("error:%v\n", resultDb.Error)
log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
```

### 高级查询

#### 条件构造器

gorm-plus 提供了强大的条件构造器,通过构造器能够组合不同的查询条件。

|方法名|操作|
|-|-|
|Eq|等于 |
|Ne|不等于|
|Gt|大于|
|Ge|大于等于|
|Lt|小于|
|Le|小于等于|
|Like|LIKE '%值%'|
|NotLike|NOT LIKE '%值%'|
|LikeLeft|LIKE '%值'|
|LikeRight|LIKE '值%'|
|IsNull|字段 IS NULL|
|IsNotNull|字段 IS NOT NULL|
|In|字段 IN (值1，值2)|
|NotIn|字段 NOT IN (值1，值2)|
|Between|字段 BETWEEN 值1 ADN 值2|
|NotBetween|字段 NOT BETWEEN 值1 ADN 值2|


#### 指针字段查询

通过NewQuery函数返回的第二个参数model，我们可以使用指针的方式来指定要使用的字段，无需再担心字段名写错。

```Go
query, model := gplus.NewQuery[Student]()
query.Eq(&model.Name, "zhangsan1")
studentResult, resultDb := gplus.SelectOne(query)
log.Printf("error:%v\n", resultDb.Error)
log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
log.Printf("studentResult:%+v\n", studentResult)
```

#### Query泛型简化

如果不希望每次创建Query对象的时候携带上泛型，我们可以提供一个全局的泛型Dao。

```Go
var dao gplus.Dao[Student]
func main() {
    query, model := dao.NewQuery()
    query.Eq(&model.Name, "zhangsan")
    list, resultDb := gplus.SelectList(query)
    fmt.Println(resultDb.RowsAffected)
    for _, v := range list {
        marshal, _ := json.Marshal(v)
        fmt.Println(string(marshal))
    }
}
```

我们也可以把`gplus.Dao`组合到我们自己定义的Dao对象中

```Go
type StudentDao struct {
    gplus.Dao[Student]
}
var studentDao StudentDao
func main() {
    query, model := studentDao.NewQuery()
    query.Eq(&model.Name, "zhangsan")
    list, resultDb := gplus.SelectList(query)
    fmt.Println(resultDb.RowsAffected)
    for _, v := range list {
        marshal, _ := json.Marshal(v)
        fmt.Println(string(marshal))
    }
}
```

#### 查询指定字段

通过Select来设置待查询的字段

```Go
query, model := gplus.NewQuery[Student]()
query.Eq(&model.Name, "zhangsan1").Select(&model.Name)
studentResult, resultDb := gplus.SelectOne(query)
log.Printf("error:%v\n", resultDb.Error)
log.Printf("RowsAffected:%v\n", resultDb.RowsAffected)
log.Printf("studentResult:%+v\n", studentResult)
```



#### 排序

```Go
query, model := gplus.NewQuery[Student]()
query.Eq(&model.Name, "zhangsan1").OrderByAsc(&model.Age)
students, resultDb := gplus.SelectList(query)
log.Printf("error:%+v", resultDb.Error)
fmt.Println("RowsAffected:", resultDb.RowsAffected)
for _, student := range students {
    log.Printf("student:%+v", student)
}
```

#### 事务

```Go
// 开启事务
tx := gplus.Begin()

// 使用defer，实现遇到错误时回滚
defer tx.Rollback()

// 新增，传入tx
student := Student{Name: "zhangsan", Age: 18, Email: "123@11.com", Birthday: time.Now()}
result := gplus.Insert(&student, tx)
...

// 更新，传入tx
query, model := gplus.NewQuery[Student]()
query.Eq(&model.Name, "zhangsan").Set(&model.Age, 30)
resultDb := gplus.Update[Student](query, tx)
...

// 查询，如果需要在事务中查询，则也需要传入tx
resultStudent, db := gplus.SelectById[Student](student.ID, tx)
...

// 提交事务，否则数据库数据不更改
tx.Commit()
```