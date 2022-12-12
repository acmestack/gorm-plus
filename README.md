# gorm-plus
这是一个gorm的增强版，类似mybatis-plus语法。

This is an plus version of gorm, which is similar to the mybatis-plus syntax.

## 下载：

go get github.com/acmestack/gorm-plus

go install github.com/acmestack/gorm-plus/cmd/gplus@latest

## 生成工具gplus

### 为什么有需要 gplus？

我们在使用gorm的时候，是需要手写字段名称的，例如这样：

```go
gormDb.Where("username = ? age = ?","zhangsan",18)
```

一旦名称长，非常容易误写，而且如果有字段名称修改的话，还需要全局搜索一个个地修改，比较麻烦。

Go没有提供类似Java的lambad表达式或者C#中 `nameof `方式直接获取某个对象的字段名称的操作，但是我们可以通过生成代码的方式生成字段名。

所有就有了gplus，它作用就是自动识别结构体，把结构体的字段名生成出来。



### 使用gplus

通过 `gplus gen paths=路径`，gplus 会自动识别带有`// +gplus:column=true`注释的结构体，给这个结构体生成字段。

gplus 会在输入的路径下面生成 `zz_gen.column.go`文件。

例如：

在example目录下创建了了一个users.go 目录，执行 `gplus en paths=./eample`

users.go

```go

// +gplus:column=true

type User struct {
	ID        int64
	Username  string `gorm:"column:username"`
	Password  string
	Address   string
	Age       int
	Phone     string
	Score     int
	Dept      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
```



zz_gen.column.go （自动生成的）

~~~go
var UserColumn = struct {
	ID        string
	Username  string
	Password  string
	Address   string
	Age       string
	Phone     string
	Score     string
	Dept      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Username:  "username",
	Password:  "password",
	Address:   "address",
	Age:       "age",
	Phone:     "phone",
	Score:     "score",
	Dept:      "dept",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}
~~~



## 基础操作：

users 表

~~~sql
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
                          `id` int(0) NOT NULL AUTO_INCREMENT,
                          `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `age` int(0) NULL DEFAULT NULL,
                          `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `score` int(0) NULL DEFAULT NULL,
                          `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `dept` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `created_at` datetime(0) NULL DEFAULT NULL,
                          `updated_at` datetime(0) NULL DEFAULT NULL,
                          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
~~~



### 初始化

```go

var GormDb *gorm.DB

func init() {
	dsn := "root:root-abcd-1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	GormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}
	gplus.Init(GormDb)
}


// +gplus:column=true

type User struct {
	ID        int64
	Username  string `gorm:"column:username"`
	Password  string
	Address   string
	Age       int
	Phone     string
	Score     int
	Dept      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "Users"
}

```



### Select

```go

func TestSelectById(t *testing.T) {
	user, resultDb := gplus.SelectById[User](1)
	if resultDb.Error != nil {
		if errors.Is(resultDb.Error, gorm.ErrRecordNotFound) {
			log.Fatalln("SelectById Data not found:", resultDb.Error)
		}
		log.Fatalln("SelectById error:", resultDb.Error)
	}
	log.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(user)
	log.Println(string(marshal))
}

func TestSelectByIds(t *testing.T) {
	var ids []int
	ids = append(ids, 1)
	ids = append(ids, 2)
	users, resultDb := gplus.SelectByIds[User](ids)
	if resultDb.Error != nil {
		log.Fatalln(resultDb.Error)
	}
	log.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(users)
	log.Println(string(marshal))
}

func TestSelectOne1(t *testing.T) {
	q := gplus.NewQuery[User]()
	q.Eq(UserColumn.Username, "zhangsan1")
	user, resultDb := gplus.SelectOne(q)

	if resultDb.Error != nil {
		if errors.Is(resultDb.Error, gorm.ErrRecordNotFound) {
			log.Fatalln("SelectOne Data not found:", resultDb.Error)
		}
		log.Fatalln("SelectOne error:", resultDb.Error)
	}

	log.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(user)
	log.Println(string(marshal))
}

func TestSelectOne2(t *testing.T) {
	q := gplus.NewQuery[User]()
	q.Eq(UserColumn.Username, "zhangsan").
		Select(UserColumn.Username, UserColumn.Password)
	user, resultDb := gplus.SelectOne(q)

	if resultDb.Error != nil {
		if errors.Is(resultDb.Error, gorm.ErrRecordNotFound) {
			log.Fatalln("SelectOne Data not found:", resultDb.Error)
		}
		log.Fatalln("SelectOne error:", resultDb.Error)
	}

	log.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(user)
	log.Println(string(marshal))
}

func TestSelectList(t *testing.T) {
	q := gplus.NewQuery[User]()
	q.Eq(UserColumn.Username, "zhangsan")
	users, resultDb := gplus.SelectList(q)
	if resultDb.Error != nil {
		log.Fatalln("error:", resultDb.Error)
	}
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		log.Println(string(marshal))
	}
}

func TestSelectBracketList(t *testing.T) {
	q := gplus.NewQuery[User]()
	bracketQuery := gplus.NewQuery[User]()
	bracketQuery.Eq(UserColumn.Address, "上海").Or().Eq(UserColumn.Address, "北京")

	q.Eq(UserColumn.Username, "zhangsan").AndBracket(bracketQuery)
	users, resultDb := gplus.SelectList(q)
	if resultDb.Error != nil {
		log.Fatalln("error:", resultDb.Error)
	}
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		log.Println(string(marshal))
	}
}

func TestSelectTableList(t *testing.T) {
	type deptCount struct {
		Dept  string
		Count string
	}
	q := gplus.NewQuery[User]()
	q.Group(UserColumn.Dept).Select(UserColumn.Dept, "count(*) as count")
	users, resultDb := gplus.SelectListModel[User, deptCount](q)
	if resultDb.Error != nil {
		log.Fatalln("error:", resultDb.Error)
	}
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		log.Println(string(marshal))
	}
}

func TestSelectPage(t *testing.T) {
	q := gplus.NewQuery[User]()
	q.Eq(UserColumn.Age, 18)
	page := gplus.NewPage[User](1, 10)
	pageResult, resultDb := gplus.SelectPage(page, q)
	if resultDb.Error != nil {
		log.Fatalln("error:", resultDb.Error)
	}
	log.Println("total", pageResult.Total)
	for _, u := range pageResult.Records {
		marshal, _ := json.Marshal(u)
		log.Println(string(marshal))
	}
}

func TestSelectTablePage(t *testing.T) {
	type deptCount struct {
		Dept  string
		Count string
	}
	q := gplus.NewQuery[User]()
	q.Group(UserColumn.Dept).Select(UserColumn.Dept, "count(*) as count")
	page := gplus.NewPage[deptCount](1, 2)
	pageResult, resultDb := gplus.SelectPageModel[User, deptCount](page, q)
	if resultDb.Error != nil {
		log.Fatalln("error:", resultDb.Error)
	}
	log.Println("total:", pageResult.Total)
	for _, u := range pageResult.Records {
		marshal, _ := json.Marshal(u)
		log.Println(string(marshal))
	}
}

func TestSelectCount(t *testing.T) {
	q := gplus.NewQuery[User]()
	q.Eq(UserColumn.Age, 18)
	count, resultDb := gplus.SelectCount(q)
	if resultDb.Error != nil {
		log.Fatalln("error:", resultDb.Error)
	}
	log.Println("count:", count)
}
```



### Insert

```go

func TestInsert(t *testing.T) {
	user := &User{Username: "zhangsan", Password: "123456", Age: 18, Score: 100, Dept: "A部门"}
	result := gplus.Insert(user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
}

func TestInsertBatch(t *testing.T) {
	user1 := &User{Username: "zhangsan1", Password: "123456", Age: 18, Score: 12, Dept: "导弹部门"}
	user2 := &User{Username: "lisi", Password: "123456", Age: 16, Score: 34, Dept: "投诉部门"}
	user3 := &User{Username: "wangwu", Password: "123456", Age: 26, Score: 33, Dept: "研发部门"}
	user4 := &User{Username: "zhangsan4", Password: "123456", Age: 30, Score: 11, Dept: "产品部门"}
	user5 := &User{Username: "zhangsan5", Password: "123456", Age: 12, Score: 34, Dept: "产品部门1"}
	user6 := &User{Username: "zhangsan6", Password: "123456", Age: 45, Score: 123, Dept: "产品部门12"}

	var users []*User
	users = append(users, user1)
	users = append(users, user2)
	users = append(users, user3)
	users = append(users, user4)
	users = append(users, user5)
	users = append(users, user6)

	result := gplus.InsertBatch[User](users)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		fmt.Println(string(marshal))
	}
}

func TestInsertBatchSize(t *testing.T) {
	user1 := &User{Username: "zhangsan1", Password: "123456", Age: 18, Score: 12, Dept: "导弹部门"}
	user2 := &User{Username: "lisi", Password: "123456", Age: 16, Score: 34, Dept: "投诉部门"}
	user3 := &User{Username: "wangwu", Password: "123456", Age: 26, Score: 33, Dept: "研发部门"}
	user4 := &User{Username: "zhangsan4", Password: "123456", Age: 30, Score: 11, Dept: "产品部门"}
	user5 := &User{Username: "zhangsan5", Password: "123456", Age: 12, Score: 34, Dept: "产品部门1"}
	user6 := &User{Username: "zhangsan6", Password: "123456", Age: 45, Score: 123, Dept: "产品部门12"}

	var users []*User
	users = append(users, user1)
	users = append(users, user2)
	users = append(users, user3)
	users = append(users, user4)
	users = append(users, user5)
	users = append(users, user6)

	result := gplus.InsertBatchSize[User](users, 3)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		fmt.Println(string(marshal))
	}
}
```



### Update

```go

func TestUpdateById(t *testing.T) {
	user := &User{ID: 1, Username: "zhangsan", Password: "123456", Age: 18, Score: 100, Dept: "A部门asdfasdf"}
	result := gplus.UpdateById(user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
}

func TestUpdate(t *testing.T) {
	q := gplus.NewQuery[User]()
	q.Eq(UserColumn.Username, "zhangsan").Set(UserColumn.Dept, "相关部门123123").
		Set(UserColumn.Phone, 12312)
	result := gplus.Update(q)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
}

```



### Delete

~~~go

func TestDeleteById(t *testing.T) {
	result := gplus.DeleteById[User](13)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
}

func TestDelete(t *testing.T) {
	q := gplus.NewQuery[User]()
	q.Ge(UserColumn.Age, 50)
	result := gplus.Delete(q)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
}
~~~



## 通用操作

您可以在结构体中嵌入`gplus.CommonDao`获得通用的操作

### 初始化

~~~go
var GormDb *gorm.DB

func init() {
	dsn := "root:root-abcd-1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	GormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}
	gplus.Init(GormDb)
}

// +gplus:column=true

type User struct {
	ID        int64
	Username  string `gorm:"column:username"`
	Password  string
	Address   string
	Age       int
	Phone     string
	Score     int
	Dept      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "Users"
}


var userDao = NewUserDao[User]()

type UserDao[T any] struct {
	gplus.CommonDao[T]
}

func NewUserDao[T any]() *UserDao[T] {
	return &UserDao[T]{}
}

~~~



### 查询

~~~go

func TestGetById(t *testing.T) {
	user, resultDb := userDao.GetById(2)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(user)
	fmt.Println("user1:", string(marshal))
}

func TestGetByOne(t *testing.T) {
	query := gplus.NewQuery[User]()
	query.Eq(UserColumn.Username, "zhangsan1")
	user, resultDb := userDao.GetOne(query)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(user)
	fmt.Println("user1:", string(marshal))
}

func TestListAll(t *testing.T) {
	users, resultDb := userDao.ListAll()
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, v := range users {
		marshal, _ := json.Marshal(v)
		fmt.Println("u:", string(marshal))
	}
}

func TestList(t *testing.T) {
	query := gplus.NewQuery[User]()
	query.Eq(UserColumn.Username, "zhangsan1")
	users, resultDb := userDao.List(query)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, v := range users {
		marshal, _ := json.Marshal(v)
		fmt.Println("u:", string(marshal))
	}
}

func TestPageAll(t *testing.T) {
	page := gplus.NewPage[User](1, 2)
	page, resultDb := userDao.PageAll(page)
	fmt.Println("page total:", page.Total)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, v := range page.Records {
		marshal, _ := json.Marshal(v)
		fmt.Println("u:", string(marshal))
	}
}

func TestPage(t *testing.T) {
	page := gplus.NewPage[User](1, 2)
	query := gplus.NewQuery[User]()
	query.Eq(UserColumn.Username, "zhangsan1")
	page, resultDb := userDao.Page(page, query)
	fmt.Println("page total:", page.Total)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, v := range page.Records {
		marshal, _ := json.Marshal(v)
		fmt.Println("u:", string(marshal))
	}
}
~~~

### 保存

```go

func TestSave(t *testing.T) {
	user1 := &User{Username: "zhangsan1", Password: "123456", Age: 18, Score: 12, Dept: "导弹部门"}
	resultDb := userDao.Save(user1)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	marshal, _ := json.Marshal(user1)
	fmt.Println("user1:", string(marshal))
}

func TestSaveBatch(t *testing.T) {
	user1 := &User{Username: "zhangsan1", Password: "123456", Age: 11, Score: 12, Dept: "导弹部门"}
	user2 := &User{Username: "zhangsan1", Password: "123456", Age: 13, Score: 12, Dept: "导弹部门"}
	var users []*User
	users = append(users, user1)
	users = append(users, user2)
	resultDb := userDao.SaveBatch(users)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		fmt.Println("user:", string(marshal))
	}
}

func TestSaveBatchSize(t *testing.T) {
	user1 := &User{Username: "zhangsan1", Password: "123456", Age: 11, Score: 12, Dept: "导弹部门"}
	user2 := &User{Username: "zhangsan1", Password: "123456", Age: 13, Score: 12, Dept: "导弹部门"}
	user3 := &User{Username: "zhangsan1", Password: "123456", Age: 13, Score: 12, Dept: "导弹部门"}
	user4 := &User{Username: "zhangsan1", Password: "123456", Age: 13, Score: 12, Dept: "导弹部门"}
	var users []*User
	users = append(users, user1)
	users = append(users, user2)
	users = append(users, user3)
	users = append(users, user4)
	resultDb := userDao.SaveBatchSize(users, 2)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
	for _, u := range users {
		marshal, _ := json.Marshal(u)
		fmt.Println("user:", string(marshal))
	}
}
```



### 更新

```go

func TestUpdateById(t *testing.T) {
	user4 := &User{ID: 4, Username: "zhangsan666", Password: "123456", Age: 13, Score: 12, Dept: "导弹部门"}
	userDao.UpdateById(user4)
}

func TestUpdate(t *testing.T) {
	query := gplus.NewQuery[User]()
	query.Eq(UserColumn.Username, "zhangsan1").Set(UserColumn.Age, 50)
	userDao.Update(query)
}
```



### 删除

```go

func TestRemoveById(t *testing.T) {
	resultDb := userDao.RemoveById(7)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
}

func TestRemoveByIds(t *testing.T) {
	var ids []int
	ids = append(ids, 5)
	ids = append(ids, 6)
	resultDb := userDao.RemoveByIds(ids)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
}

func TestRemove(t *testing.T) {
	query := gplus.NewQuery[User]()
	query.Eq(UserColumn.Username, "lisi")
	resultDb := userDao.Remove(query)
	fmt.Println("RowsAffected:", resultDb.RowsAffected)
}
```

