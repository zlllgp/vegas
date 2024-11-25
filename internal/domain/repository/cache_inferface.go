package repository

import "time"

type CacheInterface interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, expiration time.Duration) error
	Delete(key string) error
}
