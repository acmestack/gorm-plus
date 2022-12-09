# gorm-plus
这是一个gorm的增强版，类似mybatis-plus语法。

This is an plus version of gorm, which is similar to the mybatis-plus syntax.

下载使用：

go get github.com/acmestack/gorm-plus

使用demo：

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

select

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

