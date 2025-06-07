package global

import (
	"gorm.io/gorm"
	"server/inits/config"
)

var (
	CONFIG config.Config
	DB     *gorm.DB
)
