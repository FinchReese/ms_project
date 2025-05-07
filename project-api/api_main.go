package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "test.com/project-api/api"
	"test.com/project-api/config"
	"test.com/project-api/router"
	common "test.com/project-common"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.StaticFS("/upload", http.Dir("upload"))
	common.Run(r, config.AppConf.ServerConf.Name, config.AppConf.ServerConf.Addr, nil)
}
