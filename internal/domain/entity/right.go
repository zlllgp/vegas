package entity

import (
	"github.com/zlllgp/vegas/internal/infrastructure/dal/yoda/model"
	"github.com/zlllgp/vegas/kitex_gen/api"
)

type Right struct {
	ID       int64
	UserName string
	Amt      string
}

func (a *Right) ToModel() *model.UserRight {
	return &model.UserRight{
		ID:       a.ID,
		UserName: a.UserName,
	}
}

func NewWithModel(model *model.UserRight) *Right {
	return &Right{
		ID:       model.ID,
		UserName: model.UserName,
	}
}

func (t *Right) ToDTO() *api.Right {
	return &api.Right{
		Id: t.ID,
	}
}
