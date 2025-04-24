package task_service_v1

import (
	"context"
	"strconv"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"test.com/project-common/encrypt"
	"test.com/project-common/errs"
	"test.com/project-grpc/task"
	"test.com/project-project/internal/repo"
	"test.com/project-project/pkg/model"
)

type TaskService struct {
	task.UnimplementedTaskServiceServer
	taskStage repo.TaskStageRepo
}

func NewTaskService(ts repo.TaskStageRepo) *TaskService {
	return &TaskService{
		taskStage: ts,
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
