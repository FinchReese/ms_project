package model

import (
	"test.com/project-common/errs"
)

var (
	NoLegalMobile             = errs.NewError(2001, "手机号不合法")
	CaptchaNotExist           = errs.NewError(2002, "验证码不存在或者已过期")
	CaptchaError              = errs.NewError(2003, "验证码错误")
	EmailExist                = errs.NewError(2004, "邮箱已存在")
	AccountExist              = errs.NewError(2005, "账号已存在")
	MobileExist               = errs.NewError(2006, "手机号已存在")
	QueryDbError              = errs.NewError(2007, "查询数据库错误")
	RegisterMemberError       = errs.NewError(2008, "注册成员失败")
	RegisterOrganizationError = errs.NewError(2009, "注册组织失败")
)
