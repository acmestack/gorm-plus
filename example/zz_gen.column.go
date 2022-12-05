package example

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
var ClazzColumn = struct {
	ID          string
	Name        string
	Count       string
	TeacherName string
	HouseNumber string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	Name:        "name",
	Count:       "count",
	TeacherName: "teacher_name",
	HouseNumber: "house",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}
