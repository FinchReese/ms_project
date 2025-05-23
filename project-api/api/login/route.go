package login

import (
	"log"

	"github.com/gin-gonic/gin"
	"test.com/project-api/api/midd"
	"test.com/project-api/api/rpc"
	"test.com/project-api/router"
)

type RouterLogin struct {
}

func (*RouterLogin) Register(r *gin.Engine) {
	g := r.Group("/project/login")
	g.POST("/getCaptcha", getCaptcha)
	g.POST("/register", registerUser)
	r.POST("/project/login", projectLogin)
	org := r.Group("/project/organization")
	org.Use(midd.VerifyToken())
	org.POST("/_getOrgList", getOrgList)
}

func init() {
	rpc.InitUserRpc()
	log.Println("init login router success")
	router.RegisterRouter(&RouterLogin{})
}
