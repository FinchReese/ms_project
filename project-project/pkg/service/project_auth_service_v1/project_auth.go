package project_auth_service_v1

import (
	"context"

	"github.com/jinzhu/copier"
	"test.com/project-common/errs"
	project_auth "test.com/project-grpc/project_auth"
	"test.com/project-project/internal/database/gorm"
	"test.com/project-project/internal/database/trans"
	"test.com/project-project/internal/domain"
	"test.com/project-project/pkg/model"
)

type ProjectAuthService struct {
	project_auth.UnimplementedProjectAuthServiceServer
	projectAuth           *domain.ProjectAuthDomain
	userDomain            *domain.UserDomain
	projectNodeDomain     *domain.ProjectNodeDomain
	projectAuthNodeDomain *domain.ProjectAuthNodeDomain
	tran                  *trans.TransactionImpl
}

func NewProjectAuthService(projectAuth *domain.ProjectAuthDomain, userDomain *domain.UserDomain, projectNodeDomain *domain.ProjectNodeDomain, projectAuthNodeDomain *domain.ProjectAuthNodeDomain, tran *trans.TransactionImpl) *ProjectAuthService {
	return &ProjectAuthService{
		projectAuth:           projectAuth,
		userDomain:            userDomain,
		projectNodeDomain:     projectNodeDomain,
		projectAuthNodeDomain: projectAuthNodeDomain,
		tran:                  tran,
	}
}

func (pa *ProjectAuthService) GetProjectAuthList(ctx context.Context, req *project_auth.GetProjectAuthListReq) (*project_auth.GetProjectAuthListResp, error) {
	// 根据memberId获取organizationCode
	organizationCode, err := pa.userDomain.GetOrganizationCodeByMemberId(ctx, req.MemberId)
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	// 获取项目权限列表
	projectAuthList, total, err := pa.projectAuth.GetProjectAuthListByOrganizationCode(ctx, organizationCode,
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

func (pa *ProjectAuthService) ProjectAuthNodeApply(ctx context.Context, req *project_auth.ProjectAuthNodeApplyReq) (*project_auth.ProjectAuthNodeApplyResp, error) {
	switch req.Action {
	case model.ProjectAuthApplyActionGetNode:
		return pa.getNode(ctx, req.AuthId)
	case model.ProjectAuthApplyActionSave:
		return pa.saveNode(ctx, req.AuthId, req.NodeList)
	case model.ProjectAuthApplyActionFilter:
		return &project_auth.ProjectAuthNodeApplyResp{
			List: []*project_auth.ProjectNodeMessage{},
		}, nil
	default:
		{
			return nil, errs.GrpcError(model.InvalidActionType)
		}
	}
}

func (pa *ProjectAuthService) getNode(ctx context.Context, authId int64) (*project_auth.ProjectAuthNodeApplyResp, error) {
	nodeTree, checkUrlList, err := pa.projectNodeDomain.GetProjectNodeListByAuthId(ctx, authId)
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	var projectNodeMessageList []*project_auth.ProjectNodeMessage
	copier.Copy(&projectNodeMessageList, nodeTree)
	return &project_auth.ProjectAuthNodeApplyResp{
		List:        projectNodeMessageList,
		CheckedList: checkUrlList,
	}, nil
}

func (pa *ProjectAuthService) saveNode(ctx context.Context, authId int64, nodeList []string) (*project_auth.ProjectAuthNodeApplyResp, error) {
	// 在事务中完成删除节点和添加节点操作
	err := pa.tran.ExecTran(func(dbConn trans.DbConn) error {
		conn := dbConn.(*gorm.MysqlConn)
		bErr := pa.projectAuthNodeDomain.UpdateProjectAuthNode(ctx, authId, nodeList, conn.TranDb)
		if bErr != nil {
			return errs.GrpcError(bErr)
		}
		return nil
	})
	return nil, err
}
