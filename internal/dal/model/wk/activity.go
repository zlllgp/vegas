package wk

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

const TableNameActivity = "activity"

type Activity struct {
	gorm.Model
	CreatorId   int64                 `gorm:"colum:creator_id;not null" json:"creator_id"`
	CreatorName string                `gorm:"colum:creator_name;not null" json:"creator_name"`
	Name        string                `gorm:"colum:name;not null" json:"name"`
	TenantId    int8                  `gorm:"colum:tenant_id;not null" json:"tenant_id"`
	IsDel       soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt;not null;default:0"`
}

func (a *Activity) TableName() string { return TableNameActivity }
