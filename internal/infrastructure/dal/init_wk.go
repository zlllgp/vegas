package dal

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/sirupsen/logrus"
	"github.com/zlllgp/vegas/internal/application/config"
	"github.com/zlllgp/vegas/internal/infrastructure/dal/wk/query"
	"gorm.io/gorm/logger"
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
	// 创建日志文件
	logFile, err := os.OpenFile("./logs/sql.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}
	//defer logFile.Close()
	logrus.SetOutput(logFile)

	newLogger := logger.New(
		//log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logrus.StandardLogger(),
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	dbWk, errWk = gorm.Open(mysql.New(mysql.Config{
		DSN:               config.GlobalConfig.MySQLWk.DSN,
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

	query.SetDefault(dbWk)
}
