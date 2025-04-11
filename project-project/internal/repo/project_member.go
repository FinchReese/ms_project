package repo

import (
	"context"

	"test.com/project-project/internal/data"
)

type ProjectMemberRepo interface {
	GetProjectListByMemberId(ctx context.Context, memberId int64, page int64, size int64) ([]*data.ProjectAndProjectMember, int64, error)
}
