package account_service_v1

import (
	"context"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"test.com/project-common/encrypt"
	"test.com/project-common/errs"
	"test.com/project-grpc/account"
	"test.com/project-project/internal/domain"
	"test.com/project-project/pkg/model"
)

type AccountService struct {
	account.UnimplementedAccountServiceServer
	memberAccount *domain.MemberAccountDomain
	projectAuth   *domain.ProjectAuthDomain
	user          *domain.UserDomain
}

func NewAccountService(mad *domain.MemberAccountDomain, pad *domain.ProjectAuthDomain, ud *domain.UserDomain) *AccountService {
	return &AccountService{
		memberAccount: mad,
		projectAuth:   pad,
		user:          ud,
	}
}

func (s *AccountService) GetAccountList(ctx context.Context, req *account.GetAccountListReq) (*account.GetAccountListResp, error) {
	// 获取organization code
	organizationCode, err := s.user.GetOrganizationCodeByMemberId(ctx, req.MemberId)
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	// 解析Department Code
	departmentId, _ := encrypt.DecryptToInt64(req.DepartmentCode, model.AESKey)

	// 获取account list
	memberAccountDispList, total, err := s.memberAccount.GetMemberAccountList(ctx, int(req.SearchType), organizationCode, departmentId, int(req.Page), int(req.PageSize))
	if err != nil {
		zap.L().Error("GetMemberAccountList error", zap.Error(errs.GrpcError(err)))
		return nil, errs.GrpcError(err)
	}

	// 获取 auth list
	projectAuthDispList, err := s.projectAuth.GetProjectAuthList(ctx, organizationCode)
	if err != nil {
		zap.L().Error("GetProjectAuthList error", zap.Error(errs.GrpcError(err)))
		return nil, errs.GrpcError(err)
	}

	// 组织回复消息
	var memberAccountList []*account.MemberAccount
	copier.Copy(&memberAccountList, memberAccountDispList)

	var projectAuthList []*account.ProjectAuth
	copier.Copy(&projectAuthList, projectAuthDispList)

	return &account.GetAccountListResp{
		Total:       total,
		AccountList: memberAccountList,
		AuthList:    projectAuthList,
	}, nil
}
