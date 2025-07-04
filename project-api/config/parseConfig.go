package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"test.com/project-common/logs"
)

type ServerConfig struct {
	Name string
	Addr string
}

type EtcdConfig struct {
	Addrs []string
}

type MinIOConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
	Bucket    string
}

type JaegerConfig struct {
	CollectorAddr string
}

type Config struct {
	viper      *viper.Viper
	ServerConf *ServerConfig
	EtcdConf   *EtcdConfig
	MinIOConf  *MinIOConfig
	JaegerConf *JaegerConfig
}

var AppConf = initConfig()

func initConfig() *Config {
	v := viper.New()
	conf := &Config{viper: v}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("app")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(workDir + "/config")

	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	conf.ReadServerConfig()
	conf.ReadEtcdConfig()
	conf.ReadMinIOConfig()
	conf.ReadJaegerConfig()
	return conf
}

func (c *Config) InitZapLog() {
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("maxSize"),
		MaxAge:        c.viper.GetInt("maxAge"),
		MaxBackups:    c.viper.GetInt("maxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.ServerConf = sc
}

func (c *Config) ReadEtcdConfig() {
	ec := &EtcdConfig{}
	var addrs []string
	err := c.viper.UnmarshalKey("etcd.addrs", &addrs)
	if err != nil {
		log.Fatalln(err)
	}
	ec.Addrs = addrs
	c.EtcdConf = ec
}

func (c *Config) ReadMinIOConfig() {
	minIOConf := &MinIOConfig{}
	err := c.viper.UnmarshalKey("minIO", &minIOConf)
	if err != nil {
		log.Fatalln(err)
	}
	c.MinIOConf = minIOConf
}

func (c *Config) ReadJaegerConfig() {
	jaegerConf := &JaegerConfig{}
	err := c.viper.UnmarshalKey("jaeger", &jaegerConf)
	if err != nil {
		log.Fatalln(err)
	}
	c.JaegerConf = jaegerConf
}
