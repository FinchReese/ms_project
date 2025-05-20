package config

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

const (
	MaxReconnectTimes = 10 // 最大重连次数
	ReconnectInterval = 5  // 重连间隔为5秒
)

func InitMysqlClient(cfg *MysqlConfig) {
	//配置MySQL连接参数
	username := cfg.UserName //账号
	password := cfg.Password //密码
	host := cfg.Host         //数据库地址，可以是Ip或者域名
	port := cfg.Port         //数据库端口
	dbName := cfg.Db         //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)
	for i := 0; i < MaxReconnectTimes; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
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
			custom_gorm.InitMysqlDb(db)
			return
		}
		zap.L().Error("Ping 数据库失败", zap.Error(err))
		time.Sleep(ReconnectInterval * time.Second)
	}
	zap.L().Error("连接数据库失败")
}
