package model

import (
	"github.com/zlllgp/vegas/pkg/consts"
	"gorm.io/gorm"
)

type RuleInstance struct {
	gorm.Model
	CreatorId   int64  `json:"creator_id" column:"creator_id"`
	CreatorName string `json:"creator_name" column:"creator_id"`
	RuleMetaId  int64  `json:"rule_meta_id" column:"rule_meta_id"`
	Name        string `json:"name" column:"name"`
	Description string `json:"description" column:"description"`
}

func (r *RuleInstance) TableName() string { return consts.RuleInstanceTableName }
