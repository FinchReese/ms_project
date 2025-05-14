package domain

import (
	"context"

	"go.uber.org/zap"
	"test.com/project-common/errs"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/repo"
	"test.com/project-project/pkg/model"
)

type ProjectAuthDomain struct {
	projectAuthRepo repo.ProjectAuthRepo
}

func NewProjectAuthDomain(projectAuthRepo repo.ProjectAuthRepo) *ProjectAuthDomain {
	return &ProjectAuthDomain{projectAuthRepo: projectAuthRepo}
}

func (d *ProjectAuthDomain) GetProjectAuthList(ctx context.Context, organizationCode int64) ([]*data.ProjectAuthDisplay, *errs.BError) {
	projectAuthList, err := d.projectAuthRepo.GetProjectAuthList(ctx, organizationCode)
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

func (d *ProjectAuthDomain) GetProjectAuthListByOrganizationCode(ctx context.Context, organizationCode int64, page int, pageSize int) ([]*data.ProjectAuthDisplay, int64, *errs.BError) {
	projectAuthList, total, err := d.projectAuthRepo.GetProjectAuthListByOrganizationCode(ctx, organizationCode, page, pageSize)
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
