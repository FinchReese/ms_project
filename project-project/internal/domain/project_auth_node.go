package domain

import (
	"context"

	"go.uber.org/zap"
	"test.com/project-common/errs"
	"test.com/project-project/internal/repo"
	"test.com/project-project/pkg/model"
)

type ProjectAuthNodeDomain struct {
	projectAuthNode repo.ProjectAuthNodeRepo
}

func NewProjectAuthNodeDomain(projectAuthNodeRepo repo.ProjectAuthNodeRepo) *ProjectAuthNodeDomain {
	return &ProjectAuthNodeDomain{projectAuthNode: projectAuthNodeRepo}
}

func (d *ProjectAuthNodeDomain) GetProjectAuthNodeList(ctx context.Context, authId int64) ([]string, *errs.BError) {
	nodes, err := d.projectAuthNode.GetProjectAuthNodeList(ctx, authId)
	if err != nil {
		zap.L().Error("get project auth node list error", zap.Error(err))
		return nil, model.GetProjectAuthNodeListError
	}
	return nodes, nil
}
