package cache

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/bluele/gcache"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dtm-labs/rockscache"
	"github.com/zlllgp/vegas/internal/infrastructure/redis"
	"log"
	"time"
)

var (
	bigCache *bigcache.BigCache
	gCache   gcache.Cache
	RCache   *rockscache.Client
)

func InitBigCache() {
	config := bigcache.Config{
		Shards:             2,
		LifeWindow:         10 * time.Minute,
		CleanWindow:        5 * time.Minute,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
		OnRemove:           nil,
		OnRemoveWithReason: nil,
	}

	cache, initErr := bigcache.New(context.Background(), config)
	if initErr != nil {
		log.Fatal(initErr)
	}
	klog.Infof("bigCache init success stats : %+v", cache.Stats())
}

func InitGCache() {
	gCache = gcache.New(200).
		LRU().
		Build()
}

func InitRocksCache() {
	RCache = rockscache.NewClient(redis.GetRdb(), rockscache.NewDefaultOptions())
	klog.Infof("rCache init success stats")
}
