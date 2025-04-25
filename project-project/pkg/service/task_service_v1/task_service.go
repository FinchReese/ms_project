package task_service_v1

import (
	"context"
	"strconv"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"test.com/project-common/encrypt"
	"test.com/project-common/errs"
	"test.com/project-grpc/task"
	"test.com/project-grpc/user/login"
	"test.com/project-project/internal/repo"
	"test.com/project-project/internal/rpc"
	"test.com/project-project/pkg/model"
)

type TaskService struct {
	task.UnimplementedTaskServiceServer
	taskStage  repo.TaskStageRepo
	task       repo.TaskRepo
	taskMember repo.TaskMemberRepo
}

func NewTaskService(ts repo.TaskStageRepo, t repo.TaskRepo, tm repo.TaskMemberRepo) *TaskService {
	return &TaskService{
		taskStage:  ts,
		task:       t,
		taskMember: tm,
	}
}

func (ts *TaskService) GetTaskStages(ctx context.Context, req *task.GetTaskStagesReq) (*task.GetTaskStagesResp, error) {
	// 解析请求消息获得参数
	projectCodeStr, _ := encrypt.Decrypt(req.ProjectCode, model.AESKey)
	projectId, _ := strconv.ParseInt(projectCodeStr, 10, 64)
	taskStageList, total, err := ts.taskStage.FindStagesByProjectId(ctx, projectId, int(req.Page), int(req.PageSize))
	if err != nil {
		zap.L().Error("encrypt project id err", zap.Error(err))
		return nil, errs.GrpcError(model.GetTaskStagesError)
	}

	var taskStages []*task.TaskStage
	err = copier.Copy(&taskStages, taskStageList)
	if err != nil {
		zap.L().Error("copy task stage list err", zap.Error(err))
		return nil, errs.GrpcError(model.CopyTaskStageListError)
	}
	if taskStages == nil {
		return &task.GetTaskStagesResp{List: taskStages, Total: 0}, nil
	}
	for _, taskStage := range taskStages {
		taskStage.Code, err = encrypt.EncryptInt64(int64(taskStage.Id), model.AESKey)
		if err != nil {
			zap.L().Error("encrypt task stage id err", zap.Error(err))
			return nil, errs.GrpcError(model.EncryptTaskStageIdError)
		}
		taskStage.ProjectCode = req.ProjectCode
	}

	return &task.GetTaskStagesResp{List: taskStages, Total: total}, nil
}

func (ts *TaskService) GetTasksByStageCode(ctx context.Context, in *task.GetTasksByStageCodeReq) (*task.GetTasksByStageCodeResp, error) {
	// 解析请求消息获得参数
	stageCodeStr, _ := encrypt.Decrypt(in.StageCode, model.AESKey)
	stageId, _ := strconv.ParseInt(stageCodeStr, 10, 64)
	taskList, err := ts.task.FindTasksByStageCode(ctx, int(stageId))
	if err != nil {
		zap.L().Error("find tasks by stage code err", zap.Error(err))
		return nil, errs.GrpcError(model.GetTasksByStageCodeError)
	}

	if taskList == nil || len(taskList) == 0 {
		return &task.GetTasksByStageCodeResp{List: []*task.Task{}}, nil
	}

	var dispTaskList []*task.Task
	memIdList := []int64{}
	for _, task := range taskList {
		dispTask := task.ToDisplayTask()
		// 私人的任务需要判断是否是当前用户所有
		if dispTask.Private == 1 {
			taskMemberList, err := ts.taskMember.FindTaskMembers(ctx, dispTask.Id, in.MemberId)
			if err != nil {
				zap.L().Error("find task members err", zap.Error(err))
				return nil, errs.GrpcError(model.GetTaskMembersError)
			}
			if taskMemberList != nil && len(taskMemberList) > 0 {
				dispTask.CanRead = model.CanRead
			} else {
				dispTask.CanRead = model.CanNotRead
			}
		}
		assignToStr, _ := encrypt.Decrypt(dispTask.AssignTo, model.AESKey)
		memId, _ := strconv.ParseInt(assignToStr, 10, 64)
		memIdList = append(memIdList, memId)
		dispTaskList = append(dispTaskList, dispTask)
	}

	// 向login服务一次性查询所有成员信息，提升性能
	members, err := rpc.LoginServiceClient.GetMembersByIds(ctx, &login.GetMembersByIdsReq{
		MemberIds: memIdList,
	})
	if err != nil {
		zap.L().Error("get members info error", zap.Error(err))
		return nil, errs.GrpcError(model.GetMembersInfoError)
	}
	memIdToInfo := make(map[int64]*login.MemberMessage)
	for _, member := range members.List {
		memIdToInfo[member.Id] = member
	}
	for _, dispTask := range dispTaskList {
		assignToStr, _ := encrypt.Decrypt(dispTask.AssignTo, model.AESKey)
		memId, _ := strconv.ParseInt(assignToStr, 10, 64)
		member := memIdToInfo[memId]
		e := &task.ExecutorMessage{
			Name:   member.Name,
			Avatar: member.Avatar,
		}
		dispTask.Executor = e
	}
	return &task.GetTasksByStageCodeResp{List: dispTaskList}, nil
}
