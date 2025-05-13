package department_service_v1

import (
	"context"

	"github.com/jinzhu/copier"
	"test.com/project-common/encrypt"
	"test.com/project-common/errs"
	"test.com/project-grpc/department"
	"test.com/project-project/internal/domain"
	"test.com/project-project/pkg/model"
)

type DepartmentService struct {
	department.UnimplementedDepartmentServiceServer
	department *domain.DepartmentDomain
	user       *domain.UserDomain
}

func NewDepartmentService(departmentDomain *domain.DepartmentDomain, userDomain *domain.UserDomain) *DepartmentService {
	return &DepartmentService{
		department: departmentDomain,
		user:       userDomain,
	}
}

func (ds *DepartmentService) GetDepartmentList(ctx context.Context, req *department.GetDepartmentListReq) (*department.GetDepartmentListResp, error) {
	// 根据memberId获取organizationCode
	organizationCode, bErr := ds.user.GetOrganizationCodeByMemberId(ctx, req.MemberId)
	if bErr != nil {
		return nil, errs.GrpcError(bErr)
	}
	// 解密得到pcode
	var pcode int64 = 0
	var err error
	if req.Pcode != "" {
		pcode, err = encrypt.DecryptToInt64(req.Pcode, model.AESKey)
		if err != nil {
			return nil, errs.GrpcError(model.DecryptError)
		}
	}

	// 调用domain层接口获取部门列表
	departmentList, total, bErr := ds.department.GetDepartmentList(ctx, organizationCode, pcode, int(req.Page), int(req.PageSize))
	if bErr != nil {
		return nil, errs.GrpcError(bErr)
	}
	// 组织回复消息
	var departmentMessages []*department.DepartmentMessage
	copier.Copy(&departmentMessages, departmentList)
	return &department.GetDepartmentListResp{
		Total:       total,
		Departments: departmentMessages,
	}, nil
}
