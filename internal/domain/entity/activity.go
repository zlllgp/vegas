package entity

import "github.com/zlllgp/vegas/internal/infrastructure/dal/wk/model"

type Activity struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	TenantID int32  `json:"tenant_id"`
}

func NewActivity(options ...ActivityOption) *Activity {
	activity := &Activity{}
	for _, option := range options {
		option(activity)
	}
	return activity
}

func (a *Activity) ToModel() *model.Activity {
	return &model.Activity{
		ID:   a.ID,
		Name: a.Name,
	}
}

func NewActivityWithModel(model *model.Activity) *Activity {
	return &Activity{
		ID:   model.ID,
		Name: model.Name,
	}
}
