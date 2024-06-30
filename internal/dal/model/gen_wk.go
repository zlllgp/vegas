package main

import (
	"github.com/zlllgp/vegas/config"
	"github.com/zlllgp/vegas/internal/dal/model/wk"
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

how to use
gen https://cloud.tencent.com/developer/article/2385346
gen https://cloud.tencent.com/developer/article/2038104
*/
func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "internal/dal/query/wk",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		WithUnitTest: false,
	})

	db, err = gorm.Open(mysql.Open(config.GetConf().MySQLWk.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	g.UseDB(db)
	g.ApplyBasic(wk.Activity{}, wk.Plan{})
	g.ApplyInterface(func() {}, wk.Activity{}, wk.Plan{})
	g.Execute()
}