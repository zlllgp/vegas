package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	zkregistry "github.com/kitex-contrib/registry-zookeeper/registry"
	"github.com/zlllgp/vegas/config"
	"github.com/zlllgp/vegas/initialize"
	"github.com/zlllgp/vegas/kitex_gen/api/vegas"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"time"
)

func main() {
	opts := kitexInit()
	// server
	svr := server.NewServer(opts...)
	//实例服务
	vegasImpl := &VegasImpl{}
	//注册服务
	vegas.RegisterService(svr, vegasImpl)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", config.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: config.GetConf().Kitex.Service,
	}))

	// registry
	r, err := zkregistry.NewZookeeperRegistry(config.GetConf().Registry.RegistryAddress, 40*time.Second)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithRegistry(r))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(config.LogLevel())
	klog.SetOutput(&lumberjack.Logger{
		Filename:   config.GetConf().Kitex.LogFileName,
		MaxSize:    config.GetConf().Kitex.LogMaxSize,
		MaxBackups: config.GetConf().Kitex.LogMaxBackups,
		MaxAge:     config.GetConf().Kitex.LogMaxAge,
	})

	// init other
	initialize.InitDB()
	initialize.InitRedis()
	return
}
