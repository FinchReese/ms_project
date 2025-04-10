package dao

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

type MenuDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewMenuDAO() *MenuDAO {
	return &MenuDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (md *MenuDAO) GetAllMenus(ctx context.Context) (menuList []*data.ProjectMenu, err error) {
	err = md.conn.Db.Session(&gorm.Session{Context: ctx}).Model(&data.ProjectMenu{}).Find(&menuList).Error
	return
}
