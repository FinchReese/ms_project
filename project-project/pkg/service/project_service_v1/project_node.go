package project_service_v1

import (
	"context"

	"github.com/jinzhu/copier"
	"test.com/project-common/errs"
	"test.com/project-grpc/project_node"
	"test.com/project-project/internal/domain"
)

type ProjectNodeService struct {
	project_node.UnimplementedProjectNodeServiceServer
	projectNodeDomain *domain.ProjectNodeDomain
}

func NewProjectNodeService(projectNodeDomain *domain.ProjectNodeDomain) *ProjectNodeService {
	return &ProjectNodeService{
		projectNodeDomain: projectNodeDomain,
	}
}

func (s *ProjectNodeService) GetProjectNodeList(ctx context.Context, req *project_node.GetProjectNodeListReq) (*project_node.GetProjectNodeListResp, error) {
	nodeList, berr := s.projectNodeDomain.GetAllProjectNodeList(ctx)
	if berr != nil {
		return nil, errs.GrpcError(berr)
	}
	resp := &project_node.GetProjectNodeListResp{}
	copier.Copy(&resp.Nodes, nodeList)
	return resp, nil
}
