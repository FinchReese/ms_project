package dao

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

type TaskWorkTimeDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewTaskWorkTimeDAO() *TaskWorkTimeDAO {
	return &TaskWorkTimeDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (td *TaskWorkTimeDAO) SaveTaskWorkTime(ctx context.Context, taskWorkTime *data.TaskWorkTime) error {
	session := td.conn.Db.Session(&gorm.Session{Context: ctx})
	return session.Create(taskWorkTime).Error
}

func (td *TaskWorkTimeDAO) GetTaskWorkTimeList(ctx context.Context, taskCode int64) (list []*data.TaskWorkTime, err error) {
	session := td.conn.Db.Session(&gorm.Session{Context: ctx})
	err = session.Where("task_code = ?", taskCode).Find(&list).Error
	return
}
