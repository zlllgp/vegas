package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/zlllgp/vegas/internal/application/app"
	"github.com/zlllgp/vegas/internal/application/config"
	"github.com/zlllgp/vegas/internal/domain/service"
	"github.com/zlllgp/vegas/internal/infrastructure/dal"
	"github.com/zlllgp/vegas/internal/infrastructure/mw"
	"github.com/zlllgp/vegas/internal/infrastructure/persistence"
	"github.com/zlllgp/vegas/internal/infrastructure/redis"
	"github.com/zlllgp/vegas/internal/infrastructure/rpc"
	"github.com/zlllgp/vegas/internal/infrastructure/utils"
	"github.com/zlllgp/vegas/internal/infrastructure/viper"
	"github.com/zlllgp/vegas/kitex_gen/api/rightservice"
	"github.com/zlllgp/vegas/kitex_gen/api/vegasservice"
	"go.uber.org/dig"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"time"
)

func main() {
	opts := kitexInit()

	svr := server.NewServer(opts...)
	// construct
	/*	rmbRep := rpc.NewRmbRepository()
		actRep := persistence.NewActivityRepository()
		drawService := service.NewDrawService(rmbRep, actRep)
		vegasImpl := app.NewVegasServiceImpl(drawService)*/

	// one: use wire
	/*	vegasImpl := wire.InitVegasApp()
		vegasservice.RegisterService(svr, vegasImpl)
		rightImpl := wire.InitRightApp()
		rightservice.RegisterService(svr, rightImpl)*/
	// two: use dig
	c := digInit()
	initService(c, svr)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func initService(c *dig.Container, svr server.Server) {
	var vegasService app.VegasServiceImpl
	c.Invoke(func(vegasServiceImpl app.VegasServiceImpl) {
		vegasService = vegasServiceImpl
	})
	vegasservice.RegisterService(svr, &vegasService)

	var rightService app.RightServiceImpl
	c.Invoke(func(rightServiceImpl app.RightServiceImpl) {
		rightService = rightServiceImpl
	})
	rightservice.RegisterService(svr, &rightService)
}

func digInit() *dig.Container {
	c := dig.New()
	c.Provide(persistence.NewRightRepository)
	c.Provide(service.NewRightService)
	c.Provide(app.NewRightServiceImpl)

	c.Provide(persistence.NewActivityRepository)
	c.Provide(rpc.NewRmbRepository)
	c.Provide(service.NewDrawService)
	c.Provide(app.NewVegasServiceImpl)
	return c
}

func kitexInit() (opts []server.Option) {
	// config
	viper.Init()

	// log
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(viper.LogLevel())
	klog.SetOutput(&lumberjack.Logger{
		Filename:   config.GlobalConfig.Kitex.LogFileName,
		MaxSize:    config.GlobalConfig.Kitex.LogMaxSize,
		MaxBackups: config.GlobalConfig.Kitex.LogMaxBackups,
		MaxAge:     config.GlobalConfig.Kitex.LogMaxAge,
	})

	// address
	ip, err := utils.GetOutBoundIP()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", ip+":8080")
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: config.GlobalConfig.Kitex.Service,
	}))

	//etcd
	r, err := etcd.NewEtcdRegistry([]string{config.GlobalConfig.Etcd.Address})
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithRegistry(r))

	//middleware
	opts = append(opts, server.WithMiddleware(mw.CommonMiddleware), server.WithMiddleware(mw.ServerMiddleware))

	//timeout
	opts = append(opts, server.WithReadWriteTimeout(3*time.Second))

	//limit
	opts = append(opts, server.WithLimit(&limit.Option{MaxConnections: 10000, MaxQPS: 10000}))

	// init other
	dal.InitWkDB()
	dal.InitYodaDB()
	redis.InitRedis()
	return
}
