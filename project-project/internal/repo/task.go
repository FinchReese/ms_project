package repo

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
)

type TaskStageRepo interface {
	SaveTaskStage(ctx context.Context, ts *data.TaskStage, db *gorm.DB) error
	FindStagesByProjectId(ctx context.Context, projectId int64, page int, pageSize int) (list []*data.TaskStage, total int64, err error)
}

type TaskRepo interface {
	FindTasksByStageCode(ctx context.Context, stageCode int) (list []*data.Task, err error)
}

type TaskMemberRepo interface {
	FindTaskMembers(ctx context.Context, taskCode int64, memberCode int64) (list []*data.TaskMember, err error)
}
