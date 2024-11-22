package wire

import (
	"github.com/google/wire"
	"github.com/zlllgp/vegas/internal/application/app"
	"github.com/zlllgp/vegas/internal/domain/repository"
	"github.com/zlllgp/vegas/internal/domain/service"
	"github.com/zlllgp/vegas/internal/infrastructure/persistence"
	"github.com/zlllgp/vegas/internal/infrastructure/rpc"
)

var ActivityRepoProviderSet = wire.NewSet(
	persistence.NewActivityRepository,
	wire.Bind(new(repository.ActivityRepository), new(*persistence.ActivityRepository)),
)

var RmbRepoProviderSet = wire.NewSet(
	rpc.NewRmbRepository,
	wire.Bind(new(repository.RmbRepository), new(*rpc.RmbRepository)),
)

var DrawServiceProviderSet = wire.NewSet(
	service.NewDrawService,
	ActivityRepoProviderSet,
	RmbRepoProviderSet,
)

var VegasServiceImplProviderSet = wire.NewSet(
	app.NewVegasServiceImpl,
	DrawServiceProviderSet,
)

func InitializeVegasServiceImpl() *app.VegasServiceImpl {
	wire.Build(
		ActivityRepoProviderSet,
		RmbRepoProviderSet,
		DrawServiceProviderSet,
		VegasServiceImplProviderSet,
	)
	return &app.VegasServiceImpl{}
}
