package dao

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

type TemplateTaskStageDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewTemplateTaskStageDAO() *TemplateTaskStageDAO {
	return &TemplateTaskStageDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (t *TemplateTaskStageDAO) GetTaskStagesByTemplateIds(ctx context.Context, ids []int) ([]data.TemplateTaskStage, error) {
	var taskStages []data.TemplateTaskStage
	session := t.conn.Db.Session(&gorm.Session{Context: ctx})
	err := session.Model(&data.TemplateTaskStage{}).Where("project_template_code in ?", ids).Find(&taskStages).Error
	if err != nil {
		zap.L().Error("Query template task stage err.", zap.Error(err))
		return nil, err
	}
	return taskStages, nil
}
