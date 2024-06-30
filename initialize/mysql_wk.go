package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/zlllgp/vegas/config"
	"github.com/zlllgp/vegas/internal/dal/query/wk"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbWk  *gorm.DB
	errWk error
)

/*
gorm https://gorm.io/zh_CN/docs/index.html
*/
func InitWkDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	dbWk, errWk = gorm.Open(mysql.New(mysql.Config{
		DSN:               config.GetConf().MySQLWk.DSN,
		DefaultStringSize: 256, // string 类型字段的默认长度
	}),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			Logger:                 newLogger,
		},
	)
	/*
		sqlDB, err := db.DB()
		// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)*/

	klog.Info("mysql wk init success")
	if errWk != nil {
		panic(errWk)
	}

	wk.SetDefault(dbWk)
}
