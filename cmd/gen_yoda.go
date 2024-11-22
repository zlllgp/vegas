package main

import (
	"github.com/zlllgp/vegas/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	dbYoda  *gorm.DB
	errYoda error
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
		OutPath:      "internal/dal/yoda/query",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		WithUnitTest: false,
	})

	dbYoda, errYoda = gorm.Open(mysql.Open(config.GetConf().MySQLYoda.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	g.UseDB(dbYoda)
	//g.ApplyBasic(model.UerRight{})
	//g.ApplyInterface(func() {}, model.UerRight{})
	g.ApplyBasic(g.GenerateModel("user_right"))
	g.Execute()
}
