package data

import (
	"github.com/jinzhu/copier"
	"test.com/project-common/encrypt"
	"test.com/project-common/time_format"
	"test.com/project-grpc/task"
	"test.com/project-project/pkg/model"
)

type TemplateTaskStage struct {
	Id                  int
	Name                string
	ProjectTemplateCode int
	CreateTime          int64
	Sort                int
}

func (*TemplateTaskStage) TableName() string {
	return "ms_task_stages_template"
}

type TaskStagesOnlyName struct {
	Name string
}

func CovertProjectMap(tsts []TemplateTaskStage) map[int][]*TaskStagesOnlyName {
	var tss = make(map[int][]*TaskStagesOnlyName)
	for _, v := range tsts {
		ts := &TaskStagesOnlyName{}
		ts.Name = v.Name
		tss[v.ProjectTemplateCode] = append(tss[v.ProjectTemplateCode], ts)
	}
	return tss
}

type TaskStage struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ProjectCode int64  `json:"project_code"`
	Sort        int    `json:"sort"`
	Description string `json:"description"`
	CreateTime  int64  `json:"create_time"`
	Deleted     int    `json:"deleted"`
}

func (ts *TaskStage) TableName() string {
	return "ms_task_stage"
}

func (ts *TaskStage) CreateTimeStr() string {
	return time_format.ConvertMsecToString(ts.CreateTime)
}

type Task struct {
	Id            int64  `json:"id"`
	ProjectCode   int64  `json:"project_code"`
	Name          string `json:"name"`
	Pri           int    `json:"pri"`            // 紧急程度
	ExecuteStatus int    `json:"execute_status"` // 执行状态
	Description   string `json:"description"`    // 详情
	CreateBy      int64  `json:"create_by"`      // 创建人
	DoneBy        int64  `json:"done_by"`        // 完成人
	DoneTime      int64  `json:"done_time"`      // 完成时间
	CreateTime    int64  `json:"create_time"`    // 创建日期
	AssignTo      int64  `json:"assign_to"`      // 指派给谁
	Deleted       int    `json:"deleted"`        // 回收站
	StageCode     int    `json:"stage_code"`     // 任务列表
	TaskTag       string `json:"task_tag"`       // 任务标签
	Done          int    `json:"done"`           // 是否完成
	BeginTime     int64  `json:"begin_time"`     // 开始时间
	EndTime       int64  `json:"end_time"`       // 截止时间
	RemindTime    int64  `json:"remind_time"`    // 提醒时间
	Pcode         int64  `json:"pcode"`          // 父任务id
	Sort          int    `json:"sort"`           // 排序
	Like          int    `json:"like"`           // 点赞数
	Star          int    `json:"star"`           // 收藏数
	DeletedTime   int64  `json:"deleted_time"`   // 删除时间
	Private       int    `json:"private"`        // 是否隐私模式
	IdNum         int    `json:"id_num"`         // 任务id编号
	Path          string `json:"path"`           // 上级任务路径
	Schedule      int    `json:"schedule"`       // 进度百分比
	VersionCode   int64  `json:"version_code"`   // 版本id
	FeaturesCode  int64  `json:"features_code"`  // 版本库id
	WorkTime      int    `json:"work_time"`      // 预估工时
	Status        int    `json:"status"`         // 执行状态
}

// TableName 指定表名
func (*Task) TableName() string {
	return "ms_task"
}

func (t *Task) ToDisplayTask() *task.Task {
	dispTask := &task.Task{}
	copier.Copy(&dispTask, t)
	dispTask.CreateTime = time_format.ConvertMsecToString(t.CreateTime)
	dispTask.DoneTime = time_format.ConvertMsecToString(t.DoneTime)
	dispTask.BeginTime = time_format.ConvertMsecToString(t.BeginTime)
	dispTask.EndTime = time_format.ConvertMsecToString(t.EndTime)
	dispTask.RemindTime = time_format.ConvertMsecToString(t.RemindTime)
	dispTask.DeletedTime = time_format.ConvertMsecToString(t.DeletedTime)
	dispTask.CreateBy, _ = encrypt.EncryptInt64(t.CreateBy, model.AESKey)
	dispTask.ProjectCode, _ = encrypt.EncryptInt64(t.ProjectCode, model.AESKey)
	dispTask.DoneBy, _ = encrypt.EncryptInt64(t.DoneBy, model.AESKey)
	dispTask.AssignTo, _ = encrypt.EncryptInt64(t.AssignTo, model.AESKey)
	dispTask.StageCode, _ = encrypt.EncryptInt64(int64(t.StageCode), model.AESKey)
	dispTask.Pcode, _ = encrypt.EncryptInt64(t.Pcode, model.AESKey)
	dispTask.VersionCode, _ = encrypt.EncryptInt64(t.VersionCode, model.AESKey)
	dispTask.FeaturesCode, _ = encrypt.EncryptInt64(t.FeaturesCode, model.AESKey)
	dispTask.ExecuteStatus = t.GetExecuteStatusStr()
	dispTask.Code, _ = encrypt.EncryptInt64(t.Id, model.AESKey)
	dispTask.CanRead = 1
	return dispTask
}

type Executor struct {
	Name   string
	Avatar string
	Code   string
}

type MyTaskDisplay struct {
	Id                 int64
	ProjectCode        string
	Name               string
	Pri                int
	ExecuteStatus      string
	Description        string
	CreateBy           string
	DoneBy             string
	DoneTime           string
	CreateTime         string
	AssignTo           string
	Deleted            int
	StageCode          string
	TaskTag            string
	Done               int
	BeginTime          string
	EndTime            string
	RemindTime         string
	Pcode              string
	Sort               int
	Like               int
	Star               int
	DeletedTime        string
	Private            int
	IdNum              int
	Path               string
	Schedule           int
	VersionCode        string
	FeaturesCode       string
	WorkTime           int
	Status             int
	Code               string
	Cover              string `json:"cover"`
	AccessControlType  string `json:"access_control_type"`
	WhiteList          string `json:"white_list"`
	Order              int    `json:"order"`
	TemplateCode       string `json:"template_code"`
	OrganizationCode   string `json:"organization_code"`
	Prefix             string `json:"prefix"`
	OpenPrefix         int    `json:"open_prefix"`
	Archive            int    `json:"archive"`
	ArchiveTime        string `json:"archive_time"`
	OpenBeginTime      int    `json:"open_begin_time"`
	OpenTaskPrivate    int    `json:"open_task_private"`
	TaskBoardTheme     string `json:"task_board_theme"`
	AutoUpdateSchedule int    `json:"auto_update_schedule"`
	HasUnDone          int    `json:"hasUnDone"`
	ParentDone         int    `json:"parentDone"`
	PriText            string `json:"priText"`
	ProjectName        string
	Executor           *Executor
}

func (t *Task) ToMyTaskDisplay(p *Project, name string, avatar string) *MyTaskDisplay {
	td := &MyTaskDisplay{}
	copier.Copy(td, p)
	copier.Copy(td, t)
	td.Executor = &Executor{
		Name:   name,
		Avatar: avatar,
	}
	td.ProjectName = p.Name
	td.CreateTime = time_format.ConvertMsecToString(t.CreateTime)
	td.DoneTime = time_format.ConvertMsecToString(t.DoneTime)
	td.BeginTime = time_format.ConvertMsecToString(t.BeginTime)
	td.EndTime = time_format.ConvertMsecToString(t.EndTime)
	td.RemindTime = time_format.ConvertMsecToString(t.RemindTime)
	td.DeletedTime = time_format.ConvertMsecToString(t.DeletedTime)
	td.CreateBy, _ = encrypt.EncryptInt64(t.CreateBy, model.AESKey)
	td.ProjectCode, _ = encrypt.EncryptInt64(t.ProjectCode, model.AESKey)
	td.DoneBy, _ = encrypt.EncryptInt64(t.DoneBy, model.AESKey)
	td.AssignTo, _ = encrypt.EncryptInt64(t.AssignTo, model.AESKey)
	td.StageCode, _ = encrypt.EncryptInt64(int64(t.StageCode), model.AESKey)
	td.Pcode, _ = encrypt.EncryptInt64(t.Pcode, model.AESKey)
	td.VersionCode, _ = encrypt.EncryptInt64(t.VersionCode, model.AESKey)
	td.FeaturesCode, _ = encrypt.EncryptInt64(t.FeaturesCode, model.AESKey)
	td.ExecuteStatus = t.GetExecuteStatusStr()
	td.Code, _ = encrypt.EncryptInt64(t.Id, model.AESKey)
	td.AccessControlType = p.GetAccessControlType()
	td.ArchiveTime = time_format.ConvertMsecToString(p.ArchiveTime)
	td.TemplateCode, _ = encrypt.EncryptInt64(int64(p.TemplateCode), model.AESKey)
	td.OrganizationCode, _ = encrypt.EncryptInt64(p.OrganizationCode, model.AESKey)
	return td
}

const (
	Wait = iota
	Doing
	Done
	Pause
	Cancel
	Closed
)

func (t *Task) GetExecuteStatusStr() string {
	status := t.ExecuteStatus
	if status == Wait {
		return "wait"
	}
	if status == Doing {
		return "doing"
	}
	if status == Done {
		return "done"
	}
	if status == Pause {
		return "pause"
	}
	if status == Cancel {
		return "cancel"
	}
	if status == Closed {
		return "closed"
	}
	return ""
}

type TaskMember struct {
	Id         int64 `json:"id"`
	TaskCode   int64 `json:"task_code"`   // 任务ID
	IsExecutor int   `json:"is_executor"` // 执行者
	MemberCode int64 `json:"member_code"` // 成员id
	JoinTime   int64 `json:"join_time"`   // 加入时间
	IsOwner    int   `json:"is_owner"`    // 是否创建人
}

// TableName 指定表名
func (*TaskMember) TableName() string {
	return "ms_task_member"
}
