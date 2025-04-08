package repo

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-user/pkg/data/organization"
)

type OrganizationRepo interface {
	RegisterOrganization(ctx context.Context, org *organization.Organization, db *gorm.DB) error
}
