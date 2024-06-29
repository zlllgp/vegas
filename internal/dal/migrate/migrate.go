package main

import (
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	"github.com/zlllgp/vegas/config"
	"github.com/zlllgp/vegas/internal/dal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
model to table
*/
func main() {
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               config.GetConf().MySQL.DSN,
		DefaultStringSize: 256, // string 类型字段的默认长度
	}),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.Activity{}, &model.Plan{}, &model.RuleMeta{}, &model.RuleInstance{})
	log.Info("Auto Migration Completed")
}
