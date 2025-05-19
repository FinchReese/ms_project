package gorm

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"test.com/project-project/config"
)

const (
	MaxReconnectTimes = 10 // 最大重连次数
	ReconnectInterval = 5  // 重连间隔为5秒
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
	for i := 0; i < MaxReconnectTimes; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			zap.L().Error("连接数据库失败, 5秒后尝试重连" + err.Error())
			time.Sleep(ReconnectInterval * time.Second)
			continue
		}
		// 验证连接是否真正可用
		sqlDB, err := db.DB()
		if err != nil {
			zap.L().Error("获取 SQL DB 失败", zap.Error(err))
			time.Sleep(ReconnectInterval * time.Second)
			continue
		}
		if err = sqlDB.Ping(); err == nil {
			zap.L().Info("数据库连接成功")
			break
		}
		zap.L().Error("Ping 数据库失败", zap.Error(err))
		time.Sleep(ReconnectInterval * time.Second)
	}
}
