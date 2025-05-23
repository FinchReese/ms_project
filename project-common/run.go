package common

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(router *gin.Engine, serverName string, addr string, stop func()) {
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	//保证下面的优雅启停
	go func() {
		log.Printf("%s running in %s \n", serverName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal)
	//SIGINT 用户发送INTR字符(Ctrl+C)触发
	//SIGTERM 结束程序(可以被捕获、阻塞或忽略)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutting Down %s ...\n", serverName)
	if stop != nil {
		stop()
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("%s Shutdown, cause by: %v\n", serverName, err)
	}
	log.Printf("%s stop success...\n", serverName)
}
