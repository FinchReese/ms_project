package project

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	model_project "test.com/project-api/pkg/model/project"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-grpc/project"
)

const (
	serviceTimeOut = 2 // 服务超时时间设置为2秒
)

func index(ctx *gin.Context) {
	result := &common.Result{}
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()
	indexResponse, err := projectServiceClient.Index(grpcCtx, &project.IndexMessage{})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}
	var resp []*model_project.Menu
	copier.Copy(&resp, indexResponse.Menus)
	ctx.JSON(http.StatusOK, result.Success(resp))
}

func selfList(ctx *gin.Context) {
	result := &common.Result{}
	var req model_project.SelfListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	val, _ := ctx.Get("memberId")
	memberId := val.(int64)
	grpcCtx, cancelFunc := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancelFunc()
	memberName := ctx.GetString("memberName")
	resp, err := projectServiceClient.GetProjectList(grpcCtx, &project.GetProjectListReq{MemberId: memberId, MemberName: memberName, Page: req.Page, Size: req.PageSize})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}
	if resp.ProjectList == nil {
		resp.ProjectList = []*project.ProjectMemberMessage{}
	}
	var projectList []*model_project.ProjectAndMember
	copier.Copy(&projectList, resp.ProjectList)

	ctx.JSON(http.StatusOK, result.Success(gin.H{
		"list":  projectList,
		"total": resp.Total,
	}))
}
