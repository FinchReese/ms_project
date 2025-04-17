package dao

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
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
