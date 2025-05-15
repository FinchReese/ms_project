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

func (d *ProjectAuthNodeDAO) DeleteProjectAuthNode(ctx context.Context, authId int64, db *gorm.DB) error {
	return db.Session(&gorm.Session{Context: ctx}).Where("auth = ?", authId).Delete(&data.ProjectAuthNode{}).Error
}

func (d *ProjectAuthNodeDAO) AddProjectAuthNode(ctx context.Context, authId int64, nodeList []string, db *gorm.DB) error {
	// 先生成auth_node表数据
	authNodes := make([]*data.ProjectAuthNode, len(nodeList))
	for i, node := range nodeList {
		authNodes[i] = &data.ProjectAuthNode{
			Auth: authId,
			Node: node,
		}
	}
	// 批量插入auth_node表数据
	return db.Session(&gorm.Session{Context: ctx}).Create(authNodes).Error
}
