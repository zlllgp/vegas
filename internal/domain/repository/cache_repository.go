package repository

import "time"

type CacheRepository[T any] interface {
	Get(key string) (T, bool)
	Set(key string, value T, expiration time.Duration) error
	Delete(key string) error
}
