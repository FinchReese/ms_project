package domain

import (
	"context"

	"test.com/project-common/errs"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/repo"
	"test.com/project-project/pkg/model"
)

type MenuDomain struct {
	menuRepo repo.MenuRepo
}

func NewMenuDomain(menuRepo repo.MenuRepo) *MenuDomain {
	return &MenuDomain{menuRepo: menuRepo}
}

func (md *MenuDomain) GetMenuTree(ctx context.Context) ([]*data.ProjectMenuNode, *errs.BError) {
	menuList, err := md.menuRepo.GetAllMenus(ctx)
	if err != nil {
		return nil, model.GetAllMenusError
	}
	menuTree := data.ConvertMenuListToTreeList(menuList)
	return menuTree, nil
}
