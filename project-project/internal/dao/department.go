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

func (d *DepartmentDAO) GetDepartmentList(ctx context.Context, organizationCode int64, pcode int64, page int, pageSize int) ([]*data.Department, int64, error) {
	session := d.conn.Db.Session(&gorm.Session{Context: ctx})
	var departments []*data.Department
	var total int64
	var offset int = (page - 1) * pageSize

	session = session.Model(&data.Department{}).Where("organization_code = ?", organizationCode)
	if pcode > 0 {
		session = session.Where("pcode = ?", pcode)
	}
	err := session.Limit(pageSize).Offset(offset).Find(&departments).Error
	if err != nil {
		return nil, 0, err
	}
	err = session.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return departments, total, nil
}

func (d *DepartmentDAO) AddDepartment(ctx context.Context, department *data.Department) error {
	session := d.conn.Db.Session(&gorm.Session{Context: ctx})
	return session.Create(department).Error
}

func (d *DepartmentDAO) SearchDepartmentList(ctx context.Context, organizationCode int64, pcode int64, name string) ([]*data.Department, error) {
	session := d.conn.Db.Session(&gorm.Session{Context: ctx})
	var departments []*data.Department

	session = session.Model(&data.Department{}).Where("organization_code = ? AND name = ?", organizationCode, name)
	if pcode > 0 {
		session = session.Where("pcode = ?", pcode)
	}
	err := session.Find(&departments).Error
	return departments, err
}

func (d *DepartmentDAO) GetDepartmentById(ctx context.Context, id int64) (*data.Department, error) {
	session := d.conn.Db.Session(&gorm.Session{Context: ctx})
	var department data.Department
	err := session.Model(&data.Department{}).Where("id = ?", id).First(&department).Error
	return &department, err
}
