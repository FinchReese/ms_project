package dao

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

type ProjectAuthNodeDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewProjectAuthNodeDAO() *ProjectAuthNodeDAO {
	return &ProjectAuthNodeDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (d *ProjectAuthNodeDAO) GetProjectAuthNodeList(ctx context.Context, authId int64) ([]string, error) {
	session := d.conn.Db.Session(&gorm.Session{Context: ctx})
	var nodes []string
	err := session.Model(&data.ProjectAuthNode{}).Where("auth = ?", authId).Pluck("node", &nodes).Error
	return nodes, err
}
