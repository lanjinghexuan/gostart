package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone    string `gorm:"type:char(11);not null;" json:"phone"`
	Password string `gorm:"type:varchar(256);not null;" json:"password"`
}
type UserCenter struct {
	gorm.Model
	Name          string `gorm:"type:varchar(256);not null;" json:"name"`          //名称
	NickName      string `gorm:"type:varchar(256);not null;" json:"nickName"`      //昵称
	UserCode      string `gorm:"type:varchar(256);not null;" json:"userCode"`      //编号
	Signature     string `gorm:"type:varchar(256);not null;" json:"signature"`     //签名
	Sex           string `gorm:"type:varchar(256);not null;" json:"sex"`           //性别
	IpAddress     string `gorm:"type:varchar(256);not null;" json:"ipAddress"`     //ip地址
	Constellation string `gorm:"type:varchar(256);not null;" json:"constellation"` //星座
	AttendCount   int64  `gorm:"type:int(10);not null;" json:"attendCount"`        //关注数
	FanCount      int64  `gorm:"type:int(10);not null;" json:"fanCount"`           //粉丝数
	ZanCount      int64  `gorm:"type:int(10);not null;" json:"zanCount"`           //点赞数
	Status        string `gorm:"type:varchar(256);not null;" json:"status"`        //用户状态
	ImageId       int64  `gorm:"type:int(10);not null;" json:"imageId"`            //头像
	Info          string `gorm:"type:varchar(256);not null;" json:"info"`          //认证信息
	Mobile        string `gorm:"type:varchar(256);not null;" json:"mobile"`        //手机号
	RealNameAuth  string `gorm:"type:varchar(256);not null;" json:"realNameAuth"`  //实名认证状态
	Age           int64  `gorm:"type:int(10);not null;" json:"age"`                //年龄
	OnlineStatus  string `gorm:"type:varchar(256);not null;" json:"onlineStatus"`  //在线状态
	Type          string `gorm:"type:varchar(256);not null;" json:"type"`          //认证类型
	Level         string `gorm:"type:varchar(256);not null;" json:"level"`         //用户等级
	Balance       string `gorm:"type:varchar(256);not null;" json:"balance"`       //余额
}
