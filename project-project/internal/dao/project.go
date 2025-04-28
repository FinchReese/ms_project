package dao

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
	"test.com/project-project/pkg/model"
)

type ProjectDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewProjectDAO() *ProjectDAO {
	return &ProjectDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (pd *ProjectDAO) SaveProject(ctx context.Context, p *data.Project, db *gorm.DB) error {
	return db.Session(&gorm.Session{Context: ctx}).Create(p).Error
}

func (pd *ProjectDAO) UpdateProjectDeletedState(ctx context.Context, projectId int64, deleted bool) error {
	deletedState := model.NotDeleted
	if deleted {
		deletedState = model.Deleted
	}
	return pd.conn.Db.Session(&gorm.Session{Context: ctx}).Model(&data.Project{}).
		Where("id = ?", projectId).
		Update("deleted", deletedState).Error
}

func (pd *ProjectDAO) UpdateProject(ctx context.Context, project *data.Project) error {
	return pd.conn.Db.Session(&gorm.Session{Context: ctx}).Model(&data.Project{}).Where("id=?", project.Id).
		Updates(project).Error
}

// GetProjectByID 根据项目id获取项目信息
func (pd *ProjectDAO) GetProjectByID(ctx context.Context, projectID int64) (*data.Project, error) {
	var project data.Project
	err := pd.conn.Db.Session(&gorm.Session{Context: ctx}).
		Where("id = ?", projectID).
		First(&project).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

// GetProjectsByIds 根据项目id列表获取项目信息
func (pd *ProjectDAO) GetProjectsByIds(ctx context.Context, projectIds []int64) ([]*data.Project, error) {
	var projects []*data.Project
	err := pd.conn.Db.Session(&gorm.Session{Context: ctx}).Where("id IN ?", projectIds).Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}
