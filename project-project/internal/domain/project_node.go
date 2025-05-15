package domain

import (
	"context"
	"test.com/project-project/internal/data"

	"go.uber.org/zap"
	"test.com/project-common/errs"
	"test.com/project-project/internal/repo"
	"test.com/project-project/pkg/model"
)

type ProjectNodeDomain struct {
	proNode repo.ProjectNodeRepo
}

func NewProjectNodeDomain(proNode repo.ProjectNodeRepo) *ProjectNodeDomain {
	return &ProjectNodeDomain{
		proNode: proNode,
	}
}

func (pn *ProjectNodeDomain) GetAllProjectNodeList(ctx context.Context) ([]*data.ProjectNodeTree, *errs.BError) {
	nodeList, err := pn.proNode.GetAllProjectNodeList(ctx)
	if err != nil {
		zap.L().Error("GetAllProjectNodeList error", zap.Error(err))
		return nil, model.GetProjectNodeListError
	}
	return data.ToNodeTreeList(nodeList), nil
}
