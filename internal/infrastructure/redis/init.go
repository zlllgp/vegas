package redis

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"github.com/zlllgp/vegas/internal/application/config"
)

var rdb *redis.Client

/*
go get github.com/redis/go-redis/v9
doc：https://redis.uptrace.dev/zh/guide/go-redis.html
*/
func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.GlobalConfig.Redis.Address,
		Username: config.GlobalConfig.Redis.Username,
		Password: config.GlobalConfig.Redis.Password,
		DB:       config.GlobalConfig.Redis.DB,
	})
	klog.Info("redis init success")
	//todo redis can't connect don't panic
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
