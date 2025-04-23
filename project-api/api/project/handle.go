package project

import (
	"context"
	"fmt"
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
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	val, _ := ctx.Get("memberId")
	memberId := val.(int64)
	grpcCtx, cancelFunc := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancelFunc()
	memberName := ctx.GetString("memberName")
	resp, err := projectServiceClient.GetProjectList(grpcCtx, &project.GetProjectListReq{
		MemberId:   memberId,
		MemberName: memberName,
		SelectBy:   req.SelectBy,
		Page:       req.Page,
		Size:       req.PageSize,
	})
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

func projectTemplate(ctx *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	memberId := ctx.GetInt64("memberId")
	var req model_project.GetProjectTemplateReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	// 调用grpc接口
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()
	grpcReq := &project.GetProjectTemplatesReq{
		MemberId: memberId,
		Page:     req.Page,
		PageSize: req.PageSize,
		ViewType: req.ViewType,
	}
	grpcResp, err := projectServiceClient.GetProjectTemplates(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}
	var templates []*model_project.ProjectTemplate
	copier.Copy(&templates, grpcResp.Ptm)
	if templates == nil {
		templates = []*model_project.ProjectTemplate{}
	}
	for _, template := range templates {
		if template.TaskStages == nil {
			template.TaskStages = []*model_project.TaskStagesOnlyName{}
		}
	}
	ctx.JSON(http.StatusOK, result.Success(gin.H{
		"list":  templates,
		"total": grpcResp.Total,
	}))
}

func saveProject(ctx *gin.Context) {
	result := &common.Result{}
	// 获取参数
	memberId := ctx.GetInt64("memberId")
	var req model_project.SaveProjectReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	// 调用grpc接口
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()
	grpcReq := &project.SaveProjectReq{
		MemberId:     memberId,
		TemplateCode: req.TemplateCode,
		Name:         req.Name,
		Description:  req.Description,
	}
	grpcResp, err := projectServiceClient.SaveProject(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}

	resp := &model_project.SaveProjectResp{}
	copier.Copy(&resp, grpcResp)
	ctx.JSON(http.StatusOK, result.Success(resp))
}

func getProjectInfo(ctx *gin.Context) {
	fmt.Println("call getProjectInfo")
	result := &common.Result{}
	// 获取参数
	memberId := ctx.GetInt64("memberId")
	projectCode := ctx.PostForm("projectCode")
	// 通过grpc接口查询
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()
	grpcReq := &project.GetProjectDetailReq{MemberId: memberId, ProjectCode: projectCode}
	grpcResp, err := projectServiceClient.GetProjectDetail(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}
	// 组织回复消息
	resp := &model_project.GetProjectDetailResp{}
	copier.Copy(&resp, grpcResp)
	ctx.JSON(http.StatusOK, result.Success(resp))
}

func collectProject(ctx *gin.Context) {
	result := &common.Result{}
	// 获取参数
	memberId := ctx.GetInt64("memberId")
	var req model_project.CollectProjectReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}

	// 调用grpc接口
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()
	grpcReq := &project.CollectProjectReq{
		MemberId:    memberId,
		ProjectCode: req.ProjectCode,
		Type:        req.Type,
	}
	_, err = projectServiceClient.CollectProject(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}

	ctx.JSON(http.StatusOK, result.Success([]int{}))
}

func UpdateProjectDeletedState(ctx *gin.Context, deletedState bool) {
	result := &common.Result{}
	// 1. 解析请求消息
	var req model_project.UpdateProjectDeletedStateReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}

	// 2. 创建2秒超时的上下文
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()

	// 3. 调用gRPC服务
	_, err = projectServiceClient.UpdateProjectDeletedState(grpcCtx, &project.UpdateProjectDeletedStateReq{
		ProjectCode:  req.ProjectCode,
		DeletedState: deletedState,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}

	// 4. 返回结果
	ctx.JSON(http.StatusOK, result.Success([]int{}))
}

// 移入回收站
func recycleProject(c *gin.Context) {
	UpdateProjectDeletedState(c, true)
}

// 从回收站还原
func recoveryProject(c *gin.Context) {
	UpdateProjectDeletedState(c, false)
}

// 更新项目
func updateProject(ctx *gin.Context) {
	result := &common.Result{}
	// 1. 解析请求消息
	var req model_project.UpdateProjectReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}

	// 2. 调用gRPC服务
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()

	grpcReq := &project.UpdateProjectReq{}
	copier.Copy(grpcReq, req)
	_, err = projectServiceClient.UpdateProject(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}

	// 4. 返回结果
	ctx.JSON(http.StatusOK, result.Success([]int{}))
}
