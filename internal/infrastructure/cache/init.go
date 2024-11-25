package cache

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/bluele/gcache"
	"github.com/cloudwego/kitex/pkg/klog"
	"log"
	"time"
)

var (
	bigCache *bigcache.BigCache
	gCache   gcache.Cache
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
