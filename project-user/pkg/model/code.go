package model

import (
	"test.com/project-common/errs"
)

var (
	NoLegalMobile = errs.NewError(2001, "手机号不合法")
)
