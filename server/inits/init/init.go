package init

import (
	"fmt"
	"github.com/spf13/viper"
	"server/global"
	"server/inits/mysql"
	"server/inits/redis"
)

func init() {
	InitConfig()
	mysql.InitMysql()
	redis.InitRedis()
}

func InitConfig() {

	viper.SetConfigFile("../inits/config/dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}
	err = viper.Unmarshal(&global.CONFIG)
	fmt.Printf("global.CONFIG:%v\n", global.CONFIG)
}
