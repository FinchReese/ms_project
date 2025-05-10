package model

const AESKey = "sdfgyrhgbxcdgryfhgywertd"

const (
	NotDeleted = iota
	Deleted
)

const (
	NoArchive = iota
	Archive
)

const (
	Open = iota
	Private
	Custom
)

const (
	Default = "default"
	Simple  = "simple"
)

const (
	NotCollected = iota
	Collected
)

const (
	CanNotRead = iota
	CanRead
)

const (
	NoExecutor = iota
	Executor
)

const (
	NotDone = 0 // 0表示未完成
	Done    = 1 // 1表示已完成
)

const (
	TaskTypeAssignedTo = 1 // 分配的任务
	TaskTypeInvolved   = 2 // 参与的任务
	TaskTypeCreatedBy  = 3 // 创建的任务
)

// link type
const (
	LinkTypeTask = "task"
)

// log type
const (
	NotCommentLog = 0 // 非评论日志
	CommentLog    = 1 // 评论日志
)

// 是否能删除
const (
	CanDelete    = 1 // 能删除
	CanNotDelete = 0 // 不能删除
)

// 账号状态
const (
	AccountStatusDisabled = 0 // 禁用
	AccountStatusEnabled  = 1 // 启用
)

// 查询成员账号类型
const (
	QueryMemberAccountTypeEnableAccount      = 1 // 使用中
	QueryMemberAccountTypeNullDepartmentCode = 2 // 部门id为空
	QueryMemberAccountTypeDisableAccount     = 3 // 禁用
	QueryMemberAccountTypeByDepartmentCode   = 4 // 使用中并根据部门id查询
)
