package dao

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

type TaskDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewTaskDAO() *TaskDAO {
	return &TaskDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (td *TaskDAO) FindTasksByStageCode(ctx context.Context, stageCode int) (list []*data.Task, err error) {
	session := td.conn.Db.Session(&gorm.Session{Context: ctx})
	err = session.Model(&data.Task{}).
		Where("stage_code = ? and deleted = 0", stageCode).
		Order("sort asc").
		Find(&list).Error
	return
}
