package domain

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	"test.com/project-common/encrypt"
	"test.com/project-common/errs"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/repo"
	"test.com/project-project/pkg/model"
)

type ProjectAuthDomain struct {
	projectAuthRepo     repo.ProjectAuthRepo
	memberAccountRepo   repo.MemberAccountRepo
	projectAuthNode     *ProjectAuthNodeDomain
	taskDomain          *TaskDomain
	projectMemberDomain *ProjectMemberDomain
}

func NewProjectAuthDomain(projectAuthRepo repo.ProjectAuthRepo, memberAccountRepo repo.MemberAccountRepo, projectAuthNodeDomain *ProjectAuthNodeDomain,
	taskDomain *TaskDomain, projectMemberDomain *ProjectMemberDomain) *ProjectAuthDomain {
	return &ProjectAuthDomain{
		projectAuthRepo:     projectAuthRepo,
		memberAccountRepo:   memberAccountRepo,
		projectAuthNode:     projectAuthNodeDomain,
		taskDomain:          taskDomain,
		projectMemberDomain: projectMemberDomain,
	}
}

func (pa *ProjectAuthDomain) GetProjectAuthList(ctx context.Context, organizationCode int64) ([]*data.ProjectAuthDisplay, *errs.BError) {
	projectAuthList, err := pa.projectAuthRepo.GetProjectAuthList(ctx, organizationCode)
	if err != nil {
		zap.L().Error("get project auth list error", zap.Error(err))
		return nil, model.GetProjectAuthListError
	}
	var proAuthDispList []*data.ProjectAuthDisplay
	for _, projectAuth := range projectAuthList {
		proAuthDispList = append(proAuthDispList, projectAuth.ToDisplay())
	}
	return proAuthDispList, nil
}

func (pa *ProjectAuthDomain) GetProjectAuthListByOrganizationCode(ctx context.Context, organizationCode int64, page int, pageSize int) ([]*data.ProjectAuthDisplay, int64, *errs.BError) {
	projectAuthList, total, err := pa.projectAuthRepo.GetProjectAuthListByOrganizationCode(ctx, organizationCode, page, pageSize)
	if err != nil {
		zap.L().Error("get project auth list error", zap.Error(err))
		return nil, 0, model.GetProjectAuthListError
	}
	var proAuthDispList []*data.ProjectAuthDisplay
	for _, projectAuth := range projectAuthList {
		proAuthDispList = append(proAuthDispList, projectAuth.ToDisplay())
	}
	return proAuthDispList, total, nil
}

// 获取指定member code的有权限节点URL列表
func (pa *ProjectAuthDomain) GetAuthNodeUrlList(ctx context.Context, memberCode int64) ([]string, *errs.BError) {
	// 1. 根据member code获取成员账号
	memberAccount, err := pa.memberAccountRepo.GetMemberAccountByMemberCode(ctx, memberCode)
	if err != nil {
		zap.L().Error("get member account by member code error", zap.Error(err))
		return nil, model.GetMemberAccountError
	}
	// 2. 解析权限id
	authId, err := strconv.ParseInt(memberAccount.Authorize, 10, 64)
	if err != nil {
		zap.L().Error("parse authorize error", zap.Error(err))
		return nil, model.ParseAuthIdError
	}
	// 3. 根据项目权限id获取项目权限节点列表
	projectAuthNodeList, bErr := pa.projectAuthNode.GetProjectAuthNodeList(ctx, authId)
	if bErr != nil {
		zap.L().Error("get project auth node list error", zap.Error(errs.GrpcError(bErr)))
		return nil, bErr
	}
	return projectAuthNodeList, nil
}

// 根据member id 、 project code 和 task code 判断是否有项目权限
func (pa *ProjectAuthDomain) CheckProjectAuth(ctx context.Context, memberId int64, projectCode string, taskCode string) (isMember bool,
	isOwner bool, isPrivateProject bool, bErr *errs.BError) {
	// 如果projectCode不为空，解析为project id
	var projectId int64
	var projectCodeValid bool = false // 记录project code是否有效
	if projectCode != "" {
		var err error
		projectId, err = encrypt.DecryptToInt64(projectCode, model.AESKey)
		if err != nil {
			zap.L().Error("parse project code error", zap.Error(err))
			return false, false, false, model.ParseProjectIdError
		}
		projectCodeValid = true
	} else { // 尝试根据task code获取project code
		// 解析task code
		taskId, err := encrypt.DecryptToInt64(taskCode, model.AESKey)
		if err != nil {
			zap.L().Error("parse task code error", zap.Error(err))
			return false, false, false, model.ParseTaskIdError
		}
		// 根据task id获取project code
		projectCode, isExist, bErr := pa.taskDomain.GetProjectCodeByTaskId(ctx, taskId)
		if bErr != nil {
			zap.L().Error("get project code by task id error", zap.Error(errs.GrpcError(bErr)))
			return false, false, false, bErr
		}
		if !isExist {
			return false, false, false, nil
		}
		projectId = projectCode
		projectCodeValid = true
	}
	if !projectCodeValid {
		return false, false, false, model.InvalidProjectCodeOrTaskCode
	}
	// 4. 根据member id和project id查询项目信息
	projectAndMember, err := pa.projectMemberDomain.GetProjectAndMember(ctx, memberId, projectId)
	if err != nil {
		zap.L().Error("get project and member error", zap.Error(err))
		return false, false, false, model.GetProjectAndMemberError
	}
	// 5. 判断是否有项目权限
	if projectAndMember == nil { // 查不到记录就是没权限
		return false, false, false, nil
	}
	isMember = true // 有记录就是项目成员
	isOwner = projectAndMember.IsOwner == memberId
	isPrivateProject = projectAndMember.Project.Private == 1
	return
}
