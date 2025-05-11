package dao

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

type DepartmentDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewDepartmentDAO() *DepartmentDAO {
	return &DepartmentDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (d *DepartmentDAO) GetDepartmentInfo(ctx context.Context, departmentCode int64) (*data.Department, error) {
	session := d.conn.Db.Session(&gorm.Session{Context: ctx})
	var department data.Department
	err := session.Model(&data.Department{}).Where("id = ?", departmentCode).First(&department).Error
	return &department, err
}
