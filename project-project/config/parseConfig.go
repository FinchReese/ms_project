package config

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"test.com/project-common/logs"
)

type ServerConfig struct {
	Name string
	Addr string
}

type GrpcConfig struct {
	Name    string
	Addr    string
	Version string
	Weight  int64
}

type EtcdConfig struct {
	Addrs []string
}

type MysqlConfig struct {
	Master    *DbConfig   `mapstructure:"master"`
	SlaveList []*DbConfig `mapstructure:"slaveList"`
}

type DbConfig struct {
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int64  `mapstructure:"port"`
	Db       string `mapstructure:"db"`
}

type Config struct {
	viper      *viper.Viper
	ServerConf *ServerConfig
	GrpcConf   *GrpcConfig
	EtcdConf   *EtcdConfig
	MysqlConf  *MysqlConfig
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
	conf.ReadGrpcConfig()
	conf.ReadServerConfig()
	conf.ReadEtcdConfig()
	conf.ReadMysqlConfig()
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

func (c *Config) InitRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     c.viper.GetString("redis.host") + ":" + c.viper.GetString("redis.port"),
		Password: c.viper.GetString("redis.password"), // no password set
		DB:       c.viper.GetInt("db"),                // use default DB
	}
}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.ServerConf = sc
}

func (c *Config) ReadGrpcConfig() {
	gc := &GrpcConfig{}
	gc.Name = c.viper.GetString("grpc.name")
	gc.Addr = c.viper.GetString("grpc.addr")
	gc.Version = c.viper.GetString("grpc.version")
	gc.Weight = c.viper.GetInt64("grpc.weight")
	c.GrpcConf = gc
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

func (c *Config) ReadMysqlConfig() {
	mc := &MysqlConfig{}
	err := c.viper.UnmarshalKey("mysql", &mc)
	if err != nil {
		log.Fatalln(err)
	}
	c.MysqlConf = mc
}
