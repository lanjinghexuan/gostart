package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"server/inits/config"
)

var (
	Client  config.Config
	DB      *gorm.DB
	RedisDb *redis.Client
)
