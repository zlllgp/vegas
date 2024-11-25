package wire

import (
	"github.com/google/wire"
	"github.com/zlllgp/vegas/internal/application/app"
	"github.com/zlllgp/vegas/internal/application/inf"
	"github.com/zlllgp/vegas/internal/domain/entity"
	"github.com/zlllgp/vegas/internal/domain/repository"
	"github.com/zlllgp/vegas/internal/domain/service"
	"github.com/zlllgp/vegas/internal/infrastructure/cache"
	"github.com/zlllgp/vegas/internal/infrastructure/persistence"
	"github.com/zlllgp/vegas/internal/infrastructure/redis"
	"github.com/zlllgp/vegas/internal/infrastructure/rpc"
)

// https://zhuanlan.zhihu.com/p/691115410
// do wire before use this code then auto generate down code
// step1 add New and append VegasProviderSet
// step2 del wire_gen.go !wireinject
// step3 add function like this
// step4 in wire dir do wire command , then modify main.go

/*
	func InitVegasApp() *app.VegasServiceImpl {
		wire.Build(VegasAppProviderSet)
		return &app.VegasServiceImpl{}
	}
*/
var ActivityRepoProvider = wire.NewSet(
	persistence.NewActivityRepository,
	wire.Bind(new(repository.ActivityRepository), new(*persistence.ActivityRepository)),
)

var RmbRepoProvider = wire.NewSet(
	rpc.NewRmbRepository,
	wire.Bind(new(repository.RmbRepository), new(*rpc.RmbRepository)),
)

var RedisActivityRepoProvider = wire.NewSet(
	redis.NewRedisActivity,
	wire.Bind(new(repository.RedisActivityRepository), new(*redis.RedisActivity)),
)

var CacheRepositoryProvider = wire.NewSet(
	// TODO error
	cache.NewBigCache[entity.Activity],
	wire.Bind(new(repository.CacheRepository[entity.Activity]), new(*cache.BigCache[entity.Activity])),
)

var DrawServiceProvider = wire.NewSet(
	service.NewDrawService,
	wire.Bind(new(inf.DrawService), new(*service.DrawService)),
)

var VegasAppProviderSet = wire.NewSet(
	app.NewVegasServiceImpl,
	ActivityRepoProvider,
	RmbRepoProvider,
	RedisActivityRepoProvider,
	CacheRepositoryProvider,
	DrawServiceProvider,
)
