package dao

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

type ProjectAuthDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewProjectAuthDAO() *ProjectAuthDAO {
	return &ProjectAuthDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (dao *ProjectAuthDAO) GetProjectAuthList(ctx context.Context, organizationCode int64) ([]*data.ProjectAuth, error) {
	session := dao.conn.Db.Session(&gorm.Session{Context: ctx})
	var list []*data.ProjectAuth
	err := session.Model(&data.ProjectAuth{}).
		Where("organization_code = ?", organizationCode).
		Order("sort asc").
		Find(&list).Error
	return list, err
}
