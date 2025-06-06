package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"server/global"
)

func RedisInit() {
	appConfig := global.Client.Redis
	global.RedisDb = redis.NewClient(&redis.Options{
		Addr:     appConfig.Host,
		Password: appConfig.Password,
		DB:       appConfig.Db,
	})
	stc := context.Background()
	err := global.RedisDb.Ping(stc).Err()
	if err != nil {
		log.Println("redis连接失败")
		return
	} else {
		log.Println("redis连接成功")
	}
}
