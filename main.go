package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/zlllgp/vegas/internal/application/config"
	"github.com/zlllgp/vegas/internal/application/vegas"
	"github.com/zlllgp/vegas/internal/infrastructure/dal"
	"github.com/zlllgp/vegas/internal/infrastructure/redis"
	"github.com/zlllgp/vegas/internal/infrastructure/viper"
	"github.com/zlllgp/vegas/kitex_gen/api/vegasservice"
	"github.com/zlllgp/vegas/pkg/consts"
	"github.com/zlllgp/vegas/pkg/mw"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"time"
)

func main() {
	opts := kitexInit()
	// server
	svr := server.NewServer(opts...)
	//实例服务
	vegasImpl := &vegas.VegasServiceImpl{}
	//注册服务
	vegasservice.RegisterService(svr, vegasImpl)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
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
	ip, err := consts.GetOutBoundIP()
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

	// registry
	//r, err := zkregistry.NewZookeeperRegistry(config.GetConf().Registry.RegistryAddress, 40*time.Second)
	//if err != nil {
	//	panic(err)
	//}

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
