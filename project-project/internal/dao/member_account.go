package dao

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
	"test.com/project-project/pkg/model"
)

type MemberAccountDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewMemberAccountDAO() *MemberAccountDAO {
	return &MemberAccountDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (d *MemberAccountDAO) GetMemberAccountList(ctx context.Context, queryType int, organizationCode int64, departmentCode int64, page int, pageSize int) ([]*data.MemberAccount, int64, error) {
	session := d.conn.Db.Session(&gorm.Session{Context: ctx})
	// 现先根据组织id筛选
	session = session.Model(&data.MemberAccount{}).Where("organization_code = ?", organizationCode)
	// 再根据查询类型筛选
	var condition string
	switch queryType {
	case model.QueryMemberAccountTypeEnableAccount:
		condition = fmt.Sprintf("status = %d", model.AccountStatusEnabled)
	case model.QueryMemberAccountTypeNullDepartmentCode:
		condition = "department_code is NULL"
	case model.QueryMemberAccountTypeDisableAccount:
		condition = fmt.Sprintf("status = %d", model.AccountStatusDisabled)
	case model.QueryMemberAccountTypeByDepartmentCode:
		condition = fmt.Sprintf("department_code = %d AND status = %d", departmentCode, model.AccountStatusEnabled)
	}
	session = session.Where(condition)

	// 查询指定页成员账号列表
	offset := (page - 1) * pageSize
	var list []*data.MemberAccount
	err := session.Offset(offset).Limit(pageSize).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	// 查询成员账号列表总数
	var total int64
	err = session.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
