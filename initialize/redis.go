package initialize

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"github.com/zlllgp/vegas/config"
)

var rdb *redis.Client

/*
go get github.com/redis/go-redis/v9
doc：https://redis.uptrace.dev/zh/guide/go-redis.html
*/
func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.GetConf().Redis.Address,
		Username: config.GetConf().Redis.Username,
		Password: config.GetConf().Redis.Password,
		DB:       config.GetConf().Redis.DB,
	})
	klog.Info("redis init success")
	//todo redis can't connect don't panic
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
