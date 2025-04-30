package data

import (
	"github.com/jinzhu/copier"
	"test.com/project-common/encrypt"
	"test.com/project-common/time_format"
	"test.com/project-project/pkg/model"
)

// TaskWorkTime 任务工时表
type TaskWorkTime struct {
	Id         int64  `db:"id"`          // 主键
	TaskCode   int64  `db:"task_code"`   // 任务ID
	MemberCode int64  `db:"member_code"` // 成员id
	CreateTime int64  `db:"create_time"` // 创建时间
	Content    string `db:"content"`     // 描述
	BeginTime  int64  `db:"begin_time"`  // 开始时间
	Num        int    `db:"num"`         // 工时
}

// TableName 返回表名
func (t *TaskWorkTime) TableName() string {
	return "ms_task_work_time"
}

type TaskWorkTimeDisplay struct {
	Id         int64
	TaskCode   string
	MemberCode string
	CreateTime string
	Content    string
	BeginTime  string
	Num        int
	Member     Member
}

func (t *TaskWorkTime) ToDisplay() *TaskWorkTimeDisplay {
	td := &TaskWorkTimeDisplay{}
	copier.Copy(td, t)
	td.MemberCode, _ = encrypt.EncryptInt64(t.MemberCode, model.AESKey)
	td.TaskCode, _ = encrypt.EncryptInt64(t.TaskCode, model.AESKey)
	td.CreateTime = time_format.ConvertMsecToString(t.CreateTime)
	td.BeginTime = time_format.ConvertMsecToString(t.BeginTime)
	return td
}
