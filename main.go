package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	zkregistry "github.com/kitex-contrib/registry-zookeeper/registry"
	"github.com/zlllgp/vegas/conf"
	"github.com/zlllgp/vegas/internal/dal/mysql"
	"github.com/zlllgp/vegas/internal/dal/redis"
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
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	// registry
	r, err := zkregistry.NewZookeeperRegistry(conf.GetConf().Registry.RegistryAddress, 40*time.Second)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithRegistry(r))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	klog.SetOutput(&lumberjack.Logger{
		Filename:   conf.GetConf().Kitex.LogFileName,
		MaxSize:    conf.GetConf().Kitex.LogMaxSize,
		MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
		MaxAge:     conf.GetConf().Kitex.LogMaxAge,
	})

	// init other
	mysql.Init()
	redis.Init()
	return
}
