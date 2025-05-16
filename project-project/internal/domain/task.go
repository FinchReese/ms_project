package domain

import (
	"context"

	"go.uber.org/zap"
	"test.com/project-common/errs"
	"test.com/project-project/internal/repo"
	"test.com/project-project/pkg/model"
)

type TaskDomain struct {
	taskRepo repo.TaskRepo
}

func NewTaskDomain(taskRepo repo.TaskRepo) *TaskDomain {
	return &TaskDomain{taskRepo: taskRepo}
}

// 根据task id获取project code
func (t *TaskDomain) GetProjectCodeByTaskId(ctx context.Context, taskId int64) (int64, bool, *errs.BError) {
	projectCode, isExist, err := t.taskRepo.GetProjectCodeByTaskId(ctx, taskId)
	if err != nil {
		zap.L().Error("get project code by task id error", zap.Error(err))
		return 0, false, model.GetProjectCodeByTaskIdError
	}
	return projectCode, isExist, nil
}
