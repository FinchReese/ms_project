package domain

import (
	"context"

	"go.uber.org/zap"
	"test.com/project-common/errs"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/repo"
	"test.com/project-project/pkg/model"
)

type DepartmentDomain struct {
	departmentRepo repo.DepartmentRepo
}

func NewDepartmentDomain(departmentRepo repo.DepartmentRepo) *DepartmentDomain {
	return &DepartmentDomain{departmentRepo: departmentRepo}
}

func (d *DepartmentDomain) GetDepartmentInfo(ctx context.Context, departmentCode int64) (*data.Department, *errs.BError) {
	department, err := d.departmentRepo.GetDepartmentInfo(ctx, departmentCode)
	if err != nil {
		zap.L().Error("get department info error", zap.Error(err))
		return nil, model.GetDepartmentByIdError
	}
	return department, nil
}
