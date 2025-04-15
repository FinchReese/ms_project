package dao

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

type ProjectMemberDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewProjectMemberDAO() *ProjectMemberDAO {
	return &ProjectMemberDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (p *ProjectMemberDAO) GetProjectListByMemberId(ctx context.Context, memberId int64, page int64, size int64) ([]*data.ProjectAndProjectMember, int64, error) {
	start := size * (page - 1)
	var projectList []*data.ProjectAndProjectMember
	session := p.conn.Db.Session(&gorm.Session{Context: ctx})
	err := session.Raw("select * from ms_project AS p, ms_project_member AS pm where p.id = pm.member_code and pm.member_code = ? order by `order` limit ?,?",
		memberId, start, size).Scan(&projectList).Error
	if err != nil {
		zap.L().Error("Get project list error.", zap.Error(err))
		return nil, 0, err
	}
	var total int64
	err = session.Model(&data.ProjectMember{}).Where("member_code=?", memberId).Count(&total).Error
	if err != nil {
		zap.L().Error("Get project num error.", zap.Error(err))
		return nil, 0, err
	}
	return projectList, total, nil
}
