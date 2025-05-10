package repo

import (
	"context"

	"test.com/project-project/internal/data"
)

type ProjectAuthRepo interface {
	GetProjectAuthList(ctx context.Context, organizationCode int64) ([]*data.ProjectAuth, error)
}
