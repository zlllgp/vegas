package model

import (
	"gorm.io/gorm"
)

const TableNameRuleInstance = "rule_instance"

type RuleInstance struct {
	gorm.Model
	CreatorId   int64  `gorm:"colum:creator_id;not null" json:"creator_id"`
	CreatorName string `gorm:"colum:creator_name;not null" json:"creator_name"`
	RuleMetaId  int64  `gorm:"colum:rule_meta_id;not null" json:"rule_meta_id"`
}

func (r *RuleInstance) TableName() string { return TableNameRuleInstance }
