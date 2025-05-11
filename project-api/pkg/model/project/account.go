package project

type GetAccountListReq struct {
	Page           int    `form:"page"`
	PageSize       int    `form:"pageSize"`
	SearchType     int    `form:"searchType"`
	DepartmentCode string `form:"departmentCode"`
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

type MemberAccount struct {
	Id                int64    `json:"id"`
	Code              string   `json:"code"`
	MemberCode        string   `json:"member_code"`
	OrganizationCode  string   `json:"organization_code"`
	DepartmentCode    string   `json:"department_code"`
	Authorize         string   `json:"authorize"`
	IsOwner           int      `json:"is_owner"`
	Name              string   `json:"name"`
	Mobile            string   `json:"mobile"`
	Email             string   `json:"email"`
	CreateTime        string   `json:"create_time"`
	LastLoginTime     string   `json:"last_login_time"`
	Status            int      `json:"status"`
	Description       string   `json:"description"`
	Avatar            string   `json:"avatar"`
	Position          string   `json:"position"`
	Department        string   `json:"department"`
	MembarAccountCode string   `json:"membar_account_code"`
	Departments       string   `json:"departments"`
	StatusText        string   `json:"statusText"`
	AuthorizeArr      []string `json:"authorizeArr"`
}

type GetAccountListResp struct {
	ProjectAuthList   []*ProjectAuth   `json:"authList"`
	MemberAccountList []*MemberAccount `json:"list"`
	Total             int64            `json:"total"`
	Page              int              `json:"page"`
}
