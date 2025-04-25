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
	"test.com/project-grpc/task"
)

func getTaskStage(ctx *gin.Context) {
	result := &common.Result{}
	// 1. 解析请求消息
	var req model_project.GetTaskStageReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}

	// 2. 调用gRPC服务
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()

	grpcReq := &task.GetTaskStagesReq{}
	copier.Copy(grpcReq, req)
	grpcResp, err := TaskServiceClient.GetTaskStages(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}
	// 3. 组织回复消息
	resp := &model_project.GetTaskStageResp{}
	resp.Total = grpcResp.Total
	resp.Page = req.Page
	copier.Copy(&resp.List, grpcResp.List)
	for _, taskStage := range resp.List {
		taskStage.TasksLoading = true  //任务加载状态
		taskStage.FixedCreator = false //添加任务按钮定位
		taskStage.ShowTaskCard = false //是否显示创建卡片
		taskStage.Tasks = []int{}
		taskStage.DoneTasks = []int{}
		taskStage.UnDoneTasks = []int{}
	}

	// 4. 返回结果
	ctx.JSON(http.StatusOK, result.Success(resp))

}

func getTaskList(ctx *gin.Context) {
	result := &common.Result{}
	// 1. 解析请求消息
	var req model_project.GetTaskListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}
	memberId := ctx.GetInt64("memberId")

	// 2. 调用gRPC服务
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()

	grpcReq := &task.GetTasksByStageCodeReq{
		MemberId:  memberId,
		StageCode: req.StageCode,
	}
	grpcResp, err := TaskServiceClient.GetTasksByStageCode(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}
	// 3. 组织回复消息
	var dispTaskList []*model_project.DispTask
	copier.Copy(&dispTaskList, grpcResp)
	if dispTaskList == nil {
		dispTaskList = []*model_project.DispTask{}
	}
	//返回给前端的数据 一定不要是null
	for _, v := range dispTaskList {
		if v.Tags == nil {
			v.Tags = []int{}
		}
		if v.ChildCount == nil {
			v.ChildCount = []int{}
		}
	}
	ctx.JSON(http.StatusOK, result.Success(dispTaskList))
}
