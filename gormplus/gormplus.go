package gormplus

import (
	"github.com/zouchangfu/gorm-plus/mapper"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	mapper.GormDb = db
}
