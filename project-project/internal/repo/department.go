package repo

import (
	"context"

	"test.com/project-project/internal/data"
)

type DepartmentRepo interface {
	GetDepartmentInfo(ctx context.Context, departmentCode int64) (*data.Department, error)
	// 根据organization code、 pcode、页号和页面大小 获取部门列表
	GetDepartmentList(ctx context.Context, organizationCode int64, pcode int64, page int, pageSize int) ([]*data.Department, int64, error)
}
