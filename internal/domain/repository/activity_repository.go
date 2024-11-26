package repository

import (
	"context"
	"github.com/zlllgp/vegas/internal/infrastructure/dal/wk/model"
)

//go:generate mockgen -source=activity_repository.go -destination=mock_repository/activity_repository_mock.go -package=mock_repository ActivityRepository
type ActivityRepository interface {
	Create(ctx context.Context, name string) error
	GetById(ctx context.Context, id int64) (*model.Activity, error)
	UpdateName(ctx context.Context, account *model.Activity) error
	IsNameExist(ctx context.Context, name string) (bool, error)
}
