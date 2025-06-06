package model

import "gorm.io/gorm"

type Work struct {
	gorm.Model
	Title          string `gorm:"type:varchar(256);not null;" json:"title"`           //标题
	Describe       string `gorm:"type:varchar(256);not null;" json:"describe"`        //描述
	MusicId        int    `gorm:"type:int(10);not null;" json:"music_id"`             //选择音乐
	Status         string `gorm:"type:varchar(256);not null;" json:"status"`          //审核状态
	WorkCate       string `gorm:"type:varchar(256);not null;" json:"work_cate"`       //作品类型
	Reviewer       string `gorm:"type:varchar(256);not null;" json:"reviewer"`        //审核人
	WorkPermission string `gorm:"type:varchar(256);not null;" json:"work_permission"` //作品权限
	Address        string `gorm:"type:varchar(256);not null;" json:"address"`         //ip地址
	LikeNum        int    `gorm:"type:int(10);not null;" json:"like_num"`             //喜欢数量
	ContentNum     int    `gorm:"type:int(10);not null;" json:"content_num"`          //评论数
	ShareNum       int    `gorm:"type:int(10);not null;" json:"share_num"`            //分享数
	CollectNum     int    `gorm:"type:int(10);not null;" json:"collect_num"`          //收藏数
	BrowseNum      int    `gorm:"type:int(10);not null;" json:"browse_num"`           //浏览量
}
