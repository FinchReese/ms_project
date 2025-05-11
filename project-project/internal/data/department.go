package data

import (
	"github.com/jinzhu/copier"
	"test.com/project-common/encrypt"
	"test.com/project-common/time_format"
	"test.com/project-project/pkg/model"
)

type Department struct {
	Id               int64
	OrganizationCode int64
	Name             string
	Sort             int
	Pcode            int64
	icon             string
	CreateTime       int64
	Path             string
}

func (*Department) TableName() string {
	return "ms_department"
}

type DepartmentDisplay struct {
	Id               int64
	Code             string
	OrganizationCode string
	Name             string
	Sort             int
	Pcode            string
	icon             string
	CreateTime       string
	Path             string
}

func (d *Department) ToDisplay() *DepartmentDisplay {
	dp := &DepartmentDisplay{}
	copier.Copy(dp, d)
	dp.Code, _ = encrypt.EncryptInt64(d.Id, model.AESKey)
	dp.CreateTime = time_format.ConvertMsecToString(d.CreateTime)
	dp.OrganizationCode, _ = encrypt.EncryptInt64(d.OrganizationCode, model.AESKey)
	if d.Pcode > 0 {
		dp.Pcode, _ = encrypt.EncryptInt64(d.Pcode, model.AESKey)
	} else {
		dp.Pcode = ""
	}
	return dp
}
