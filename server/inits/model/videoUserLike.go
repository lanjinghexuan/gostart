package model

type VideoUserLike struct {
	Id      int32 `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	VideoId int32 `gorm:"column:video_id;type:int;default:NULL;" json:"video_id"`
	UserId  int32 `gorm:"column:user_id;type:int;default:NULL;" json:"user_id"`
}
