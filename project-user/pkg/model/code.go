package model

import (
	"test.com/project-common/errs"
)

var (
	NoLegalMobile                 = errs.NewError(2001, "手机号不合法")
	CaptchaNotExist               = errs.NewError(2002, "验证码不存在或者已过期")
	CaptchaError                  = errs.NewError(2003, "验证码错误")
	EmailExist                    = errs.NewError(2004, "邮箱已存在")
	AccountExist                  = errs.NewError(2005, "账号已存在")
	MobileExist                   = errs.NewError(2006, "手机号已存在")
	QueryDbError                  = errs.NewError(2007, "查询数据库错误")
	RegisterMemberError           = errs.NewError(2008, "注册成员失败")
	RegisterOrganizationError     = errs.NewError(2009, "注册组织失败")
	AccountOrPwdError             = errs.NewError(2010, "账号或者密码错误")
	CopyMemberMsgError            = errs.NewError(2011, "复制成员信息出错")
	GetOrganizationMsgError       = errs.NewError(2012, "获取组织信息出错")
	CopyOrganizationMsgError      = errs.NewError(2013, "复制组织信息出错")
	EncryptMemberIdError          = errs.NewError(2014, "对成员id加密失败")
	EncryptOrganizationIdError    = errs.NewError(2014, "对组织id加密失败")
	VerifyTokenError              = errs.NewError(2015, "校验token失败")
	FindMemberByIdError           = errs.NewError(2016, "根据id查找成员错误")
	GetMemberFromRedisError       = errs.NewError(2017, "从Redis获取成员信息失败")
	UnmarshalMemberFromRedisError = errs.NewError(2018, "解析Redis中的成员信息失败")
	LoginTimeoutError             = errs.NewError(2019, "登录已过期，请重新登录")
	FindMembersByIdsError         = errs.NewError(2020, "根据id列表查找成员错误")
)
