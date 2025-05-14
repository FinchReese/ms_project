package menu_service_v1

import (
	"context"

	"github.com/jinzhu/copier"
	"test.com/project-common/errs"
	menu "test.com/project-grpc/menu"
	"test.com/project-project/internal/domain"
)

type MenuService struct {
	menu.UnimplementedMenuServiceServer
	menuDomain *domain.MenuDomain
}

func NewMenuService(menuDomain *domain.MenuDomain) *MenuService {
	return &MenuService{menuDomain: menuDomain}
}

func (s *MenuService) GetMenuTree(ctx context.Context, req *menu.GetMenuTreeReq) (*menu.GetMenuTreeResp, error) {
	menuTree, err := s.menuDomain.GetMenuTree(ctx)
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	var menuTreeDisp []*menu.MenuMessage
	copier.Copy(&menuTreeDisp, menuTree)
	return &menu.GetMenuTreeResp{MenuTree: menuTreeDisp}, nil
}
