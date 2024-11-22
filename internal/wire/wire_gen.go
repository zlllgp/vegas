package wire

import (
	"github.com/google/wire"
	"github.com/zlllgp/vegas/internal/application/app"
	"github.com/zlllgp/vegas/internal/domain/repository"
	"github.com/zlllgp/vegas/internal/domain/service"
	"github.com/zlllgp/vegas/internal/infrastructure/persistence"
	"github.com/zlllgp/vegas/internal/infrastructure/rpc"
)

// Injectors from wire_gen.go:

func InitializeVegasImplService() (*app.VegasServiceImpl, error) {
	wire.Build(ProviderSet)
	return &app.VegasServiceImpl{}, nil
}

// wire_gen.go:

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

var ProviderSet = wire.NewSet(
	NewActivityRepo,
	NewRmbRepo,
	NewDrawService,
	NewVegasServiceImpl,
)
