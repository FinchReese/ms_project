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
