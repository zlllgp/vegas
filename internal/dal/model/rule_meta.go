package model

import (
	"github.com/zlllgp/vegas/pkg/consts"
	"gorm.io/gorm"
)

type RuleMeta struct {
	gorm.Model
	CreatorId   int64  `json:"creator_id" column:"creator_id"`
	CreatorName string `json:"creator_name" column:"creator_id"`
	Name        string `json:"name" column:"name"`
	Description string `json:"description" column:"description"`
}

func (r *RuleMeta) TableName() string { return consts.RuleMetaTableName }
