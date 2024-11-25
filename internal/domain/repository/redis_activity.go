package repository

import (
	"context"
	"github.com/zlllgp/vegas/internal/domain/entity"
)

type RedisActivityRepository interface {
	StoreActivity(ctx context.Context, id int64, activity *entity.Activity) error
	GetActivity(ctx context.Context, id int64) (*entity.Activity, error)
}
