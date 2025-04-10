package project

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-grpc/project"
)

const (
	indexServiceTimeOut = 2 // index接口的超时时间设置为2秒
)

func index(ctx *gin.Context) {
	result := &common.Result{}
	grpcCtx, cancel := context.WithTimeout(context.Background(), indexServiceTimeOut*time.Second)
	defer cancel()
	indexResponse, err := projectServiceClient.Index(grpcCtx, &project.IndexMessage{})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
	}
	ctx.JSON(http.StatusOK, result.Success(indexResponse.Menus))
}
