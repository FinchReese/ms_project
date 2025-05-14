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

func (dao *ProjectAuthDAO) GetProjectAuthListByOrganizationCode(ctx context.Context, organizationCode int64, page int, pageSize int) ([]*data.ProjectAuth, int64, error) {
	session := dao.conn.Db.Session(&gorm.Session{Context: ctx})
	var list []*data.ProjectAuth
	var total int64
	err := session.Model(&data.ProjectAuth{}).
		Where("organization_code = ?", organizationCode).
		Order("sort asc").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	err = session.Model(&data.ProjectAuth{}).
		Where("organization_code = ?", organizationCode).
		Count(&total).Error
	return list, total, err
}
