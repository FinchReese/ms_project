package project

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	common "test.com/project-common"
)

var whiteUriList = []string{
	"project/login/register",
	"project/login",
	"project/login/getCaptcha",
	"project/organization",
	"project/auth/apply",
}

func checkUriAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := &common.Result{}
		uri := c.Request.RequestURI
		// 白名单URL不做权限校验
		for _, whiteUri := range whiteUriList {
			if strings.Contains(uri, whiteUri) {
				c.Next()
				return
			}
		}
		// 根据memberId获取有权限的节点URL列表
		authNodeUrlList, err := GetAuthNodeUrlList(c)
		if err != nil {
			c.JSON(http.StatusForbidden, result.Fail(http.StatusForbidden, "权限校验失败"))
			c.Abort()
			return
		}
		// 判断当前请求的URI是否在有权限的节点URL列表中
		for _, authUrl := range authNodeUrlList {
			if strings.Contains(uri, authUrl) {
				c.Next()
				return
			}
		}
		// 如果当前请求的URI不在有权限的节点URL列表中，则返回401
		c.JSON(http.StatusForbidden, result.Fail(http.StatusForbidden, "权限校验失败"))
		c.Abort()
		return
	}
}
