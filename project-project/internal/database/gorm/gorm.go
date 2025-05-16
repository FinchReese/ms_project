package gorm

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"test.com/project-project/config"
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
	// 配置MySQL主仓连接参数
	masterUsername := config.AppConf.MysqlConf.Master.UserName //账号
	masterPassword := config.AppConf.MysqlConf.Master.Password //密码
	masterHost := config.AppConf.MysqlConf.Master.Host         //数据库地址，可以是Ip或者域名
	masterPort := config.AppConf.MysqlConf.Master.Port         //数据库端口
	masterDbName := config.AppConf.MysqlConf.Master.Db         //数据库名
	masterDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", masterUsername, masterPassword, masterHost, masterPort, masterDbName)
	// 创建主仓连接
	var err error
	db, err = gorm.Open(mysql.Open(masterDsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	var slaves []gorm.Dialector
	// 创建从仓连接
	for _, slave := range config.AppConf.MysqlConf.SlaveList {
		slaveUsername := slave.UserName
		slavePassword := slave.Password
		slaveHost := slave.Host
		slavePort := slave.Port
		slaveDbName := slave.Db
		slaveDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", slaveUsername, slavePassword, slaveHost, slavePort, slaveDbName)
		slaves = append(slaves, mysql.Open(slaveDsn))
	}
	// 配置读写分离插件
	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(masterDsn)},
		Replicas: slaves,
		Policy:   dbresolver.RandomPolicy{},
	}).SetConnMaxLifetime(time.Hour * 24))
	if err != nil {
		panic("failed to register dbresolver plugin")
	}
}
