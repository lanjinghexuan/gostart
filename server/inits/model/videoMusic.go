package model

type VideoMusic struct {
	Id   int32  `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Name string `gorm:"column:name;type:varchar(50);default:NULL;" json:"name"`
	Path string `gorm:"column:path;type:varchar(200);default:NULL;" json:"path"`
}
