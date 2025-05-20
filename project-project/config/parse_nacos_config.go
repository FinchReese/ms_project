package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type NacosConfig struct {
	NamespaceId string
	Group       string
	IpAddr      string
	Port        int
	ContextPath string
	Scheme      string
}

type BootConfig struct {
	viper       *viper.Viper
	NacosConfig *NacosConfig
}

func (bc *BootConfig) ReadBootConfig() {
	if bc.viper == nil {
		log.Fatalln("viper is nil")
		return
	}
	nc := &NacosConfig{}
	bc.viper.UnmarshalKey("nacos", nc)
	bc.NacosConfig = nc
}

func InitBootConfig() *BootConfig {
	v := viper.New()
	conf := &BootConfig{viper: v}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("nacos_config")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(workDir + "/config")

	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	conf.ReadBootConfig()
	return conf
}
