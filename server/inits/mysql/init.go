package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"server/global"
	"server/inits/model"
)

func MysqlInit() {
	var err error
	appConfig := global.Client.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", appConfig.User, appConfig.Password, appConfig.Host, appConfig.Port, appConfig.Database)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("数据库连接失败")
		return
	} else {
		log.Println("数据库连接成功")
	}
	global.DB.AutoMigrate(&model.User{}, &model.Work{})
}
