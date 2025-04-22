package dao

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

type ProjectCollectDao struct {
	conn *custom_gorm.MysqlConn
}

func NewProjectCollectDao() *ProjectCollectDao {
	return &ProjectCollectDao{
		conn: custom_gorm.NewMysqlConn(),
	}
}

// 收藏项目
func (p *ProjectCollectDao) Collect(ctx context.Context, memberId int64, projectId int64, createTime int64) error {
	collection := &data.ProjectCollection{
		MemberCode:  memberId,
		ProjectCode: projectId,
		CreateTime:  createTime,
	}
	session := p.conn.Db.Session(&gorm.Session{Context: ctx})
	err := session.Create(collection).Error
	if err != nil {
		zap.L().Error("Collect project error", zap.Error(err))
		return err
	}
	return nil
}

// 取消收藏项目
func (p *ProjectCollectDao) CancelCollect(ctx context.Context, memberId int64, projectId int64) error {
	session := p.conn.Db.Session(&gorm.Session{Context: ctx})
	err := session.Where("member_code = ? and project_code = ?", memberId, projectId).
		Delete(&data.ProjectCollection{}).Error
	if err != nil {
		zap.L().Error("Cancel collect project error", zap.Error(err))
		return err
	}
	return nil
}
