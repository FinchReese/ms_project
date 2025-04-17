package trans

import (
	database_gorm "test.com/project-project/internal/database/gorm"
)

type DbConn interface {
	Begin()
	Rollback()
	Commit()
}

type TransactionImpl struct {
	conn DbConn
}

func (t *TransactionImpl) ExecTran(f func(conn DbConn) error) error {
	t.conn.Begin()
	err := f(t.conn)
	if err != nil {
		t.conn.Rollback()
		return err
	}
	t.conn.Commit()
	return nil
}

func NewTransaction() *TransactionImpl {
	return &TransactionImpl{
		conn: database_gorm.NewMysqlConn(),
	}
}
