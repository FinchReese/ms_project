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
	session := pld.conn.Db.Session(&gorm.Session{Context: ctx})
	return session.Create(log).Error
}

func (pld *ProjectLogDAO) GetAllProjectLogListByTaskId(ctx context.Context, taskId int64, isComment int) ([]*data.ProjectLog, int64, error) {
	var logs []*data.ProjectLog
	session := pld.conn.Db.Session(&gorm.Session{Context: ctx})
	err := session.Model(&data.ProjectLog{}).
		Where("source_code = ? AND is_comment = ?", taskId, isComment).
		Find(&logs).Error
	if err != nil {
		return nil, 0, err
	}
	return logs, int64(len(logs)), nil
}

func (pld *ProjectLogDAO) GetProjectLogListByTaskIdAndPage(ctx context.Context, taskId int64, isComment int, page int64, pageSize int64) ([]*data.ProjectLog, int64, error) {
	var logs []*data.ProjectLog
	var total int64
	offset := (page - 1) * pageSize
	session := pld.conn.Db.Session(&gorm.Session{Context: ctx})
	err := session.Model(&data.ProjectLog{}).
		Where("source_code = ? AND is_comment = ?", taskId, isComment).
		Offset(int(offset)).
		Limit(int(pageSize)).
		Find(&logs).Error
	if err != nil {
		return nil, 0, err
	}
	err = session.Model(&data.ProjectLog{}).
		Where("source_code = ? AND is_comment = ?", taskId, isComment).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return logs, total, nil
}
