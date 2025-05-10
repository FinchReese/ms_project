package data

import (
	"github.com/jinzhu/copier"
	"test.com/project-common/encrypt"
	"test.com/project-common/time_format"
	"test.com/project-project/pkg/model"
)

type ProjectAuth struct {
	Id               int64  `json:"id"`
	OrganizationCode int64  `json:"organization_code"`
	Title            string `json:"title"`
	CreateAt         int64  `json:"create_at"`
	Sort             int    `json:"sort"`
	Status           int    `json:"status"`
	Desc             string `json:"desc"`
	CreateBy         int64  `json:"create_by"`
	IsDefault        int    `json:"is_default"`
	Type             string `json:"type"`
}

func (pa *ProjectAuth) TableName() string {
	return "ms_project_auth"
}

type ProjectAuthDisplay struct {
	Id               int64  `json:"id"`
	OrganizationCode string `json:"organization_code"`
	Title            string `json:"title"`
	CreateAt         string `json:"create_at"`
	Sort             int    `json:"sort"`
	Status           int    `json:"status"`
	Desc             string `json:"desc"`
	CreateBy         int64  `json:"create_by"`
	IsDefault        int    `json:"is_default"`
	Type             string `json:"type"`
	CanDelete        int    `json:"canDelete"`
}

func (pa *ProjectAuth) ToDisplay() *ProjectAuthDisplay {
	p := &ProjectAuthDisplay{}
	copier.Copy(p, pa)
	p.OrganizationCode, _ = encrypt.EncryptInt64(pa.OrganizationCode, model.AESKey)
	p.CreateAt = time_format.ConvertMsecToString(pa.CreateAt)
	if pa.Type == "admin" || pa.Type == "member" {
		//不能删除
		p.CanDelete = model.CanNotDelete
	} else {
		p.CanDelete = model.CanDelete
	}
	return p
}
