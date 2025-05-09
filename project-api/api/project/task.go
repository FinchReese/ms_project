package project

import (
	"context"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"test.com/project-api/pkg/model/project"
	model_project "test.com/project-api/pkg/model/project"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-common/fs"
	"test.com/project-common/time_format"
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

func getTaskMemberList(ctx *gin.Context) {
	result := &common.Result{}
	// 1. 解析请求消息
	var req model_project.GetTaskMemberListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}

	// 2. 调用gRPC服务
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()

	grpcReq := &task.GetTaskMemberListReq{
		TaskCode: req.TaskCode,
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
	}
	grpcResp, err := TaskServiceClient.GetTaskMemberList(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}

	resp := &model_project.GetTaskMemberListResp{}
	copier.Copy(&resp.List, grpcResp.List)
	if resp.List == nil {
		resp.List = []*model_project.TaskMember{}
	}
	resp.Total = grpcResp.Total
	resp.Page = req.Page

	// 4. 返回结果
	ctx.JSON(http.StatusOK, result.Success(resp))
}

func getTaskLogList(ctx *gin.Context) {
	result := &common.Result{}
	// 1. 解析请求消息
	var req model_project.GetTaskLogListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}
	// 2. 调用gRPC服务
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()

	grpcReq := &task.GetTaskLogListReq{
		TaskCode: req.TaskCode,
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
		All:      int32(req.All),
		Comment:  int32(req.Comment),
	}
	grpcResp, err := TaskServiceClient.GetTaskLogList(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}

	// 3. 组织回复消息
	resp := &model_project.GetTaskLogListResp{}
	copier.Copy(&resp.List, grpcResp.List)
	resp.Total = grpcResp.Total
	resp.Page = req.Page
	ctx.JSON(http.StatusOK, result.Success(resp))
}

func getTaskWorkTimeList(ctx *gin.Context) {
	result := &common.Result{}
	// 1. 解析请求消息
	var req model_project.GetTaskWorkTimeListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}

	// 2. 调用gRPC服务
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()

	grpcReq := &task.GetTaskWorkTimeListReq{
		TaskCode: req.TaskCode,
	}
	grpcResp, err := TaskServiceClient.GetTaskWorkTimeList(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}

	// 3. 组织回复消息
	var taskWorkTimeList []*model_project.TaskWorkTime
	copier.Copy(&taskWorkTimeList, grpcResp.List)
	if taskWorkTimeList == nil {
		taskWorkTimeList = []*model_project.TaskWorkTime{}
	}

	// 4. 返回结果
	ctx.JSON(http.StatusOK, result.Success(taskWorkTimeList))
}

func saveTaskWorkTime(ctx *gin.Context) {
	result := &common.Result{}
	// 1. 解析请求消息
	var req model_project.SaveTaskWorkTimeReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}
	beginTime, err := time_format.ParseTimeStr(req.BeginTime)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}
	memberId := ctx.GetInt64("memberId")

	// 2. 调用gRPC服务
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()
	grpcReq := &task.SaveTaskWorkTimeReq{
		TaskCode:  req.TaskCode,
		Content:   req.Content,
		BeginTime: beginTime,
		Num:       int32(req.Num),
		MemberId:  memberId,
	}
	_, err = TaskServiceClient.SaveTaskWorkTime(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}

	ctx.JSON(http.StatusOK, result.Success([]int{}))
}

func uploadFile(ctx *gin.Context) {
	result := &common.Result{}
	// 1. 解析请求消息
	var req model_project.UploadFileReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}
	// 获取上传文件
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "获取文件失败"))
		return
	}
	// 打开上传文件
	fileContent, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "打开文件内容失败"))
		return
	}
	saveDir := "upload/" + req.ProjectCode + "/" + req.TaskCode + "/" + time_format.ConvertTimeToDate(time.Now())
	// 如果文件夹不存在，则创建文件夹
	if !fs.FileExists(saveDir) {
		os.MkdirAll(saveDir, os.ModePerm)
	}
	savePath := saveDir + "/" + file.Filename
	// 将上传文件的内容追加到保存文件
	saveFile, err := os.OpenFile(savePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		zap.L().Error("打开文件错误", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "打开文件错误"))
		return
	}
	defer saveFile.Close()
	buf := make([]byte, req.CurrentChunkSize)
	_, err = io.ReadFull(fileContent, buf)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "读取上传文件内容失败"))
		return
	}
	_, err = saveFile.Write(buf)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "保存文件内容失败"))
		return
	}
	// 最后一个文件分片处理完成需要记录文件信息
	if req.ChunkNumber == req.TotalChunks {
		// 调用GRPC服务存储文件信息
		grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
		defer cancel()
		grpcReq := &task.SaveUploadFileInfoReq{
			TaskCode:    req.TaskCode,
			ProjectCode: req.ProjectCode,
			PathName:    savePath,
			FileName:    req.Filename,
			Size:        int64(req.TotalSize),
			Extension:   path.Ext(savePath),
			FileUrl:     "http://localhost/" + savePath,
			FileType:    file.Header.Get("Content-Type"),
			MemberId:    ctx.GetInt64("memberId"),
		}
		_, err = TaskServiceClient.SaveUploadFileInfo(grpcCtx, grpcReq)
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
			return
		}
	}

	// 组织回复消息
	resp := &model_project.UploadFileResp{
		File:        savePath,
		Hash:        "",
		Key:         savePath,
		Url:         "http://localhost/" + savePath,
		ProjectName: req.ProjectName,
	}
	ctx.JSON(http.StatusOK, result.Success(resp))
}

func getTaskLinkFiles(ctx *gin.Context) {
	result := &common.Result{}
	// 1. 解析请求消息
	var req model_project.GetTaskLinkFilesReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}
	// 2. 调用gRPC服务获取任务关联文件
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()
	grpcReq := &task.GetTaskLinkFilesReq{
		TaskCode: req.TaskCode,
	}
	grpcResp, err := TaskServiceClient.GetTaskLinkFiles(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}
	// 3. 组织回复消息
	var taskLinkFiles []*model_project.SourceLink
	copier.Copy(&taskLinkFiles, grpcResp.List)
	ctx.JSON(http.StatusOK, result.Success(taskLinkFiles))
}

func createComment(ctx *gin.Context) {
	result := &common.Result{}
	// 解析请求消息
	var req model_project.CreateCommentReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}
	// 调用grpc接口实现功能
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()
	grpcReq := &task.CreateTaskCommentReq{
		TaskCode:       req.TaskCode,
		MemberId:       ctx.GetInt64("memberId"),
		CommentContent: req.Comment,
	}
	_, err = TaskServiceClient.CreateTaskComment(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}
	// 回复
	ctx.JSON(http.StatusOK, result.Success(""))
}

func getUserProjectLogList(ctx *gin.Context) {
	result := &common.Result{}
	// 1. 解析请求消息
	var req model_project.GetUserProjectLogListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数错误"))
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	// 2. 调用gRPC服务获取用户项目日志列表
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()
	grpcReq := &task.GetUserProjectLogListReq{
		MemberId: ctx.GetInt64("memberId"),
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
	}
	grpcResp, err := TaskServiceClient.GetUserProjectLogList(grpcCtx, grpcReq)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}
	// 3. 组织回复消息
	// 如果列表为空，则返回空列表
	if len(grpcResp.List) == 0 || grpcResp.Total == 0 {
		ctx.JSON(http.StatusOK, result.Success([]*model_project.ProjectLog{}))
		return
	}
	var projectLogList []*model_project.ProjectLog
	copier.Copy(&projectLogList, grpcResp.List)
	ctx.JSON(http.StatusOK, result.Success(projectLogList))
}
