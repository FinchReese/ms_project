package project

import (
	"github.com/gin-gonic/gin"
	"test.com/project-api/api/midd"
	"test.com/project-api/router"
)

type ProjectRouter struct {
}

func (*ProjectRouter) Register(r *gin.Engine) {
	group := r.Group("/project/index")
	group.Use(midd.VerifyToken())
	group.POST("", index)
}

func init() {
	InitUserRpc()
	router.RegisterRouter(&ProjectRouter{})
}
