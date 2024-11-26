package repository

import "time"

//go:generate mockgen -source=cache_repository.go -destination=mock_repository/cache_repository_mock.go -package=mock_repository CacheRepository
type CacheRepository interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, expiration time.Duration) error
	Delete(key string) error
}
