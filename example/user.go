package example

import (
	"time"
)

// +gplus:column=true

type User struct {
	ID        int64  `gorm:"primaryKey"`
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
