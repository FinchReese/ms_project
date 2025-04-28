package repo

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
)

type ProjectRepo interface {
	SaveProject(ctx context.Context, p *data.Project, db *gorm.DB) error
	UpdateProjectDeletedState(ctx context.Context, projectId int64, deleted bool) error
	UpdateProject(ctx context.Context, project *data.Project) error
	// 根据项目id获取项目信息
	GetProjectByID(ctx context.Context, projectID int64) (*data.Project, error)
}

type ProjectMemberRepo interface {
	GetProjectList(ctx context.Context, memberId int64, selectBy string, page int64, size int64) ([]*data.ProjectAndProjectMember, int64, error)
	SaveProjectMember(ctx context.Context, p *data.ProjectMember, db *gorm.DB) error
	GetProjectAndMember(ctx context.Context, memberId int64, projectId int64) (*data.ProjectAndProjectMember, error)
	IsCollectedProject(ctx context.Context, memberId int64, projectId int64) (bool, error)
	GetProjectMemberList(ctx context.Context, projectId int64, page, pageSize int) ([]*data.ProjectMember, int64, error)
}

type ProjectTemplateRepo interface {
	GetSystemProjectTemplates(ctx context.Context, page int64, size int64) ([]data.ProjectTemplate, int64, error)
	GetCustomProjectTemplates(ctx context.Context, memId int64, page int64, size int64) ([]data.ProjectTemplate, int64, error)
	GetAllProjectTemplates(ctx context.Context, page int64, size int64) ([]data.ProjectTemplate, int64, error)
}

type TemplateTaskStageRepo interface {
	GetTaskStagesByTemplateIds(ctx context.Context, ids []int) ([]data.TemplateTaskStage, error)
}

type ProjectCollectRepo interface {
	// 收藏项目
	Collect(ctx context.Context, memberId int64, projectId int64, createTime int64) error
	// 取消收藏项目
	CancelCollect(ctx context.Context, memberId int64, projectId int64) error
}
