package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"server/global"
)

func InitRedis() {
	global.REDIS = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.CONFIG.Redis.Host, global.CONFIG.Redis.Port),
	})

	global.REDIS.Ping(global.Ctx)

	global.REDIS.Set(global.Ctx, "a", 1, 60)
	fmt.Println(global.REDIS.Get(global.Ctx, "a"))

}
