package dao

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

type TaskStageDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewTaskStageDAO() *TaskStageDAO {
	return &TaskStageDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (tsd *TaskStageDAO) SaveTaskStage(ctx context.Context, ts *data.TaskStage, db *gorm.DB) error {
	return db.Session(&gorm.Session{Context: ctx}).Create(ts).Error
}

func (tsd *TaskStageDAO) FindStagesByProjectId(ctx context.Context, projectCode int64, page int, pageSize int) (list []*data.TaskStage, total int64, err error) {
	session := tsd.conn.Db.Session(&gorm.Session{Context: ctx})
	offset := pageSize * (page - 1)
	err = session.Model(&data.TaskStage{}).
		Where("project_code=?", projectCode).
		Offset(offset).
		Limit(pageSize).
		Find(&list).Error
	if err != nil {
		list = nil
		total = 0
		return
	}
	err = session.Model(&data.TaskStage{}).Where("project_code=?", projectCode).Count(&total).Error
	return
}
