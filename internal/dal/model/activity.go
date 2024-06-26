package model

import (
	"github.com/zlllgp/vegas/pkg/consts"
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	CreatorId   int64  `json:"creator_id" column:"creator_id"`
	CreatorName string `json:"creator_name" column:"creator_name"`
	Name        string `json:"name" column:"name"`
}

func (a *Activity) TableName() string { return consts.ActivityTableName }
