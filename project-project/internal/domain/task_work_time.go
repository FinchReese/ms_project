package domain

import (
	"context"

	"go.uber.org/zap"
	"test.com/project-common/encrypt"
	"test.com/project-common/errs"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/repo"
	"test.com/project-project/pkg/model"
)

type TaskWorkTimeDomain struct {
	taskWorkTime repo.TaskWorkTimeRepo
	userDomain   *UserDomain
}

func NewTaskWorkTimeDomain(taskWorkTimeRepo repo.TaskWorkTimeRepo, userDomain *UserDomain) *TaskWorkTimeDomain {
	return &TaskWorkTimeDomain{
		taskWorkTime: taskWorkTimeRepo,
		userDomain:   userDomain,
	}
}

func (t *TaskWorkTimeDomain) GetTaskWorkTimeList(ctx context.Context, taskId int64) ([]*data.TaskWorkTimeDisplay, *errs.BError) {
	taskWorkTimeList, err := t.taskWorkTime.GetTaskWorkTimeList(ctx, taskId)
	if err != nil {
		zap.L().Error("get task work time list error", zap.Error(err))
		return nil, model.GetTaskWorkTimeListError
	}
	if len(taskWorkTimeList) == 0 {
		return []*data.TaskWorkTimeDisplay{}, nil
	}
	// 收集任务工时的成员id，统一查找成员信息
	memberIdList := []int64{}
	for _, taskWorkTime := range taskWorkTimeList {
		memberIdList = append(memberIdList, taskWorkTime.MemberCode)
	}
	memberIdToInfo, bErr := t.userDomain.GetIdToMemberMap(ctx, memberIdList)
	if bErr != nil {
		zap.L().Error("get id to member map error", zap.Error(errs.GrpcError(bErr)))
		return nil, bErr
	}

	taskWorkTimeDisplayList := []*data.TaskWorkTimeDisplay{}
	for _, taskWorkTime := range taskWorkTimeList {
		taskWorkTimeDisplay := taskWorkTime.ToDisplay()
		memberInfo := memberIdToInfo[taskWorkTime.MemberCode]
		taskWorkTimeDisplay.Member.Name = memberInfo.Name
		taskWorkTimeDisplay.Member.Avatar = memberInfo.Avatar
		taskWorkTimeDisplay.Member.Id = memberInfo.Id
		taskWorkTimeDisplay.Member.Code, _ = encrypt.EncryptInt64(memberInfo.Id, model.AESKey)
		taskWorkTimeDisplayList = append(taskWorkTimeDisplayList, taskWorkTimeDisplay)
	}
	return taskWorkTimeDisplayList, nil
}
