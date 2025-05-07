package dao

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
	"test.com/project-project/internal/repo"
)

type SourceLinkDao struct {
	conn *custom_gorm.MysqlConn
}

// 创建SourceLinkRepo接口实例
func NewSourceLinkDao() repo.SourceLinkRepo {
	return &SourceLinkDao{
		conn: custom_gorm.NewMysqlConn(),
	}
}

// SaveSourceLink 保存资源关联记录
func (s *SourceLinkDao) SaveSourceLink(ctx context.Context, sourceLink *data.SourceLink, db *gorm.DB) error {
	if db == nil {
		return errors.New("db is nil")
	}
	return db.Create(sourceLink).Error
}
