package repo

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-user/pkg/data/member"
)

type Member interface {
	IsEmailRegistered(ctx context.Context, email string) (bool, error)
	IsAccountRegistered(ctx context.Context, account string) (bool, error)
	IsMobileRegistered(ctx context.Context, mobile string) (bool, error)
	RegisterMember(ctx context.Context, member *member.Member, db *gorm.DB) error
	LoginVerify(ctx context.Context, account string, pwd string) (*member.Member, error)
	FindMemberById(ctx context.Context, id int64) (*member.Member, error)
	FindMembersByIds(ctx context.Context, memberIds []int64) ([]*member.Member, error)
}
