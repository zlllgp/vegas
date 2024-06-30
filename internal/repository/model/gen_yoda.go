package main

import (
	"github.com/zlllgp/vegas/config"
	"github.com/zlllgp/vegas/internal/repository/model/yoda"
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
func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "internal/repository/query/yoda",
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
	g.ApplyBasic(model.UerRight{})
	g.ApplyInterface(func() {}, model.UerRight{})
	g.Execute()
}
