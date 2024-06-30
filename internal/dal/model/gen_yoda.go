package main

import (
	"github.com/zlllgp/vegas/config"
	"github.com/zlllgp/vegas/internal/dal/model/yoda"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

/*
doc
gen https://gorm.io/zh_CN/gen/

how to use
gen https://cloud.tencent.com/developer/article/2385346
gen https://cloud.tencent.com/developer/article/2038104
*/
func main1() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "internal/dal/query/yoda",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		WithUnitTest: false,
	})

	db, err = gorm.Open(mysql.Open(config.GetConf().MySQLYoda.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	g.UseDB(db)
	g.ApplyBasic(yoda.UerRight{})
	g.ApplyInterface(func() {}, yoda.UerRight{})
	g.Execute()
}
