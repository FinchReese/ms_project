package domain

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	"test.com/project-common/errs"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/repo"
	"test.com/project-project/pkg/model"
)

type ProjectAuthDomain struct {
	projectAuthRepo   repo.ProjectAuthRepo
	memberAccountRepo repo.MemberAccountRepo
	projectAuthNode   *ProjectAuthNodeDomain
}

func NewProjectAuthDomain(projectAuthRepo repo.ProjectAuthRepo, memberAccountRepo repo.MemberAccountRepo, projectAuthNodeDomain *ProjectAuthNodeDomain) *ProjectAuthDomain {
	return &ProjectAuthDomain{
		projectAuthRepo:   projectAuthRepo,
		memberAccountRepo: memberAccountRepo,
		projectAuthNode:   projectAuthNodeDomain,
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
