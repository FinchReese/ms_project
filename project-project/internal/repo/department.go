package repo

import (
	"context"

	"test.com/project-project/internal/data"
)

type DepartmentRepo interface {
	GetDepartmentInfo(ctx context.Context, departmentCode int64) (*data.Department, error)
	// 根据organization code、 pcode、页号和页面大小 获取部门列表
	GetDepartmentList(ctx context.Context, organizationCode int64, pcode int64, page int, pageSize int) ([]*data.Department, int64, error)
	// 添加新记录
	AddDepartment(ctx context.Context, department *data.Department) error
	// 根据organization code、 pcode和name 获取部门列表
	SearchDepartmentList(ctx context.Context, organizationCode int64, pcode int64, name string) ([]*data.Department, error)
	// 根据id获取部门
	GetDepartmentById(ctx context.Context, id int64) (*data.Department, error)
}
