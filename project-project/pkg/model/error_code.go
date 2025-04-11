package model

import (
	"test.com/project-common/errs"
)

var (
	GetAllMenusError    = errs.NewError(2000, "获取菜单出错")
	GetProjectListError = errs.NewError(2001, "获取项目列表出错")
)
