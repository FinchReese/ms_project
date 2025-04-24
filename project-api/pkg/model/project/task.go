package project

type GetTaskStageReq struct {
	ProjectCode string `form:"project_code"`
	Page        int32  `form:"page"`
	PageSize    int32  `form:"pageSize"`
}

type GetTaskStageResp struct {
	Total int64        `json:"total"`
	Page  int32        `json:"page"`
	List  []*TaskStage `json:"list"`
}

type TaskStage struct {
	Name         string `json:"name"`
	ProjectCode  string `json:"project_code"`
	Sort         int    `json:"sort"`
	Description  string `json:"description"`
	CreateTime   string `json:"create_time" copier:"CreateTimeStr"`
	Code         string `json:"code"`
	Deleted      int    `json:"deleted"`
	TasksLoading bool   `json:"tasksLoading"`
	FixedCreator bool   `json:"fixedCreator"`
	ShowTaskCard bool   `json:"showTaskCard"`
	Tasks        []int  `json:"tasks"`
	DoneTasks    []int  `json:"doneTasks"`
	UnDoneTasks  []int  `json:"unDoneTasks"`
}
