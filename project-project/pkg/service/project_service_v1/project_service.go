package project_service_v1

import (
	"context"

	"github.com/jinzhu/copier"
	"test.com/project-common/errs"
	"test.com/project-grpc/project"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/repo"
	"test.com/project-project/pkg/model"
)

type ProjectService struct {
	project.UnimplementedProjectServiceServer
	menuRepo repo.MenuRepo
}

func NewProjectService(mr repo.MenuRepo) *ProjectService {
	return &ProjectService{
		menuRepo: mr,
	}
}

func (p *ProjectService) Index(ctx context.Context, req *project.IndexMessage) (resp *project.IndexResponse, err error) {
	// 数据库查询所有菜单信息
	menuList, err := p.menuRepo.GetAllMenus(context.TODO())
	if err != nil {
		return nil, errs.GrpcError(model.GetAllMenusError)
	}
	menuTree := data.ConvertMenuListToTreeList(menuList)
	var menus []*project.MenuMessage
	copier.Copy(&menus, menuTree)
	return &project.IndexResponse{Menus: menus}, nil
}
