package wire

import (
	"github.com/google/wire"
	"github.com/zlllgp/vegas/internal/application/app"
	"github.com/zlllgp/vegas/internal/domain/repository"
	"github.com/zlllgp/vegas/internal/domain/service"
	"github.com/zlllgp/vegas/internal/infrastructure/persistence"
)

// https://zhuanlan.zhihu.com/p/691115410
// do wire before use this code then auto generate down code
// step1 add New and append VegasProviderSet
// step2 del wire_gen.go !wireinject
// step3 add function like this
// step4 in wire dir do wire command , then modify main.go

/*
	func InitRightApp() *app.RightServiceImpl {
		wire.Build(RightAppProviderSet)
		return &app.RightServiceImpl{}
	}
*/
var RightRepoProvider = wire.NewSet(
	persistence.NewRightRepository,
	wire.Bind(new(repository.RightRepository), new(*persistence.RightRepository)),
)

var RightServiceProvider = wire.NewSet(
	service.NewRightServiceImpl,
	wire.Bind(new(service.RightService), new(*service.RightServiceImpl)),
)

var RightAppProviderSet = wire.NewSet(
	app.NewRightServiceImpl,
	RightRepoProvider,
	RightServiceProvider,
)
