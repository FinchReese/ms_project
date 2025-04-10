package model

import (
	"test.com/project-common/errs"
)

var (
	GetAllMenusError = errs.NewError(2000, "获取菜单失败")
)
