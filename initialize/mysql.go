package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/zlllgp/vegas/config"
	"github.com/zlllgp/vegas/internal/dal/query"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDB() {
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               config.GetConf().MySQL.DSN,
		DefaultStringSize: 256, // string 类型字段的默认长度
	}),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	klog.Info("mysql init success")
	if err != nil {
		panic(err)
	}

	query.SetDefault(DB)
}
