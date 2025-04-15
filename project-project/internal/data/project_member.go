package data

import (
	"go.uber.org/zap"
	"test.com/project-common/encrypt"
	"test.com/project-common/time_format"
)

const (
	aesKey = "sdfgyrhgbxcdgryfhgywertd"
)

type ProjectMember struct {
	Id          int64
	ProjectCode int64
	MemberCode  int64
	JoinTime    int64
	IsOwner     int64
	Authorize   string
}

func (*ProjectMember) TableName() string {
	return "ms_project_member"
}

type ProjectAndProjectMember struct {
	Project
	ProjectCode int64
	MemberCode  int64
	JoinTime    int64
	IsOwner     int64
	Authorize   string
}

func (p *ProjectAndProjectMember) AccessControlTypeStr() string {
	var str string
	switch p.AccessControlType {
	case 0:
		str = "open"
	case 1:
		str = "private"
	case 2:
		str = "custom"
	default:
		str = ""
	}
	return str
}

func (p *ProjectAndProjectMember) EncryptedOrganizationCode() string {
	res, err := encrypt.EncryptInt64(p.OrganizationCode, aesKey)
	if err != nil {
		zap.L().Error("encrypt organizationCode err", zap.Error(err))
		return ""
	}
	return res
}

func (p *ProjectAndProjectMember) JoinTimeStr() string {
	return time_format.ConvertMsecToString(p.JoinTime)
}

func (p *ProjectAndProjectMember) CreateTimeStr() string {
	return time_format.ConvertMsecToString(p.CreateTime)
}
