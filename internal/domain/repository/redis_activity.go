package repository

import (
	"context"
	"github.com/zlllgp/vegas/internal/domain/entity"
)

//go:generate mockgen -source=redis_activity.go -destination=mock_repository/redis_activity_mock.go -package=mock_repository RedisActivityRepository
type RedisActivityRepository interface {
	StoreActivity(ctx context.Context, id int64, activity *entity.Activity) error
	GetActivity(ctx context.Context, id int64) (*entity.Activity, error)
}
