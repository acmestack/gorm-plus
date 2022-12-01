package gormplus

import (
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	GormDb = db
}
