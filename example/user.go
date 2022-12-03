package example

import "time"

type User struct {
	ID        int64  `gorm:"primaryKey"`
	Username  string `gorm:"column:username"`
	Password  string
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
