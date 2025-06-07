package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/global"
)

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.CONFIG.Mysql.User,
		global.CONFIG.Mysql.Pass,
		global.CONFIG.Mysql.Host,
		global.CONFIG.Mysql.Port,
		global.CONFIG.Mysql.DB,
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("链接mysql出错.error:", err)
	}
	fmt.Println("mysql链接成功.")
}
