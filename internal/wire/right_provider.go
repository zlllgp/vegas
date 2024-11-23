package wire

import (
	"github.com/google/wire"
	"github.com/zlllgp/vegas/internal/application/app"
	"github.com/zlllgp/vegas/internal/domain/repository"
	"github.com/zlllgp/vegas/internal/domain/service"
	"github.com/zlllgp/vegas/internal/infrastructure/persistence"
)

func NewRightRepo() repository.RightRepository {
	return persistence.NewRightRepository()
}

func NewRightService(rightRep repository.RightRepository) *service.RightService {
	return service.NewRightService(rightRep)
}

func NewRightServiceImpl(rightService *service.RightService) *app.RightServiceImpl {
	return app.NewRightServiceImpl(rightService)
}

var RightProviderSet = wire.NewSet(
	NewRightRepo,
	NewRightService,
	NewRightServiceImpl,
)
