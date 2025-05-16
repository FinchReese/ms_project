package project

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-grpc/project_auth"
)

func checkProjectAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := &common.Result{}
		// 检查是否需要进行项目权限校验
		hasProjectCode := false
		hasTaskCode := false
		memberId := c.GetInt64("memberId")

		projectCode := c.PostForm("projectCode")
		if projectCode != "" {
			hasProjectCode = true
		}
		taskCode := c.PostForm("taskCode")
		if taskCode != "" {
			hasTaskCode = true
		}

		if !hasProjectCode && !hasTaskCode { // 没有projectCode和taskCode，则不需要进行项目权限校验
			c.Next()
			return
		}

		// 调用grpc服务获取项目权限信息
		grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
		defer cancel()
		grpcResp, err := ProjectAuthServiceClient.CheckProjectAuth(grpcCtx, &project_auth.CheckProjectAuthReq{
			MemberId:    memberId,
			ProjectCode: projectCode,
			TaskCode:    taskCode,
		})
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusInternalServerError, result.Fail(code, msg))
			c.Abort()
			return
		}
		if !grpcResp.IsMember {
			c.JSON(http.StatusForbidden, result.Fail(http.StatusForbidden, "不是项目成员，无操作权限"))
			c.Abort()
			return
		}
		if grpcResp.IsPrivateProject {
			if !grpcResp.IsOwner {
				c.JSON(http.StatusForbidden, result.Fail(http.StatusForbidden, "私有项目，不是Owner无操作权限"))
				c.Abort()
				return
			}
		}
		// 不是私有项目，普通成员有操作权限
		c.Next()
	}
}
