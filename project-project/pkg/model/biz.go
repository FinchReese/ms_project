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
