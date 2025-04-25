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
