package gorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"test.com/project-user/config"
)

var MysqlConn *gorm.DB

func init() {
	//配置MySQL连接参数
	username := config.AppConf.MysqlConf.UserName //账号
	password := config.AppConf.MysqlConf.Password //密码
	host := config.AppConf.MysqlConf.Host         //数据库地址，可以是Ip或者域名
	port := config.AppConf.MysqlConf.Port         //数据库端口
	db := config.AppConf.MysqlConf.Db             //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, db)
	var err error
	MysqlConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
}
