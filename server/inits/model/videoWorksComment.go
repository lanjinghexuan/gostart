package model

import "time"

type VideoWorkComment struct {
	Id            int32     `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	WorkId        int16     `gorm:"column:work_id;type:smallint;comment:作品ID;default:NULL;" json:"work_id"`     // 作品ID
	UserId        int16     `gorm:"column:user_id;type:smallint;comment:用户ID;default:NULL;" json:"user_id"`     // 用户ID
	Content       string    `gorm:"column:content;type:varchar(100);comment:评论内容;default:NULL;" json:"content"` // 评论内容
	Tag           int16     `gorm:"column:tag;type:smallint;comment:评论标签表;default:NULL;" json:"tag"`            // 评论标签表
	Pid           int16     `gorm:"column:pid;type:smallint;comment:父级ID;default:NULL;" json:"pid"`             // 父级ID
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;default:NULL;" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime;default:NULL;" json:"updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at;type:datetime;default:NULL;" json:"deleted_at"`
	ReplyCount    int16     `gorm:"column:reply_count;type:smallint;comment:子评论数量;default:0;" json:"reply_count"`       // 子评论数量
	RepliedUserId int32     `gorm:"column:replied_user_id;type:int;comment:被回复用户;default:NULL;" json:"replied_user_id"` // 被回复用户
}
