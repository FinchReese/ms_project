package data

import "test.com/project-common/time_format"

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
