package data

import (
	"test.com/project-common/encrypt"
	"test.com/project-common/time_format"
	"test.com/project-project/pkg/model"
)

type Project struct {
	Id                 int64
	Cover              string
	Name               string
	Description        string
	AccessControlType  int
	WhiteList          string
	Order              int
	Deleted            int
	TemplateCode       int
	Schedule           float64
	CreateTime         int64
	OrganizationCode   int64
	DeletedTime        string
	Private            int
	Prefix             string
	OpenPrefix         int
	Archive            int
	ArchiveTime        int64
	OpenBeginTime      int
	OpenTaskPrivate    int
	TaskBoardTheme     string
	BeginTime          int64
	EndTime            int64
	AutoUpdateSchedule int
}

func (p *Project) TableName() string {
	return "ms_project"
}

func (m *Project) GetAccessControlType() string {
	if m.AccessControlType == 0 {
		return "open"
	}
	if m.AccessControlType == 1 {
		return "private"
	}
	if m.AccessControlType == 2 {
		return "custom"
	}
	return ""
}

type ProjectTemplate struct {
	Id               int
	Name             string
	Description      string
	Sort             int
	CreateTime       int64
	OrganizationCode int64
	Cover            string
	MemberCode       int64
	IsSystem         int
}

func (*ProjectTemplate) TableName() string {
	return "ms_project_template"
}

type CompleteProjectTemplate struct {
	Id               int
	Name             string
	Description      string
	Sort             int
	CreateTime       string
	OrganizationCode string
	Cover            string
	MemberCode       string
	IsSystem         int
	TaskStages       []*TaskStagesOnlyName
	Code             string
}

func (pt *ProjectTemplate) Convert(taskStages []*TaskStagesOnlyName) *CompleteProjectTemplate {
	organizationCode, _ := encrypt.EncryptInt64(pt.OrganizationCode, model.AESKey)
	memberCode, _ := encrypt.EncryptInt64(pt.MemberCode, model.AESKey)
	code, _ := encrypt.EncryptInt64(int64(pt.Id), model.AESKey)
	pta := &CompleteProjectTemplate{
		Id:               pt.Id,
		Name:             pt.Name,
		Description:      pt.Description,
		Sort:             pt.Sort,
		CreateTime:       time_format.ConvertMsecToString(pt.CreateTime),
		OrganizationCode: organizationCode,
		Cover:            pt.Cover,
		MemberCode:       memberCode,
		IsSystem:         pt.IsSystem,
		TaskStages:       taskStages,
		Code:             code,
	}
	return pta
}
func ToProjectTemplateIds(pts []ProjectTemplate) []int {
	var ids []int
	for _, v := range pts {
		ids = append(ids, v.Id)
	}
	return ids
}

type ProjectCollection struct {
	Id          int64
	ProjectCode int64
	MemberCode  int64
	CreateTime  int64
}

func (*ProjectCollection) TableName() string {
	return "ms_project_collection"
}
