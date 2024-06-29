package main

import (
	"github.com/zlllgp/vegas/config"
	"github.com/zlllgp/vegas/internal/dal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	DB, err = gorm.Open(mysql.Open(config.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	g.UseDB(DB)
	g.ApplyBasic(model.Activity{}, model.Plan{}, model.RuleMeta{}, model.RuleInstance{})
	g.ApplyInterface(func() {}, model.Activity{}, model.Plan{}, model.RuleMeta{}, model.RuleInstance{})
	g.Execute()
}
