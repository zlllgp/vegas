package redis

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/zlllgp/vegas/internal/domain/entity"
	"github.com/zlllgp/vegas/pkg/errno"
	"strconv"
)

type RedisRepository struct {
}

func NewRedisRepository() *RedisRepository {
	return &RedisRepository{}
}

func (s *RedisRepository) StoreActivity(ctx context.Context, id int64, activity *entity.Activity) error {
	aJson, err := json.Marshal(activity)
	if err != nil {
		return err
	}
	err = rdb.Set(ctx, "activity:"+strconv.FormatInt(id, 10), aJson, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *RedisRepository) GetActivity(ctx context.Context, id int64) (*entity.Activity, error) {
	result, err := rdb.Get(ctx, "activity:"+strconv.FormatInt(id, 10)).Result()
	if errors.Is(err, redis.Nil) {
		return nil, errno.ActivityNotFoundErr
	}
	if err != nil {
		return nil, err
	}
	var activity entity.Activity
	err = json.Unmarshal([]byte(result), &activity)
	if err != nil {
		return nil, err
	}
	return &activity, nil
}
