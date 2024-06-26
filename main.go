package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	zkregistry "github.com/kitex-contrib/registry-zookeeper/registry"
	"github.com/zlllgp/vegas/kitex_gen/api/vegas"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	//log file
	f, err := os.OpenFile("logs/application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//log klog
	klog.SetLevel(klog.LevelInfo)
	klog.SetOutput(f)
	klog.Infof("main las run")

	//zookeeper
	r, err := zkregistry.NewZookeeperRegistry([]string{"192.168.3.147:2181"}, 40*time.Second)
	if err != nil {
		panic(err)
	}

	// server
	svr := server.NewServer(server.WithRegistry(r), server.WithServiceAddr(&net.TCPAddr{IP: net.ParseIP("192.168.3.71"), Port: 8888}), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "vegas"}))
	//实例服务
	vegasImpl := &VegasImpl{}
	//注册服务
	vegas.RegisterService(svr, vegasImpl)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
