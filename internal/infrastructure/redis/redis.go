package redis

import "context"

type RedisManager interface {
	GetActivity(c context.Context, activityId int64)
}
