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
	proNode         repo.ProjectNodeRepo
	projectAuthNode *ProjectAuthNodeDomain
}

func NewProjectNodeDomain(proNode repo.ProjectNodeRepo, projectAuthNodeDomain *ProjectAuthNodeDomain) *ProjectNodeDomain {
	return &ProjectNodeDomain{
		proNode:         proNode,
		projectAuthNode: projectAuthNodeDomain,
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

// 根据auth_id获取节点信息和有权限节点URL列表
func (pn *ProjectNodeDomain) GetProjectNodeListByAuthId(ctx context.Context, authId int64) ([]*data.ProjectNodeAuthTree, []string, *errs.BError) {
	// 获取节点列表
	nodeList, err := pn.proNode.GetAllProjectNodeList(ctx)
	if err != nil {
		zap.L().Error("GetAllProjectNodeList error", zap.Error(err))
		return nil, nil, model.GetProjectNodeListError
	}
	// 获取有权限节点URL列表
	checkUrlList, bErr := pn.projectAuthNode.GetProjectAuthNodeList(ctx, authId)
	if bErr != nil {
		zap.L().Error("GetProjectAuthNodeList error", zap.Error(errs.GrpcError(bErr)))
		return nil, nil, bErr
	}
	// 将节点列表转换为树结构
	nodeTree := data.ToAuthNodeTreeList(nodeList, checkUrlList)
	return nodeTree, checkUrlList, nil
}
