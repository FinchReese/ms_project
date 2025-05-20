package gorm

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var db *gorm.DB

type MysqlConn struct {
	Db     *gorm.DB
	TranDb *gorm.DB
}

func NewMysqlConn() *MysqlConn {
	return &MysqlConn{
		Db: db,
	}
}

func (mc *MysqlConn) Begin() {
	zap.L().Info("call Begin")
	mc.TranDb = mc.Db.Begin()
}

func (mc *MysqlConn) Commit() {
	zap.L().Info("call Commit")
	mc.TranDb.Commit()
}

func (mc *MysqlConn) Rollback() {
	mc.TranDb.Rollback()
	zap.L().Info("call Rollback")
}

func InitMysqlDb(_db *gorm.DB) {
	db = _db
}
