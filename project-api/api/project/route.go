package project

import (
	"github.com/gin-gonic/gin"
	"test.com/project-api/api/midd"
	"test.com/project-api/router"
)

type ProjectRouter struct {
}

func (*ProjectRouter) Register(r *gin.Engine) {
	indexGroup := r.Group("/project/index")
	indexGroup.Use(midd.VerifyToken())
	indexGroup.POST("", index)
	projectGroup := r.Group("/project/project")
	projectGroup.Use(midd.VerifyToken())
	projectGroup.POST("/selfList", selfList)
}

func init() {
	InitUserRpc()
	router.RegisterRouter(&ProjectRouter{})
}
