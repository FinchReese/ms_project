package repo

import (
	"context"
	"test.com/project-project/internal/data"
)

type ProjectNodeRepo interface {
	GetAllProjectNodeList(ctx context.Context) ([]*data.ProjectNode, error)
}
