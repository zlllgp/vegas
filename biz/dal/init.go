package dal

import (
	"github.com/zlllgp/vegas/biz/dal/mysql"
	"github.com/zlllgp/vegas/biz/dal/redis"
)

func Init() {
	mysql.Init()
	redis.Init()
}
