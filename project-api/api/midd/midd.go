package midd

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"test.com/project-api/api/rpc"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-grpc/user/login"
)

const (
	verifyTokenTimeout = 2 // 校验token超时时间设置为2秒
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取header的token
		token := c.GetHeader("Authorization")
		// 调用grpc接口校验token
		ctx, cancelFunc := context.WithTimeout(context.Background(), verifyTokenTimeout*time.Second)
		defer cancelFunc()
		resp, err := rpc.LoginServiceClient.VerifyToken(ctx, &login.VerifyTokenReq{Token: token})

		// 根据校验token结果处理
		result := &common.Result{}
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusUnauthorized, result.Fail(code, msg))
			c.Abort()
			return
		}
		c.Set("memberId", resp.Member.Id)
		c.Set("memberName", resp.Member.Name)
		// 继续处理请求
		c.Next()
	}
}
