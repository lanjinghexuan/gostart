package model

type VideoObject struct {
	Id   int32  `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Path string `gorm:"column:path;type:varchar(500);default:NULL;" json:"path"`
}
