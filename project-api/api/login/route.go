package login

import (
	"log"

	"github.com/gin-gonic/gin"
	"test.com/project-api/router"
)

type RouterLogin struct {
}

func (*RouterLogin) Register(r *gin.Engine) {
	g := r.Group("/project/login")
	g.POST("/getCaptcha", getCaptcha)
	g.POST("/register", registerUser)
}

func init() {
	InitUserRpc()
	log.Println("init login router success")
	router.RegisterRouter(&RouterLogin{})
}
