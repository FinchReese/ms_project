package domain

import (
	"context"
	"time"

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

func (d *DepartmentDomain) GetDepartmentList(ctx context.Context, organizationCode int64, pcode int64, page int, pageSize int) ([]*data.DepartmentDisplay, int64, *errs.BError) {
	departments, total, err := d.departmentRepo.GetDepartmentList(ctx, organizationCode, pcode, page, pageSize)
	if err != nil {
		zap.L().Error("get department list error", zap.Error(err))
		return nil, 0, model.GetDepartmentListError
	}
	departmentDisplays := make([]*data.DepartmentDisplay, 0, total)
	for _, department := range departments {
		departmentDisplays = append(departmentDisplays, department.ToDisplay())
	}
	return departmentDisplays, total, nil
}

func (d *DepartmentDomain) AddDepartment(ctx context.Context, organizationCode int64, pcode int64, name string) (*data.DepartmentDisplay, *errs.BError) {
	// 检查部门是否存在
	departmentList, err := d.departmentRepo.SearchDepartmentList(ctx, organizationCode, pcode, name)
	if err != nil {
		zap.L().Error("search department list error", zap.Error(err))
		return nil, model.SearchDepartmentListError
	}
	if len(departmentList) == 0 {
		department := &data.Department{
			OrganizationCode: organizationCode,
			Name:             name,
			CreateTime:       time.Now().UnixMilli(),
		}
		if pcode > 0 {
			department.Pcode = pcode
		}
		err := d.departmentRepo.AddDepartment(ctx, department)
		if err != nil {
			zap.L().Error("add department error", zap.Error(err))
			return nil, model.AddDepartmentError
		}
		return department.ToDisplay(), nil
	}
	return departmentList[0].ToDisplay(), nil
}

func (d *DepartmentDomain) GetDepartmentById(ctx context.Context, id int64) (*data.DepartmentDisplay, *errs.BError) {
	department, err := d.departmentRepo.GetDepartmentById(ctx, id)
	if err != nil {
		zap.L().Error("get department by id error", zap.Error(err))
		return nil, model.GetDepartmentByIdError
	}
	return department.ToDisplay(), nil
}
