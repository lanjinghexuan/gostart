package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"server/inits/config"
)

var (
	CONFIG config.Config
	DB     *gorm.DB
	REDIS  *redis.Client
	Ctx    = context.Background()
)
