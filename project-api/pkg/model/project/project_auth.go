package project

type GetProjectAuthListReq struct {
	Page     int64 `form:"page"`
	PageSize int64 `form:"pageSize"`
}

type GetProjectAuthListResp struct {
	Total int64          `json:"total"`
	Page  int64          `json:"page"`
	List  []*ProjectAuth `json:"list"`
}

type ProjectAuth struct {
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
