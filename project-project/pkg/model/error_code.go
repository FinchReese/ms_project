package model

import (
	"test.com/project-common/errs"
)

var (
	GetAllMenusError               = errs.NewError(2000, "获取菜单出错")
	GetProjectListError            = errs.NewError(2001, "获取项目列表出错")
	InvalidViewType                = errs.NewError(2002, "错误的view类型")
	QueryProjectTemplateError      = errs.NewError(2003, "查询项目模板出错")
	QueryTemplateTaskStagesError   = errs.NewError(2004, "查询模板任务步骤出错")
	SaveProjectError               = errs.NewError(2005, "保存项目出错")
	SaveProjectMembertError        = errs.NewError(2006, "保存项目成员出错")
	EncryptProjectIdError          = errs.NewError(2007, "对项目id加密出错")
	EncryptOrganizationIdError     = errs.NewError(2008, "对项目id加密出错")
	GetOrganizationError           = errs.NewError(2009, "获取组织信息出错")
	GetProjectAndMemberError       = errs.NewError(2010, "获取项目和成员信息出错")
	GetMemberByIdError             = errs.NewError(2011, "根据id获取成员信息出错")
	GetProjectCollectedStateError  = errs.NewError(2012, "获取项目收藏状态出错")
	DecryptProjectCodeError        = errs.NewError(2013, "解密项目编码出错")
	ParseProjectIdError            = errs.NewError(2014, "解析项目ID出错")
	CollectProjectError            = errs.NewError(2015, "收藏项目出错")
	CancelCollectProjectError      = errs.NewError(2016, "取消收藏项目出错")
	InvalidCollectType             = errs.NewError(2017, "无效的收藏类型")
	UpdateProjectDeletedStateError = errs.NewError(2018, "更新项目删除状态出错")
	UpdateProjectError             = errs.NewError(2019, "更新项目出错")
	GetTaskStagesError             = errs.NewError(2020, "获取任务步骤出错")
	CopyTaskStageListError         = errs.NewError(2021, "复制任务步骤列表出错")
	EncryptTaskStageIdError        = errs.NewError(2022, "加密任务步骤id出错")
	TemplateTaskStagesNotFound     = errs.NewError(2023, "模板任务步骤不存在")
	SaveTaskStageError             = errs.NewError(2024, "保存任务步骤出错")
)
