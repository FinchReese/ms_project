package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	_ "test.com/project-api/api"
	"test.com/project-api/config"
	"test.com/project-api/router"
	common "test.com/project-common"
	"test.com/project-common/trace"
)

func main() {
	r := gin.Default()
	// 初始化jaeger追踪器
	tp, err := trace.JaegerTraceProvider(config.AppConf.JaegerConf.CollectorAddr, config.AppConf.ServerConf.Name, "dev")
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	// 使用otelgin中间件
	r.Use(otelgin.Middleware("project-api"))
	router.InitRouter(r)
	r.StaticFS("/upload", http.Dir("upload"))
	common.Run(r, config.AppConf.ServerConf.Name, config.AppConf.ServerConf.Addr, nil)
}
