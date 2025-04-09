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

func (oDao *OrganizationDao) GetOrganizationByMemberId(ctx context.Context, memberId int64) ([]*organization.Organization, error) {
	var orgs []*organization.Organization
	err := oDao.conn.Session(&gorm.Session{Context: ctx}).Model(&organization.Organization{}).Where("member_id=?", memberId).Find(&orgs).Error
	return orgs, err
}
