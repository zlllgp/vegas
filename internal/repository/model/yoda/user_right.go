package model

import (
	"gorm.io/gorm"
)

const TableNameUerRight = "user_right"

type UerRight struct {
	gorm.Model
	UserId           int64  `gorm:"colum:user_id;not null" json:"user_id"`
	UserName         string `gorm:"colum:user_name;not null" json:"user_name"`
	ActivityId       int64  `gorm:"colum:activity_id;not null" json:"activity_id"`
	PlanId           int64  `gorm:"colum:plan_id;not null" json:"plan_id"`
	RightsTemplateId int64  `gorm:"colum:rights_template_id;not null" json:"rights_template_id"`
	RightsInstanceId int64  `gorm:"colum:rights_instance_id;not null" json:"rights_instance_id"`
}

func (r *UerRight) TableName() string { return TableNameUerRight }
