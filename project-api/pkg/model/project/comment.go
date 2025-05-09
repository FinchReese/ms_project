package project

type CreateCommentReq struct {
	TaskCode string   `form:"taskCode"`
	Comment  string   `form:"comment"`
	Mentions []string `form:"mentions"`
}
