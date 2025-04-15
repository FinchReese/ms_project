package dao

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"test.com/project-user/pkg/data/member"
	database_gorm "test.com/project-user/pkg/database/gorm"
)

type MemberDao struct {
	conn *gorm.DB
}

var MDao *MemberDao

func init() {
	MDao = &MemberDao{
		conn: database_gorm.NewMysqlConn().Db,
	}
}

func (md *MemberDao) IsEmailRegistered(ctx context.Context, email string) (bool, error) {
	var count int64
	err := md.conn.Session(&gorm.Session{Context: ctx}).Model(&member.Member{}).Where("email=?", email).Count(&count).Error
	return count > 0, err
}

func (md *MemberDao) IsAccountRegistered(ctx context.Context, account string) (bool, error) {
	var count int64
	err := md.conn.Session(&gorm.Session{Context: ctx}).Model(&member.Member{}).Where("account=?", account).Count(&count).Error
	return count > 0, err
}

func (md *MemberDao) IsMobileRegistered(ctx context.Context, mobile string) (bool, error) {
	var count int64
	err := md.conn.Session(&gorm.Session{Context: ctx}).Model(&member.Member{}).Where("mobile=?", mobile).Count(&count).Error
	return count > 0, err
}

func (md *MemberDao) RegisterMember(ctx context.Context, member *member.Member, db *gorm.DB) error {
	return db.Session(&gorm.Session{Context: ctx}).Create(member).Error
}

func (md *MemberDao) LoginVerify(ctx context.Context, account string, pwd string) (*member.Member, error) {
	var mem *member.Member
	err := md.conn.Session(&gorm.Session{Context: ctx}).Model(&member.Member{}).Where("account=? and password=?", account, pwd).First(&mem).Error
	if err != nil {
		zap.L().Error("find member by id error", zap.String("account", account), zap.Error(err))
		return nil, err
	}
	return mem, nil
}

func (md *MemberDao) FindMemberById(ctx context.Context, id int64) (*member.Member, error) {
	var mem *member.Member
	err := md.conn.Session(&gorm.Session{Context: ctx}).Model(&member.Member{}).Where("id=?", id).First(&mem).Error
	if err != nil {
		zap.L().Error("find member by id error", zap.Int64("id", id), zap.Error(err))
		return nil, err
	}
	return mem, nil
}
