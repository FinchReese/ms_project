package project_auth_service_v1

import (
	"context"

	"github.com/jinzhu/copier"
	"test.com/project-common/errs"
	project_auth "test.com/project-grpc/project_auth"
	"test.com/project-project/internal/domain"
)

type ProjectAuthService struct {
	project_auth.UnimplementedProjectAuthServiceServer
	projectAuth *domain.ProjectAuthDomain
	userDomain  *domain.UserDomain
}

func NewProjectAuthService(projectAuth *domain.ProjectAuthDomain, userDomain *domain.UserDomain) *ProjectAuthService {
	return &ProjectAuthService{projectAuth: projectAuth, userDomain: userDomain}
}

func (s *ProjectAuthService) GetProjectAuthList(ctx context.Context, req *project_auth.GetProjectAuthListReq) (*project_auth.GetProjectAuthListResp, error) {
	// 根据memberId获取organizationCode
	organizationCode, err := s.userDomain.GetOrganizationCodeByMemberId(ctx, req.MemberId)
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	// 获取项目权限列表
	projectAuthList, total, err := s.projectAuth.GetProjectAuthListByOrganizationCode(ctx, organizationCode,
		int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	// 组织回复消息
	var projectAuthListResp []*project_auth.ProjectAuth
	copier.Copy(&projectAuthListResp, projectAuthList)
	return &project_auth.GetProjectAuthListResp{
		List:  projectAuthListResp,
		Total: total,
	}, nil
}
