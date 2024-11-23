package repository

import (
	"context"
	"github.com/zlllgp/vegas/internal/infrastructure/dal/yoda/model"
	"github.com/zlllgp/vegas/internal/infrastructure/persistence"
	"gorm.io/gen"
)

type RightRepository interface {
	GetById(ctx context.Context, userId int64) (*model.UserRight, error)
	Query(ctx context.Context, offset, limit int, conditions ...gen.Condition) ([]*model.UserRight, error)
}

var _ RightRepository = (*persistence.RightRepository)(nil)
