package domain

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
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

func (d *ProjectAuthNodeDomain) UpdateProjectAuthNode(ctx context.Context, authId int64, nodeList []string, db *gorm.DB) *errs.BError {
	// 先删除指定auth id的所有节点
	err := d.projectAuthNode.DeleteProjectAuthNode(ctx, authId, db)
	if err != nil {
		zap.L().Error("delete project auth node error", zap.Error(err))
		return model.DeleteProjectAuthNodeError
	}
	// 再根据nodeList添加新的节点
	err = d.projectAuthNode.AddProjectAuthNode(ctx, authId, nodeList, db)
	if err != nil {
		zap.L().Error("add project auth node error", zap.Error(err))
		return model.AddProjectAuthNodeError
	}
	return nil
}
