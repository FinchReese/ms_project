package dao

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-user/pkg/data/organization"
	database_gorm "test.com/project-user/pkg/database/gorm"
)

type OrganizationDao struct {
	conn *gorm.DB
}

var ODao *OrganizationDao

func init() {
	ODao = &OrganizationDao{
		conn: database_gorm.NewMysqlConn().Db,
	}
}

func (oDao *OrganizationDao) RegisterOrganization(ctx context.Context, org *organization.Organization, db *gorm.DB) error {
	err := db.Session(&gorm.Session{Context: ctx}).Create(org).Error
	return err
}
