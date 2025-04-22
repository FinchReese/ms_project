package repo

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
)

type ProjectRepo interface {
	SaveProject(ctx context.Context, p *data.Project, db *gorm.DB) error
}

type ProjectMemberRepo interface {
	GetProjectList(ctx context.Context, memberId int64, selectBy string, page int64, size int64) ([]*data.ProjectAndProjectMember, int64, error)
	SaveProjectMember(ctx context.Context, p *data.ProjectMember, db *gorm.DB) error
	GetProjectAndMember(ctx context.Context, memberId int64, projectId int64) (*data.ProjectAndProjectMember, error)
	IsCollectedProject(ctx context.Context, memberId int64, projectId int64) (bool, error)
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
