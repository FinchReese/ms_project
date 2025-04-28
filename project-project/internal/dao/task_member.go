package dao

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

type TaskMemberDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewTaskMemberDAO() *TaskMemberDAO {
	return &TaskMemberDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (tmd *TaskMemberDAO) FindTaskMembers(ctx context.Context, taskCode int64, memberCode int64) (list []*data.TaskMember, err error) {
	session := tmd.conn.Db.Session(&gorm.Session{Context: ctx})
	err = session.Model(&data.TaskMember{}).Where("task_code = ? and member_code = ?", taskCode, memberCode).Find(&list).Error
	return
}

// SaveTaskMember 保存任务成员关系
func (tmd *TaskMemberDAO) SaveTaskMember(ctx context.Context, taskMember *data.TaskMember, db *gorm.DB) error {
	if db == nil {
		db = tmd.conn.Db.Session(&gorm.Session{Context: ctx})
	}
	return db.Save(taskMember).Error
}
