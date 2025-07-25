package config

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"test.com/project-common/logs"
)

const (
	AppConfigDataId = "app.yaml"
)

type ServerConfig struct {
	Name string
	Addr string
}

type GrpcConfig struct {
	Name        string
	ListenAddr  string
	ConnectAddr string
	Version     string
	Weight      int64
}

type EtcdConfig struct {
	Addrs []string
}

type MysqlConfig struct {
	UserName string
	Password string
	Host     string
	Port     int64
	Db       string
}

type KafkaConfig struct {
	Addr string
}

type JaegerConfig struct {
	CollectorAddr string
}

type Config struct {
	viper      *viper.Viper
	ServerConf *ServerConfig
	GrpcConf   *GrpcConfig
	EtcdConf   *EtcdConfig
	MysqlConf  *MysqlConfig
	KafkaConf  *KafkaConfig
	JaegerConf *JaegerConfig
}

var AppConf = initConfig()

// 从nacos读取配置
func getConfigFromNacos(conf *Config) error {
	// 创建nacos客户端
	nacosClient := InitNacosClient()
	if nacosClient == nil {
		return errors.New("Init nacos client err")
	}
	// 读取nacos配置
	configContent, err := nacosClient.confClient.GetConfig(vo.ConfigParam{
		DataId: AppConfigDataId,
		Group:  nacosClient.group,
	})
	if err != nil {
		return fmt.Errorf("nacos get config err, err msg: %v", err)
	}
	if configContent == "" {
		return fmt.Errorf("nacos not found config")
	}
	log.Printf("Get config from nacos success, config content:\n%s\n", configContent)
	// 将读取到的配置信息传给viper
	err = conf.viper.ReadConfig(bytes.NewBuffer([]byte(configContent)))
	if err != nil {
		return err
	}
	// 监听nacos配置文件变化
	err = nacosClient.confClient.ListenConfig(vo.ConfigParam{
		DataId: AppConfigDataId,
		Group:  nacosClient.group,
		OnChange: func(namespace, group, dataId, data string) {
			log.Printf("load nacos config changed %s \n", data)
			err := conf.viper.ReadConfig(bytes.NewBuffer([]byte(data)))
			if err != nil {
				log.Printf("load nacos config changed err : %s \n", err.Error())
				return
			}
			//所有的配置应该重新读取
			conf.ReLoadAllConfig()
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func getConfigFromLocalFile(conf *Config) error {
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("app")
	conf.viper.AddConfigPath(workDir + "/config")
	err := conf.viper.ReadInConfig()
	return err
}

func initConfig() *Config {
	v := viper.New()
	conf := &Config{viper: v}
	conf.viper.SetConfigType("yaml")
	err := getConfigFromNacos(conf)
	// 如果从nacos读取配置失败，则尝试从本地配置文件读取
	if err != nil {
		log.Println(err)
		err = getConfigFromLocalFile(conf)
		if err != nil {
			log.Printf("Get config from local file err: %v\n", err)
			return nil
		}
	}
	log.Println("Get config from local file success")
	conf.ReLoadAllConfig()
	return conf
}

func (c *Config) ReLoadAllConfig() {
	c.ReadGrpcConfig()
	c.ReadServerConfig()
	c.ReadEtcdConfig()
	c.ReadMysqlConfig()
	c.ReadKafkaConfig()
	c.ReadJaegerConfig()
	InitRedisClient(c.InitRedisOptions())
	InitMysqlClient(c.MysqlConf)
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
	gc.ListenAddr = c.viper.GetString("grpc.listenAddr")
	gc.ConnectAddr = c.viper.GetString("grpc.connectAddr")
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
	mc.UserName = c.viper.GetString("mysql.username")
	mc.Password = c.viper.GetString("mysql.password")
	mc.Host = c.viper.GetString("mysql.host")
	mc.Port = c.viper.GetInt64("mysql.port")
	mc.Db = c.viper.GetString("mysql.db")
	c.MysqlConf = mc
}

func (c *Config) ReadKafkaConfig() {
	kafkaConf := &KafkaConfig{}
	err := c.viper.UnmarshalKey("kafka", &kafkaConf)
	if err != nil {
		log.Fatalln(err)
	}
	c.KafkaConf = kafkaConf
}

func (c *Config) ReadJaegerConfig() {
	jaegerConf := &JaegerConfig{}
	err := c.viper.UnmarshalKey("jaeger", &jaegerConf)
	if err != nil {
		log.Fatalln(err)
	}
	c.JaegerConf = jaegerConf
}
