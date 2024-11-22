package main

import (
	"github.com/zlllgp/vegas/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	dbWk  *gorm.DB
	errWk error
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
		OutPath:      "internal/dal/wk/query",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		WithUnitTest: false,
	})

	dbWk, errWk = gorm.Open(mysql.Open(config.GetConf().MySQLWk.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	g.UseDB(dbWk)
	//g.ApplyBasic(model.Activity{}, model.Plan{})
	//g.ApplyInterface(func() {}, model.Activity{}, model.Plan{})

	g.ApplyBasic(g.GenerateModel("activity"))
	g.ApplyBasic(g.GenerateModel("plan"))
	g.ApplyBasic(g.GenerateModel("rule_meta"))
	g.Execute()
}
