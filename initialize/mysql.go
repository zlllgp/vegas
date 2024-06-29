package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/zlllgp/vegas/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDB() {
	DB, err = gorm.Open(mysql.Open(config.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	klog.Info("mysql init success")
	if err != nil {
		panic(err)
	}
}
