package cache

import (
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

type GCache struct {
}

func NewGCache() *GCache {
	return &GCache{}
}

func (c *GCache) Get(key string) (interface{}, bool) {
	if entry, err := gCache.Get(key); err == nil {
		klog.Infof("gCache hitRate : %+v", gCache.HitRate())
		return entry, true
	} else {
		klog.Errorf("[cache] get cache failed, key: %s, error: %v", key, err)
		return nil, false
	}
}
func (c *GCache) Set(key string, value interface{}, expiration time.Duration) error {
	if entry, err := json.Marshal(value); err == nil {
		gCache.Set(key, entry)
		return nil
	} else {
		return err
	}
}
func (c *GCache) Delete(key string) error {
	return nil
}
