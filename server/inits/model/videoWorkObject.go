package model

import "time"

type VideoWorkObject struct {
	Id        int32     `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	WorkId    int16     `gorm:"column:work_id;type:smallint;default:NULL;" json:"work_id"`
	ObjectId  int16     `gorm:"column:object_id;type:smallint;comment:作品对象ID;default:NULL;" json:"object_id"` // 作品对象ID
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	CreatedBy uint64    `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;default:NULL;" json:"created_by"` // 创建者
	UpdatedBy uint64    `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;default:NULL;" json:"updated_by"` // 更新者
	DeletedBy uint64    `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;default:NULL;" json:"deleted_by"` // 删除者
}
