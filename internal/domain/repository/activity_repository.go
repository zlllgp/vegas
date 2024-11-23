package repository

import (
	"context"
	"github.com/zlllgp/vegas/internal/infrastructure/dal/wk/model"
)

type ActivityRepository interface {
	Create(ctx context.Context, name string) error
	GetById(ctx context.Context, id int64) (*model.Activity, error)
	UpdateName(ctx context.Context, account *model.Activity) error
	IsNameExist(ctx context.Context, name string) (bool, error)
}
