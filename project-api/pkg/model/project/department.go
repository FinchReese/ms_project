package project

type GetDepartmentListReq struct {
	DepartmentCode       string `form:"departmentCode"`
	ParentDepartmentCode string `form:"parentDepartmentCode"`
	Name                 string `form:"name"`
	Page                 int64  `form:"page"`
	PageSize             int64  `form:"pageSize"`
	Pcode                string `form:"pcode"`
}

type Department struct {
	Id               int64  `json:"id"`
	Code             string `json:"code"`
	OrganizationCode string `json:"organization_code"`
	Name             string `json:"name"`
	Pcode            string `json:"pcode"`
	Path             string `json:"path"`
	CreateTime       string `json:"create_time"`
}

type GetDepartmentListResp struct {
	Total       int64         `json:"total"`
	Page        int64         `json:"page"`
	Departments []*Department `json:"departments"`
}
