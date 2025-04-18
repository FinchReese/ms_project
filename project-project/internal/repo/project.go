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
