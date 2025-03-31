package router

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	Register(r *gin.Engine)
}

var routers = []Router{}

func RegisterRouter(ro ...Router) {
	routers = append(routers, ro...)
}

func InitRouter(r *gin.Engine) {
	for _, router := range routers {
		router.Register(r)
	}
}
