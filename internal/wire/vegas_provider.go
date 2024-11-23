package wire

import (
	"github.com/google/wire"
	"github.com/zlllgp/vegas/internal/application/app"
	"github.com/zlllgp/vegas/internal/domain/repository"
	"github.com/zlllgp/vegas/internal/domain/service"
	"github.com/zlllgp/vegas/internal/infrastructure/persistence"
	"github.com/zlllgp/vegas/internal/infrastructure/rpc"
)

// do wire before use this code then auto generate down code
// step1 add New and append VegasProviderSet
// step2 del wire_gen.go !wireinject
// step3 add function like this
// step4 in wire dir do wire command , then modify main.go

/*func InitializeVegasImplService() (*app.VegasServiceImpl, error) {
	wire.Build(VegasProviderSet)
	return &app.VegasServiceImpl{}, nil
}*/

func NewActivityRepo() repository.ActivityRepository {
	return persistence.NewActivityRepository()
}

func NewRmbRepo() repository.RmbRepository {
	return rpc.NewRmbRepository()
}

func NewDrawService(rmbRep repository.RmbRepository, actRep repository.ActivityRepository) *service.DrawService {
	return service.NewDrawService(rmbRep, actRep)
}

func NewVegasServiceImpl(drawService *service.DrawService) *app.VegasServiceImpl {
	return app.NewVegasServiceImpl(drawService)
}

var VegasProviderSet = wire.NewSet(
	NewActivityRepo,
	NewRmbRepo,
	NewDrawService,
	NewVegasServiceImpl,
)
