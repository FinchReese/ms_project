package domain

import (
	"context"

	"go.uber.org/zap"
	"test.com/project-common/errs"
	"test.com/project-grpc/user/login"
	"test.com/project-project/internal/rpc"
	"test.com/project-project/pkg/model"
)

type UserDomain struct {
	loginServiceClient login.LoginServiceClient
}

func NewUserDomain() *UserDomain {
	return &UserDomain{
		loginServiceClient: rpc.LoginServiceClient,
	}
}

func (u *UserDomain) GetIdToMemberMap(ctx context.Context, userIdList []int64) (map[int64]*login.MemberMessage, *errs.BError) {
	members, err := u.loginServiceClient.GetMembersByIds(ctx, &login.GetMembersByIdsReq{
		MemberIds: userIdList,
	})
	if err != nil {
		zap.L().Error("get members info error", zap.Error(err))
		return nil, model.GetMembersInfoError
	}
	idToMemberMap := make(map[int64]*login.MemberMessage)
	for _, member := range members.List {
		idToMemberMap[member.Id] = member
	}
	return idToMemberMap, nil
}

func (u *UserDomain) GetMemberInfo(ctx context.Context, memberCode int64) (*login.MemberMessage, *errs.BError) {
	member, err := u.loginServiceClient.GetMemberById(ctx, &login.GetMemberByIdReq{
		MemberId: memberCode,
	})
	if err != nil {
		zap.L().Error("get member account error", zap.Error(err))
		return nil, model.GetMemberByIdError
	}
	return member, nil
}

// GetOrganizationCodeByMemberId 根据member id查询organization code
func (u *UserDomain) GetOrganizationCodeByMemberId(ctx context.Context, memberId int64) (int64, *errs.BError) {
	// 调用login服务的RPC接口查询organization信息
	resp, err := u.loginServiceClient.GetOrganizationList(ctx, &login.GetOrganizationListReq{
		MemberId: memberId,
	})
	if err != nil {
		zap.L().Error("get organization by member id error", zap.Error(err))
		return 0, model.GetOrganizationListError
	}
	if len(resp.OrgList) == 0 {
		return 0, nil
	}
	return resp.OrgList[0].Id, nil
}
