package config

import (
	"log"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type NacosClient struct {
	confClient config_client.IConfigClient
	group      string
}

func InitNacosClient() *NacosClient {
	bootConf := InitBootConfig()
	clientConfig := constant.ClientConfig{
		NamespaceId:         bootConf.NacosConfig.NamespaceId,
		NotLoadCacheAtStart: true,
		LogDir:              "E:\\闭关\\项目\\gin管理系统\\nacos\\log",
		CacheDir:            "E:\\闭关\\项目\\gin管理系统\\nacos\\cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      bootConf.NacosConfig.IpAddr,
			ContextPath: bootConf.NacosConfig.ContextPath,
			Port:        uint64(bootConf.NacosConfig.Port),
			Scheme:      bootConf.NacosConfig.Scheme,
		},
	}
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
	nc := &NacosClient{
		confClient: configClient,
		group:      bootConf.NacosConfig.Group,
	}
	return nc
}
