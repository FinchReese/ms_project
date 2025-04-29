package project

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"test.com/project-api/pkg/model/project"
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
	copier.Copy(&dispTaskList, grpcResp.List)
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

// saveTask 创建任务处理函数
func saveTask(ctx *gin.Context) {
	result := &common.Result{}
	// 1. 解析请求消息
	var req model_project.SaveTaskReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}
	// 获取当前登录用户ID
	memberId := ctx.GetInt64("memberId")

	// 2. 创建一个等待2秒钟的上下文
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()

	// 3. 调用gRPC服务
	grpcReq := &task.SaveTaskReq{
		Name:        req.Name,
		StageCode:   req.StageCode,
		ProjectCode: req.ProjectCode,
		AssignTo:    req.AssignTo,
		MemberId:    memberId,
	}

	grpcResp, err := TaskServiceClient.SaveTask(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}

	// 4. 组织回复消息
	resp := &model_project.SaveTaskResp{}
	copier.Copy(resp, grpcResp)
	if resp != nil {
		if resp.Tags == nil {
			resp.Tags = []int{}
		}
		if resp.ChildCount == nil {
			resp.ChildCount = []int{}
		}
	}

	// 5. 返回结果
	ctx.JSON(http.StatusOK, result.Success(resp))
}

func moveTask(ctx *gin.Context) {
	result := &common.Result{}
	// 1. 解析请求消息
	var req model_project.MoveTaskReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}
	// 2. 调用gRPC服务
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()

	grpcReq := &task.MoveTaskReq{
		OriginTaskCode:  req.PreTaskCode,
		TargetTaskCode:  req.NextTaskCode,
		TargetStageCode: req.ToStageCode,
	}
	_, err = TaskServiceClient.MoveTask(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}

	// 3. 返回结果
	ctx.JSON(http.StatusOK, result.Success([]int{}))
}

func getTaskListByType(ctx *gin.Context) {
	result := &common.Result{}
	// 1. 解析请求消息
	var req model_project.GetTaskListByTypeReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}
	memberId := ctx.GetInt64("memberId")

	// 2. 调用gRPC服务
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()

	grpcReq := &task.GetTaskListReq{
		TaskType: req.TaskType,
		MemberId: memberId,
		Done:     int32(req.Done),
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
	}
	grpcResp, err := TaskServiceClient.GetTaskList(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}

	// 3. 组织回复消息
	resp := &model_project.GetTaskListByTypeResp{}
	copier.Copy(resp, grpcResp)
	if resp.List == nil {
		resp.List = []*model_project.MyTaskDisplay{}
	}
	for _, v := range resp.List {
		v.ProjectInfo = project.ProjectInfo{
			Name: v.ProjectName,
			Code: v.ProjectCode,
		}
	}

	// 4. 返回结果
	ctx.JSON(http.StatusOK, result.Success(resp))
}

func getTaskDetail(ctx *gin.Context) {
	result := &common.Result{}
	// 1. 解析请求消息
	var req model_project.GetTaskDetailReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}
	memberId := ctx.GetInt64("memberId")

	// 2. 调用gRPC服务
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()

	grpcReq := &task.GetTaskDetailReq{
		TaskCode: req.TaskCode,
		MemberId: memberId,
	}
	grpcResp, err := TaskServiceClient.GetTaskDetail(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}
	resp := &model_project.GetTaskDetailResp{}
	copier.Copy(&resp, grpcResp)
	if resp != nil {
		if resp.Tags == nil {
			resp.Tags = []int{}
		}
		if resp.ChildCount == nil {
			resp.ChildCount = []int{}
		}
	}
	ctx.JSON(http.StatusOK, result.Success(resp))
}
