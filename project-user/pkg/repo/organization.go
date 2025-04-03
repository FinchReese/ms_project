package repo

import (
	"context"

	"test.com/project-user/pkg/data/organization"
)

type OrganizationRepo interface {
	RegisterOrganization(ctx context.Context, org *organization.Organization) error
}
