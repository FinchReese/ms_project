package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	common "test.com/project-common"
	"test.com/project-common/trace"
	"test.com/project-user/config"
	"test.com/project-user/router"
)

func main() {
	//从配置中读取日志配置，初始化日志
	config.AppConf.InitZapLog()
	r := gin.Default()
	// 初始化jaeger追踪器
	tp, err := trace.JaegerTraceProvider(config.AppConf.JaegerConf.CollectorAddr, config.AppConf.ServerConf.Name, "dev")
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	// 使用otelgin中间件
	r.Use(otelgin.Middleware("project-user"))
	router.InitRouter(r)
	//grpc服务注册到etcd
	router.RegisterGrpcAddrConf()
	grpcServer := router.RegisterGrpc()
	stop := func() {
		grpcServer.Stop()
	}
	common.Run(r, config.AppConf.ServerConf.Name, config.AppConf.ServerConf.Addr, stop)
}
