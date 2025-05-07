package dao

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
	"test.com/project-project/internal/repo"
)

type FileDao struct {
	conn *custom_gorm.MysqlConn
}

// 创建FileRepo接口实例
func NewFileDao() repo.FileRepo {
	return &FileDao{
		conn: custom_gorm.NewMysqlConn(),
	}
}

// SaveFile 保存文件记录
func (f *FileDao) SaveFile(ctx context.Context, file *data.File, db *gorm.DB) error {
	if db == nil {
		return errors.New("db is nil")
	}
	return db.Create(file).Error
}
