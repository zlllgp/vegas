package mysql

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/zlllgp/vegas/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
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
