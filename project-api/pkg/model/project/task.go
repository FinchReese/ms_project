package project

type GetTaskStageReq struct {
	ProjectCode string `form:"projectCode"`
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

type Executor struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Code   string `json:"code"`
}

type DispTask struct {
	ProjectCode   string   `json:"project_code"`
	Name          string   `json:"name"`
	Pri           int      `json:"pri"`
	ExecuteStatus string   `json:"execute_status"`
	Description   string   `json:"description"`
	CreateBy      string   `json:"create_by"`
	DoneBy        string   `json:"done_by"`
	DoneTime      string   `json:"done_time"`
	CreateTime    string   `json:"create_time"`
	AssignTo      string   `json:"assign_to"`
	Deleted       int      `json:"deleted"`
	StageCode     string   `json:"stage_code"`
	TaskTag       string   `json:"task_tag"`
	Done          int      `json:"done"`
	BeginTime     string   `json:"begin_time"`
	EndTime       string   `json:"end_time"`
	RemindTime    string   `json:"remind_time"`
	Pcode         string   `json:"pcode"`
	Sort          int      `json:"sort"`
	Like          int      `json:"like"`
	Star          int      `json:"star"`
	DeletedTime   string   `json:"deleted_time"`
	Private       int      `json:"private"`
	IdNum         int      `json:"id_num"`
	Path          string   `json:"path"`
	Schedule      int      `json:"schedule"`
	VersionCode   string   `json:"version_code"`
	FeaturesCode  string   `json:"features_code"`
	WorkTime      int      `json:"work_time"`
	Status        int      `json:"status"`
	Code          string   `json:"code"`
	CanRead       int      `json:"canRead"`
	HasUnDone     int      `json:"hasUnDone"`
	ParentDone    int      `json:"parentDone"`
	HasComment    int      `json:"hasComment"`
	HasSource     int      `json:"hasSource"`
	Executor      Executor `json:"executor"`
	PriText       string   `json:"priText"`
	StatusText    string   `json:"statusText"`
	Liked         int      `json:"liked"`
	Stared        int      `json:"stared"`
	Tags          []int    `json:"tags"`
	ChildCount    []int    `json:"childCount"`
}

type SaveTaskReq struct {
	Name        string `form:"name"`
	StageCode   string `form:"stage_code"`
	ProjectCode string `form:"project_code"`
	AssignTo    string `form:"assign_to"`
}

type SaveTaskResp struct {
	ProjectCode   string   `json:"project_code"`
	Name          string   `json:"name"`
	Pri           int      `json:"pri"`
	ExecuteStatus string   `json:"execute_status"`
	Description   string   `json:"description"`
	CreateBy      string   `json:"create_by"`
	DoneBy        string   `json:"done_by"`
	DoneTime      string   `json:"done_time"`
	CreateTime    string   `json:"create_time"`
	AssignTo      string   `json:"assign_to"`
	Deleted       int      `json:"deleted"`
	StageCode     string   `json:"stage_code"`
	TaskTag       string   `json:"task_tag"`
	Done          int      `json:"done"`
	BeginTime     string   `json:"begin_time"`
	EndTime       string   `json:"end_time"`
	RemindTime    string   `json:"remind_time"`
	Pcode         string   `json:"pcode"`
	Sort          int      `json:"sort"`
	Like          int      `json:"like"`
	Star          int      `json:"star"`
	DeletedTime   string   `json:"deleted_time"`
	Private       int      `json:"private"`
	IdNum         int      `json:"id_num"`
	Path          string   `json:"path"`
	Schedule      int      `json:"schedule"`
	VersionCode   string   `json:"version_code"`
	FeaturesCode  string   `json:"features_code"`
	WorkTime      int      `json:"work_time"`
	Status        int      `json:"status"`
	Code          string   `json:"code"`
	CanRead       int      `json:"canRead"`
	HasUnDone     int      `json:"hasUnDone"`
	ParentDone    int      `json:"parentDone"`
	HasComment    int      `json:"hasComment"`
	HasSource     int      `json:"hasSource"`
	Executor      Executor `json:"executor"`
	PriText       string   `json:"priText"`
	StatusText    string   `json:"statusText"`
	Liked         int      `json:"liked"`
	Stared        int      `json:"stared"`
	Tags          []int    `json:"tags"`
	ChildCount    []int    `json:"childCount"`
}

type MoveTaskReq struct {
	PreTaskCode  string `form:"preTaskCode"`
	NextTaskCode string `form:"nextTaskCode"`
	ToStageCode  string `form:"toStageCode"`
}

type GetTaskListByTypeReq struct {
	TaskType int32 `form:"taskType"`
	Done     int   `form:"type"`
	Page     int   `form:"page"`
	PageSize int   `form:"pageSize"`
}

type ProjectInfo struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type MyTaskDisplay struct {
	ProjectCode        string      `json:"project_code"`
	Name               string      `json:"name"`
	Pri                int         `json:"pri"`
	ExecuteStatus      string      `json:"execute_status"`
	Description        string      `json:"description"`
	CreateBy           string      `json:"create_by"`
	DoneBy             string      `json:"done_by"`
	DoneTime           string      `json:"done_time"`
	CreateTime         string      `json:"create_time"`
	AssignTo           string      `json:"assign_to"`
	Deleted            int         `json:"deleted"`
	StageCode          string      `json:"stage_code"`
	TaskTag            string      `json:"task_tag"`
	Done               int         `json:"done"`
	BeginTime          string      `json:"begin_time"`
	EndTime            string      `json:"end_time"`
	RemindTime         string      `json:"remind_time"`
	Pcode              string      `json:"pcode"`
	Sort               int         `json:"sort"`
	Like               int         `json:"like"`
	Star               int         `json:"star"`
	DeletedTime        string      `json:"deleted_time"`
	Private            int         `json:"private"`
	IdNum              int         `json:"id_num"`
	Path               string      `json:"path"`
	Schedule           int         `json:"schedule"`
	VersionCode        string      `json:"version_code"`
	FeaturesCode       string      `json:"features_code"`
	WorkTime           int         `json:"work_time"`
	Status             int         `json:"status"`
	Code               string      `json:"code"`
	ProjectName        string      `json:"project_name"`
	Cover              string      `json:"cover"`
	AccessControlType  string      `json:"access_control_type"`
	WhiteList          string      `json:"white_list"`
	Order              int         `json:"order"`
	TemplateCode       string      `json:"template_code"`
	OrganizationCode   string      `json:"organization_code"`
	Prefix             string      `json:"prefix"`
	OpenPrefix         int         `json:"open_prefix"`
	Archive            int         `json:"archive"`
	ArchiveTime        string      `json:"archive_time"`
	OpenBeginTime      int         `json:"open_begin_time"`
	OpenTaskPrivate    int         `json:"open_task_private"`
	TaskBoardTheme     string      `json:"task_board_theme"`
	AutoUpdateSchedule int         `json:"auto_update_schedule"`
	HasUnDone          int         `json:"hasUnDone"`
	ParentDone         int         `json:"parentDone"`
	PriText            string      `json:"priText"`
	Executor           Executor    `json:"executor"`
	ProjectInfo        ProjectInfo `json:"projectInfo"`
}

type GetTaskListByTypeResp struct {
	Total int64            `json:"total"`
	List  []*MyTaskDisplay `json:"list"`
}

type GetTaskDetailReq struct {
	TaskCode string `form:"taskCode"`
}

type GetTaskDetailResp struct {
	ProjectCode   string   `json:"project_code"`
	Name          string   `json:"name"`
	Pri           int      `json:"pri"`
	ExecuteStatus string   `json:"execute_status"`
	Description   string   `json:"description"`
	CreateBy      string   `json:"create_by"`
	DoneBy        string   `json:"done_by"`
	DoneTime      string   `json:"done_time"`
	CreateTime    string   `json:"create_time"`
	AssignTo      string   `json:"assign_to"`
	Deleted       int      `json:"deleted"`
	StageCode     string   `json:"stage_code"`
	TaskTag       string   `json:"task_tag"`
	Done          int      `json:"done"`
	BeginTime     string   `json:"begin_time"`
	EndTime       string   `json:"end_time"`
	RemindTime    string   `json:"remind_time"`
	Pcode         string   `json:"pcode"`
	Sort          int      `json:"sort"`
	Like          int      `json:"like"`
	Star          int      `json:"star"`
	DeletedTime   string   `json:"deleted_time"`
	Private       int      `json:"private"`
	IdNum         int      `json:"id_num"`
	Path          string   `json:"path"`
	Schedule      int      `json:"schedule"`
	VersionCode   string   `json:"version_code"`
	FeaturesCode  string   `json:"features_code"`
	WorkTime      int      `json:"work_time"`
	Status        int      `json:"status"`
	Code          string   `json:"code"`
	CanRead       int      `json:"canRead"`
	HasUnDone     int      `json:"hasUnDone"`
	ParentDone    int      `json:"parentDone"`
	HasComment    int      `json:"hasComment"`
	HasSource     int      `json:"hasSource"`
	Executor      Executor `json:"executor"`
	PriText       string   `json:"priText"`
	StatusText    string   `json:"statusText"`
	Liked         int      `json:"liked"`
	Stared        int      `json:"stared"`
	Tags          []int    `json:"tags"`
	ChildCount    []int    `json:"childCount"`
	ProjectName   string   `json:"projectName"`
	StageName     string   `json:"stageName"`
}

type GetTaskMemberListReq struct {
	TaskCode string `form:"taskCode"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
}

type GetTaskMemberListResp struct {
	Total int64         `json:"total"`
	Page  int           `json:"page"`
	List  []*TaskMember `json:"list"`
}

type TaskMember struct {
	Id                int64  `json:"id"`
	Name              string `json:"name"`
	Avatar            string `json:"avatar"`
	Code              string `json:"code"`
	MembarAccountCode string `json:"membar_account_code"`
	IsExecutor        int    `json:"is_executor"`
	IsOwner           int    `json:"is_owner"`
}

type GetTaskLogListReq struct {
	TaskCode string `form:"taskCode"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	All      int    `form:"all"`
	Comment  int    `form:"comment"`
}

type GetTaskLogListResp struct {
	Total int64      `json:"total"`
	Page  int        `json:"page"`
	List  []*TaskLog `json:"list"`
}

type TaskLog struct {
	Id           int64  `json:"id"`
	MemberCode   string `json:"member_code"`
	Content      string `json:"content"`
	Remark       string `json:"remark"`
	Type         string `json:"type"`
	CreateTime   string `json:"create_time"`
	SourceCode   string `json:"source_code"`
	ActionType   string `json:"action_type"`
	ToMemberCode string `json:"to_member_code"`
	IsComment    int    `json:"is_comment"`
	ProjectCode  string `json:"project_code"`
	Icon         string `json:"icon"`
	IsRobot      int    `json:"is_robot"`
	Member       Member `json:"member"`
}

type Member struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Avatar string `json:"avatar"`
}

type GetTaskWorkTimeListReq struct {
	TaskCode string `form:"taskCode"`
}

type TaskWorkTime struct {
	Id         int64  `json:"id"`
	TaskCode   string `json:"task_code"`
	MemberCode string `json:"member_code"`
	CreateTime string `json:"create_time"`
	Content    string `json:"content"`
	BeginTime  string `json:"begin_time"`
	Num        int    `json:"num"`
	Code       string `json:"code"`
	Member     Member `json:"member"`
}

type SaveTaskWorkTimeReq struct {
	TaskCode  string `form:"taskCode"`
	Content   string `form:"content"`
	BeginTime string `form:"beginTime"`
	Num       int    `form:"num"`
}

type UploadFileReq struct {
	TaskCode         string `form:"taskCode"`
	ProjectCode      string `form:"projectCode"`
	ProjectName      string `form:"projectName"`
	TotalChunks      int    `form:"totalChunks"`
	RelativePath     string `form:"relativePath"`
	Filename         string `form:"filename"`
	ChunkNumber      int    `form:"chunkNumber"`
	ChunkSize        int    `form:"chunkSize"`
	CurrentChunkSize int    `form:"currentChunkSize"`
	TotalSize        int    `form:"totalSize"`
	Identifier       string `form:"identifier"`
}

type UploadFileResp struct {
	File        string `json:"file"`
	Hash        string `json:"hash"`
	Key         string `json:"key"`
	Url         string `json:"url"`
	ProjectName string `json:"projectName"`
}

type GetUserProjectLogListReq struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

type ProjectLog struct {
	Content      string `json:"content"`
	Remark       string `json:"remark"`
	CreateTime   string `json:"create_time"`
	SourceCode   string `json:"source_code"`
	IsComment    int    `json:"is_comment"`
	ProjectCode  string `json:"project_code"`
	ProjectName  string `json:"project_name"`
	MemberAvatar string `json:"member_avatar"`
	MemberName   string `json:"member_name"`
	TaskName     string `json:"task_name"`
}
