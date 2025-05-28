package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	common "test.com/project-common"
	"test.com/project-common/trace"
	"test.com/project-project/config"
	"test.com/project-project/internal/kafka_log"
	"test.com/project-project/router"
)

func main() {
	//从配置中读取日志配置，初始化日志
	config.AppConf.InitZapLog()
	r := gin.Default()
	// 初始化jaeger追踪器
	tp, err := trace.JaegerTraceProvider(config.AppConf.ServerConf.Name, "dev")
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	// 使用otelgin中间件
	r.Use(otelgin.Middleware("project-project"))
	router.InitRouter(r)
	//grpc服务注册到etcd
	router.RegisterGrpcAddrConf()
	// 启动grpc服务器
	grpcServer := router.RegisterGrpc()
	// 初始化kafka日志写入器
	writerStopFunc := kafka_log.InitKafkaWriter()
	// 初始化kafka缓存读取器
	cacheReader := kafka_log.NewKafkaCache()
	go cacheReader.DeleteCache()
	stop := func() {
		grpcServer.Stop()
		writerStopFunc()
		cacheReader.Close()
	}
	common.Run(r, config.AppConf.ServerConf.Name, config.AppConf.ServerConf.Addr, stop)
}
