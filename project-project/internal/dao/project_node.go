package dao

import (
	"context"

	"gorm.io/gorm"

	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

type ProjectNodeDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewProjectNodeDAO() *ProjectNodeDAO {
	return &ProjectNodeDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (d *ProjectNodeDAO) GetAllProjectNodeList(ctx context.Context) ([]*data.ProjectNode, error) {
	session := d.conn.Db.Session(&gorm.Session{Context: ctx})
	var nodes []*data.ProjectNode
	err := session.Model(&data.ProjectNode{}).Find(&nodes).Error
	return nodes, err
}
