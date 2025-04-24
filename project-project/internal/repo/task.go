package repo

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
)

type TaskStageRepo interface {
	SaveTaskStage(ctx context.Context, ts *data.TaskStage, db *gorm.DB) error
	FindStagesByProjectId(ctx context.Context, projectCode int64, page int, pageSize int) (list []*data.TaskStage, total int64, err error)
}
