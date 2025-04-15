package organization

import "test.com/project-common/time_format"

const (
	OrganizationPersion = 1
	OrganizationPublic  = 2
)

type Organization struct {
	Id          int64
	Name        string
	Avatar      string
	Description string
	MemberId    int64
	CreateTime  int64
	Personal    int32
	Address     string
	Province    int32
	City        int32
	Area        int32
}

func (o *Organization) TableName() string {
	return "ms_organization"
}

func (o *Organization) CreateTimeStr() string {
	return time_format.ConvertMsecToString(o.CreateTime)
}
