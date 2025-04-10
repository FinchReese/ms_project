package repo

import (
	"context"

	"test.com/project-project/internal/data"
)

type MenuRepo interface {
	GetAllMenus(ctx context.Context) (menuList []*data.ProjectMenu, err error)
}
