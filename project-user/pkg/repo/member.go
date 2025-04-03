package repo

import (
	"context"

	"test.com/project-user/pkg/data/member"
)

type Member interface {
	IsEmailRegistered(ctx context.Context, email string) (bool, error)
	IsAccountRegistered(ctx context.Context, account string) (bool, error)
	IsMobileRegistered(ctx context.Context, mobile string) (bool, error)
	RegisterMember(ctx context.Context, member *member.Member) error
}
