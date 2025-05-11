package domain

import (
	"context"

	"go.uber.org/zap"
	"test.com/project-common/errs"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/repo"
	"test.com/project-project/pkg/model"
)

type MemberAccountDomain struct {
	memberAccountRepo repo.MemberAccountRepo
	userDomain        *UserDomain
	departmentDomain  *DepartmentDomain
}

func NewMemberAccountDomain(memberAccountRepo repo.MemberAccountRepo, userDomain *UserDomain, departmentDomain *DepartmentDomain) *MemberAccountDomain {
	return &MemberAccountDomain{
		memberAccountRepo: memberAccountRepo,
		userDomain:        userDomain,
		departmentDomain:  departmentDomain,
	}
}

func (m *MemberAccountDomain) GetMemberAccountList(ctx context.Context, queryType int, organizationCode int64, departmentCode int64, page int,
	pageSize int) ([]*data.MemberAccountDisplay, int64, *errs.BError) {
	memberAccountList, total, err := m.memberAccountRepo.GetMemberAccountList(ctx, queryType, organizationCode, departmentCode, page, pageSize)
	if err != nil {
		zap.L().Error("get member account list error", zap.Error(err))
		return nil, 0, model.GetMemberAccountListError
	}
	memberAccountDisplayList := make([]*data.MemberAccountDisplay, 0)
	for _, memberAccount := range memberAccountList {
		memberAccountDisplay := memberAccount.ToDisplay()
		// 查询成员信息
		memberInfo, err := m.userDomain.GetMemberInfo(ctx, memberAccount.MemberCode)
		if err != nil {
			zap.L().Error("get member info error", zap.Error(errs.GrpcError(err)))
			return nil, 0, model.GetMemberByIdError
		}
		memberInfo.Avatar = memberInfo.Avatar
		// 查询部门信息
		if memberAccount.DepartmentCode > 0 {
			departmentInfo, err := m.departmentDomain.GetDepartmentInfo(ctx, memberAccount.DepartmentCode)
			if err != nil {
				zap.L().Error("get department info error", zap.Error(errs.GrpcError(err)))
				return nil, 0, model.GetDepartmentByIdError
			}
			memberAccountDisplay.Departments = departmentInfo.Name
		}
		memberAccountDisplayList = append(memberAccountDisplayList, memberAccountDisplay)
	}
	return memberAccountDisplayList, total, nil
}
