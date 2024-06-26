package model

import (
	"github.com/zlllgp/vegas/pkg/consts"
	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	CreatorId   int64  `json:"creator_id" column:"creator_id"`
	CreatorName string `json:"creator_name" column:"creator_name"`
	ActivityId  int64  `json:"activity_id" column:"activity_id"`
	Name        string `json:"name" column:"name"`
}

func (p *Plan) TableName() string { return consts.PlanTableName }
