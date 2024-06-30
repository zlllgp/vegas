package model

import (
	"gorm.io/gorm"
)

const TableNamePlan = "plan"

type Plan struct {
	gorm.Model
	CreatorId   int64  `gorm:"colum:creator_id;not null" json:"creator_id"`
	CreatorName string `gorm:"colum:creator_name;not null" json:"creator_name"`
	Name        string `gorm:"colum:name;not null" json:"name"`
	ActivityId  int64  `gorm:"colum:activity_id;not null" json:"activity_id"`
}

func (p *Plan) TableName() string { return TableNamePlan }
