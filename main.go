package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var (
	endpoint    string
	namespaceId string
	accessKey   string
	secretKey   string
	dataId      string
	group       string
	output      string
	openKMS     bool
	regionId    string
)

func init() {
	flag.StringVar(&endpoint, "endpoint", "acm.aliyun.com", "Nacos/ACM 服务器地址")
	flag.StringVar(&namespaceId, "ns", "", "Nacos/ACM 命名空间ID")
	flag.StringVar(&accessKey, "ak", "", "阿里云RAM子账号 accessKey")
	flag.StringVar(&secretKey, "sk", "", "阿里云RAM子账号 secretKey")
	flag.StringVar(&dataId, "dataId", "", "Nacos/ACM dataId")
	flag.StringVar(&group, "group", "DEFAULT_GROUP", "配置组")
	flag.StringVar(&output, "o", "stdout", "输出配置到终端或者文件")
	flag.BoolVar(&openKMS, "kms", false, "开启KMS自动解密配置内容")
	flag.StringVar(&regionId, "regionId", "cn-hangzhou", "KMS所在区域")
	flag.Parse()
}

func main() {

	clientConfig := constant.ClientConfig{
		Endpoint:       endpoint + ":8080",
		NamespaceId:    namespaceId,
		AccessKey:      accessKey,
		SecretKey:      secretKey,
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
		OpenKMS:        openKMS,
		RegionId:       regionId,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})

	if err != nil {
		panic(fmt.Sprintf("初始化客户端失败: %s\n", err.Error()))
	}

	// 获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group},
	)

	if err != nil {
		panic(fmt.Sprintf("获取配置失败: %s\n", err.Error()))
	}

	if output == "stdout" {
		fmt.Printf("Get config:\n%s\n", content)
	} else {
		err = ioutil.WriteFile(output, []byte(content), 0644)
		if err != nil {
			panic(fmt.Sprintf("写配置文件失败: %s\n", err.Error()))
		}
		fmt.Println("配置成功写入文件: ", output)
	}
}
