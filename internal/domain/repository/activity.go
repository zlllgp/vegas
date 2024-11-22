package repository

import (
	"context"
	"github.com/zlllgp/vegas/internal/infrastructure/dal/wk/model"
)

type ActivityRepository interface {
	GetById(ctx context.Context, id int64) (*model.Activity, error)
}
