package initialize

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"github.com/zlllgp/vegas/config"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.GetConf().Redis.Address,
		Username: config.GetConf().Redis.Username,
		Password: config.GetConf().Redis.Password,
		DB:       config.GetConf().Redis.DB,
	})
	klog.Info("redis init success")
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
