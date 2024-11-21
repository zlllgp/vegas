package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/zlllgp/vegas/config"
	"github.com/zlllgp/vegas/internal/dal"
	"github.com/zlllgp/vegas/internal/redis"
	"github.com/zlllgp/vegas/kitex_gen/api/vegas"
	"github.com/zlllgp/vegas/pkg/consts"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
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
	ip, err := consts.GetOutBoundIP()
	if err != nil {
		panic(err)
	}

	// address
	addr, err := net.ResolveTCPAddr("tcp", ip+":8080")
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: config.GetConf().Kitex.Service,
	}))

	// registry
	//r, err := zkregistry.NewZookeeperRegistry(config.GetConf().Registry.RegistryAddress, 40*time.Second)
	//if err != nil {
	//	panic(err)
	//}

	//etcd
	r, err := etcd.NewEtcdRegistry([]string{config.GetConf().Etcd.Address})
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
	dal.InitWkDB()
	dal.InitYodaDB()
	redis.InitRedis()
	return
}
