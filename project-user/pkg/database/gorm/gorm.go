package gorm

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"test.com/project-user/config"
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

func init() {
	//配置MySQL连接参数
	username := config.AppConf.MysqlConf.UserName //账号
	password := config.AppConf.MysqlConf.Password //密码
	host := config.AppConf.MysqlConf.Host         //数据库地址，可以是Ip或者域名
	port := config.AppConf.MysqlConf.Port         //数据库端口
	dbName := config.AppConf.MysqlConf.Db         //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
}
