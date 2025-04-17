package dao

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

const (
	isNotSystemTemplate = 0
	isSystemTemplate    = 1
)

type ProjectTemplateDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewProjectTemplateDAO() *ProjectTemplateDAO {
	return &ProjectTemplateDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (p *ProjectTemplateDAO) GetSystemProjectTemplates(ctx context.Context, page int64, size int64) ([]data.ProjectTemplate, int64, error) {
	var templates []data.ProjectTemplate
	session := p.conn.Db.Session(&gorm.Session{Context: ctx})
	err := session.Model(&data.ProjectTemplate{}).
		Where("is_system=?", isSystemTemplate).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&templates).Error
	if err != nil {
		zap.L().Error("Query system template err.", zap.Error(err))
		return nil, 0, err
	}
	var total int64
	err = session.Model(&data.ProjectTemplate{}).Where("is_system=?", isSystemTemplate).Count(&total).Error
	if err != nil {
		zap.L().Error("Get res num err.", zap.Error(err))
		return nil, 0, err
	}
	return templates, total, nil
}
func (p *ProjectTemplateDAO) GetCustomProjectTemplates(ctx context.Context, memId int64, page int64, size int64) ([]data.ProjectTemplate, int64, error) {
	var templates []data.ProjectTemplate
	session := p.conn.Db.Session(&gorm.Session{Context: ctx})
	err := session.Model(&data.ProjectTemplate{}).
		Where("member_code=?", memId).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&templates).Error
	if err != nil {
		zap.L().Error("Query system template err.", zap.Error(err))
		return nil, 0, err
	}
	var total int64
	err = session.Model(&data.ProjectTemplate{}).Where("member_code=?", memId).Count(&total).Error
	if err != nil {
		zap.L().Error("Get res num err.", zap.Error(err))
		return nil, 0, err
	}
	return templates, total, nil
}

func (p *ProjectTemplateDAO) GetAllProjectTemplates(ctx context.Context, page int64, size int64) ([]data.ProjectTemplate, int64, error) {
	var templates []data.ProjectTemplate
	session := p.conn.Db.Session(&gorm.Session{Context: ctx})
	err := session.Model(&data.ProjectTemplate{}).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&templates).Error
	if err != nil {
		zap.L().Error("Query system template err.", zap.Error(err))
		return nil, 0, err
	}
	var total int64
	err = session.Model(&data.ProjectTemplate{}).Count(&total).Error
	if err != nil {
		zap.L().Error("Get res num err.", zap.Error(err))
		return nil, 0, err
	}
	return templates, total, nil
}
