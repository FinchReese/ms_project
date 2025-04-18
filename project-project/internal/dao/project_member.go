package dao

import (
	"context"
	"fmt"

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

func (p *ProjectMemberDAO) GetProjectList(ctx context.Context, memberId int64, selectBy string, page int64, size int64) ([]*data.ProjectAndProjectMember, int64, error) {
	var condition string
	switch selectBy {
	case "", "my":
		condition = "p.deleted=0"
	case "archive":
		condition = "p.archive=1"
	case "deleted":
		condition = "p.deleted=1"
	case "collect":
		condition = "p.id in (select project_code from ms_project_collection where member_code=?)"
	default:
		return nil, 0, fmt.Errorf("invalid selectBy: %s", selectBy)
	}

	start := size * (page - 1)
	var projectList []*data.ProjectAndProjectMember
	var total int64
	session := p.conn.Db.Session(&gorm.Session{Context: ctx})

	if selectBy == "collect" { // 参数列表不同单独处理
		// 获取项目列表
		sql := fmt.Sprintf("select * from ms_project AS p, ms_project_member AS pm where p.id = pm.project_code and pm.member_code = ? and %s order by `order` limit ?,?",
			condition)
		err := session.Raw(sql, memberId, memberId, start, size).Scan(&projectList).Error
		if err != nil {
			zap.L().Error("Get project list error.", zap.Error(err))
			return nil, 0, err
		}
		// 获取查询结果数量
		sql = fmt.Sprintf("select COUNT(*) from ms_project AS p, ms_project_member AS pm where p.id = pm.project_code and pm.member_code = ? and %s order by `order` limit ?,?",
			condition)
		err = session.Raw(sql, memberId, memberId, start, size).Scan(&total).Error
		if err != nil {
			zap.L().Error("Get project list error.", zap.Error(err))
			return nil, 0, err
		}
	} else {
		// 获取项目列表
		sql := fmt.Sprintf("select * from ms_project AS p, ms_project_member AS pm where p.id = pm.project_code and pm.member_code = ? and %s order by `order` limit ?,?",
			condition)
		err := session.Raw(sql, memberId, start, size).Scan(&projectList).Error
		if err != nil {
			zap.L().Error("Get project list error.", zap.Error(err))
			return nil, 0, err
		}
		// 获取查询结果数量
		sql = fmt.Sprintf("select COUNT(*) from ms_project AS p, ms_project_member AS pm where p.id = pm.project_code and pm.member_code = ? and %s order by `order` limit ?,?",
			condition)
		err = session.Raw(sql, memberId, start, size).Scan(&total).Error
		if err != nil {
			zap.L().Error("Get project list error.", zap.Error(err))
			return nil, 0, err
		}
	}
	return projectList, total, nil
}

func (p *ProjectMemberDAO) SaveProjectMember(ctx context.Context, pm *data.ProjectMember, db *gorm.DB) error {
	return db.Session(&gorm.Session{Context: ctx}).Create(pm).Error
}

func (p *ProjectMemberDAO) GetProjectAndMember(ctx context.Context, memberId int64, projectId int64) (*data.ProjectAndProjectMember, error) {
	var pm *data.ProjectAndProjectMember
	session := p.conn.Db.Session(&gorm.Session{Context: ctx})
	sql := fmt.Sprintf("select p.*, pm.project_code, pm.member_code, pm.join_time, pm.is_owner, pm.authorize" +
		" from ms_project AS p, ms_project_member AS pm" +
		" where p.id=pm.project_code and pm.project_code=? and pm.member_code=?" +
		" limit 1")
	tx := session.Raw(sql, projectId, memberId)
	err := tx.Scan(&pm).Error
	if err != nil {
		return nil, err
	}
	if pm == nil {
		zap.L().Error("Project member not found.", zap.Int64("memberId", memberId), zap.Int64("projectId", projectId))
		return nil, fmt.Errorf("Project member not found")
	}
	return pm, nil
}

func (p *ProjectMemberDAO) IsCollectedProject(ctx context.Context, memberId int64, projectId int64) (bool, error) {
	session := p.conn.Db.Session(&gorm.Session{Context: ctx})
	var count int64
	err := session.Model(&data.ProjectCollection{}).Where("project_code = ? and member_code = ?", projectId, memberId).
		Count(&count).Error
	return count > 0, err
}
