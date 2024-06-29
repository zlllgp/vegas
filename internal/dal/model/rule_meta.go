package model

import (
	"gorm.io/gorm"
)

const TableNameRuleMeta = "rule_meta"

type RuleMeta struct {
	gorm.Model
	CreatorId   int64  `gorm:"colum:creator_id;not null" json:"creator_id"`
	CreatorName string `gorm:"colum:creator_name;not null" json:"creator_name"`
	Name        string `gorm:"colum:name;not null" json:"name"`
	Description string `gorm:"colum:description;not null" json:"description"`
}

func (r *RuleMeta) TableName() string { return TableNameRuleMeta }
