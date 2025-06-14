package model

type VideoLike struct {
	Id     int32 `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	UserId int32 `gorm:"column:user_id;type:int;default:NULL;" json:"user_id"`
	LikeId int32 `gorm:"column:like_id;type:int;default:NULL;" json:"like_id"`
}
