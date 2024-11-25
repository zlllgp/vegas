package cache

import (
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

type BigCache[T any] struct {
}

func NewBigCache[T any]() *BigCache[T] {
	return &BigCache[T]{}
}

func (c *BigCache[T]) Get(key string) (T, bool) {
	var value T
	//TODO
	if entry, err := bigCache.Get(key); err == nil {
		json.Unmarshal(entry, &value)
	} else {
		return value, false
	}
	klog.Infof("bigCache stats : %+v", bigCache.Stats())
	return value, true
}
func (c *BigCache[T]) Set(key string, value T, expiration time.Duration) error {
	if entry, err := json.Marshal(value); err == nil {
		bigCache.Set(key, entry)
		return nil
	} else {
		return err
	}
}
func (c *BigCache[T]) Delete(key string) error {
	return nil
}
