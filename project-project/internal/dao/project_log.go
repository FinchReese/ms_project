package dao

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

type ProjectLogDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewProjectLogDAO() *ProjectLogDAO {
	return &ProjectLogDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (pld *ProjectLogDAO) CreateProjectLog(ctx context.Context, log *data.ProjectLog) error {
	return pld.conn.Db.Session(&gorm.Session{Context: ctx}).Create(log).Error
}
