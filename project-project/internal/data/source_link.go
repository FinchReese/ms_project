package data

import (
	"test.com/project-common/encrypt"
	"test.com/project-common/time_format"
	"test.com/project-project/pkg/model"

	"github.com/jinzhu/copier"
)

// SourceLink 资源关联表
type SourceLink struct {
	Id               int    `db:"id"`                // 主键
	SourceType       string `db:"source_type"`       // 资源类型
	SourceCode       int64  `db:"source_code"`       // 资源编号
	LinkType         string `db:"link_type"`         // 关联类型
	LinkCode         int64  `db:"link_code"`         // 关联编号
	OrganizationCode int64  `db:"organization_code"` // 组织编码
	CreateBy         int64  `db:"create_by"`         // 创建人
	CreateTime       string `db:"create_time"`       // 创建时间
	Sort             int    `db:"sort"`              // 排序
}

// TableName 返回表名
func (s *SourceLink) TableName() string {
	return "ms_source_link"
}

type SourceLinkDisplay struct {
	Id               int64        `json:"id"`
	Code             string       `json:"code"`
	SourceType       string       `json:"source_type"`
	SourceCode       string       `json:"source_code"`
	LinkType         string       `json:"link_type"`
	LinkCode         string       `json:"link_code"`
	OrganizationCode string       `json:"organization_code"`
	CreateBy         string       `json:"create_by"`
	CreateTime       string       `json:"create_time"`
	Sort             int          `json:"sort"`
	Title            string       `json:"title"`
	SourceDetail     SourceDetail `json:"sourceDetail"`
}

type SourceDetail struct {
	Id               int64  `json:"id"`
	Code             string `json:"code"`
	PathName         string `json:"path_name"`
	Title            string `json:"title"`
	Extension        string `json:"extension"`
	Size             int    `json:"size"`
	ObjectType       string `json:"object_type"`
	OrganizationCode string `json:"organization_code"`
	TaskCode         string `json:"task_code"`
	ProjectCode      string `json:"project_code"`
	CreateBy         string `json:"create_by"`
	CreateTime       string `json:"create_time"`
	Downloads        int    `json:"downloads"`
	Extra            string `json:"extra"`
	Deleted          int    `json:"deleted"`
	FileUrl          string `json:"file_url"`
	FileType         string `json:"file_type"`
	DeletedTime      string `json:"deleted_time"`
	ProjectName      string `json:"projectName"`
	FullName         string `json:"fullName"`
}

func (s *SourceLink) ToDisplay(f *File) *SourceLinkDisplay {
	sl := &SourceLinkDisplay{}
	copier.Copy(sl, s)
	sl.SourceDetail = SourceDetail{}
	copier.Copy(&sl.SourceDetail, f)
	sl.LinkCode, _ = encrypt.EncryptInt64(s.LinkCode, model.AESKey)
	sl.OrganizationCode, _ = encrypt.EncryptInt64(s.OrganizationCode, model.AESKey)
	sl.CreateTime = s.CreateTime
	sl.CreateBy, _ = encrypt.EncryptInt64(s.CreateBy, model.AESKey)
	sl.SourceCode, _ = encrypt.EncryptInt64(s.SourceCode, model.AESKey)
	sl.SourceDetail.OrganizationCode, _ = encrypt.EncryptInt64(f.OrganizationCode, model.AESKey)
	sl.SourceDetail.CreateBy, _ = encrypt.EncryptInt64(f.CreateBy, model.AESKey)
	sl.SourceDetail.CreateTime = time_format.ConvertMsecToString(f.CreateTime)
	sl.SourceDetail.DeletedTime = time_format.ConvertMsecToString(f.DeletedTime)
	sl.SourceDetail.TaskCode, _ = encrypt.EncryptInt64(f.TaskCode, model.AESKey)
	sl.SourceDetail.ProjectCode, _ = encrypt.EncryptInt64(f.ProjectCode, model.AESKey)
	sl.SourceDetail.FullName = f.Title
	sl.Title = f.Title
	return sl
}
