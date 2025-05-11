package repo

import (
	"context"

	"test.com/project-project/internal/data"
)

type DepartmentRepo interface {
	GetDepartmentInfo(ctx context.Context, departmentCode int64) (*data.Department, error)
}
