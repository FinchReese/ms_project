package main

import (
	"github.com/gin-gonic/gin"
	common "test.com/project-common"
	"test.com/project-project/config"
	"test.com/project-project/router"
)

func main() {
	//从配置中读取日志配置，初始化日志
	config.AppConf.InitZapLog()
	r := gin.Default()
	router.InitRouter(r)
	//grpc服务注册到etcd
	router.RegisterGrpcAddrConf()
	// 启动grpc服务器
	grpcServer := router.RegisterGrpc()
	stop := func() {
		grpcServer.Stop()
	}
	common.Run(r, config.AppConf.ServerConf.Name, config.AppConf.ServerConf.Addr, stop)
}
