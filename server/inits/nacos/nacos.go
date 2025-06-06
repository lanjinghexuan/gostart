package nacos

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"server/global"
	"server/inits/config"
)

func NacosConfig() {
	appConfig := config.AppConf
	clientConfig := constant.ClientConfig{
		NamespaceId:         appConfig.NamespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      appConfig.IpAddr,
			ContextPath: "/nacos",
			Port:        uint64(appConfig.Port),
			Scheme:      "http",
		},
	}
	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		return
	}

	content, err := client.GetConfig(vo.ConfigParam{
		DataId: appConfig.DataId,
		Group:  appConfig.Group})
	var conf config.Config
	json.Unmarshal([]byte(content), &conf)
	global.Client = conf
}
