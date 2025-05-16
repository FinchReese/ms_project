package domain

import (
	"context"

	"go.uber.org/zap"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/repo"
)

type ProjectMemberDomain struct {
	projectMemberRepo repo.ProjectMemberRepo
}

func NewProjectMemberDomain(projectMemberRepo repo.ProjectMemberRepo) *ProjectMemberDomain {
	return &ProjectMemberDomain{projectMemberRepo: projectMemberRepo}
}

// 根据member id和project id查询项目信息
func (p *ProjectMemberDomain) GetProjectAndMember(ctx context.Context, memberId int64, projectId int64) (*data.ProjectAndProjectMember, error) {
	projectAndMember, err := p.projectMemberRepo.GetProjectAndMember(ctx, memberId, projectId)
	if err != nil {
		zap.L().Error("get project by member id and project id error", zap.Error(err))
		return nil, err
	}
	return projectAndMember, nil
}
