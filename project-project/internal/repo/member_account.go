package repo

import (
	"context"

	"test.com/project-project/internal/data"
)

type MemberAccountRepo interface {
	// 根据查询类型、组织id、部门id、页号、每页大小获取成员账号列表
	GetMemberAccountList(ctx context.Context, queryType int, organizationCode int64, departmentCode int64, page int, pageSize int) ([]*data.MemberAccount, int64, error)
}
