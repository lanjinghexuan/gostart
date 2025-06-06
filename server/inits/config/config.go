package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Mysql struct {
		User     string
		Password string
		Host     string
		Port     int
		Database string
	}
	Redis struct {
		Host     string
		Password string
		Db       int
	}
}
type Nacos struct {
	NamespaceId string
	IpAddr      string
	Port        int
	DataId      string
	Group       string
}

var AppConf Nacos

func GetConfig() {
	viper.SetConfigFile("D:\\goWork\\src\\video\\server\\inits\\config\\dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("配置文件读取失败")
		return
	}
	err = viper.Unmarshal(&AppConf)
	if err != nil {
		log.Println("配置文件解析失败")
		return
	}
	log.Println("配置文件读取成功")
	log.Println(AppConf)
}
