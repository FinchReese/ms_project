package login

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/project-api/router"
)

type RouterLogin struct {
}

func (*RouterLogin) Register(r *gin.Engine) {
	g := r.Group("/project/login")
	g.POST("/getCaptcha", GetCaptcha)
}

func init() {
	InitUserRpc()
	log.Println("init login router success")
	router.RegisterRouter(&RouterLogin{})
}
