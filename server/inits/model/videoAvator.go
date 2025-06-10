package model

type VideoAvator struct {
	Id     int32  `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Avator string `gorm:"column:avator;type:varchar(500);default:NULL;" json:"avator"`
}
