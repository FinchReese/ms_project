package repo

import (
	"context"

	"test.com/project-project/internal/data"
)

type ProjectAuthRepo interface {
	GetProjectAuthList(ctx context.Context, organizationCode int64) ([]*data.ProjectAuth, error)
	// 根据organization code、页号和页大小获取项目权限列表
	GetProjectAuthListByOrganizationCode(ctx context.Context, organizationCode int64, page int, pageSize int) ([]*data.ProjectAuth, int64, error)
}
