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
