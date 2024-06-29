package model

import (
	"gorm.io/gorm"
)

const TableNameActivity = "activity"

type Activity struct {
	gorm.Model
	CreatorId   int64  `gorm:"colum:creator_id;not null" json:"creator_id"`
	CreatorName string `gorm:"colum:creator_name;not null" json:"creator_name"`
	Name        string `gorm:"colum:name;not null" json:"name"`
	TenantId    int8   `gorm:"colum:tenant_id;not null" json:"tenant_id"`
}

func (a *Activity) TableName() string { return TableNameActivity }
