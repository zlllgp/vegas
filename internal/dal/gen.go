package main

import (
	"github.com/zlllgp/vegas/config"
	"github.com/zlllgp/vegas/internal/dal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

/*
doc
gen https://gorm.io/zh_CN/gen/
gorm https://gorm.io/zh_CN/docs/index.html

how to use
gen https://cloud.tencent.com/developer/article/2385346
gen https://cloud.tencent.com/developer/article/2038104
*/
func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "internal/dal/query",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		WithUnitTest: false,
	})

	db, err = gorm.Open(mysql.Open(config.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	g.UseDB(db)
	g.ApplyBasic(model.Activity{}, model.Plan{}, model.RuleMeta{}, model.RuleInstance{})
	g.ApplyInterface(func() {}, model.Activity{}, model.Plan{}, model.RuleMeta{}, model.RuleInstance{})
	g.Execute()
}
