package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/alecthomas/kingpin/v2"
	"github.com/verystar/nacos-go-sdk"
)

var (
	server      = kingpin.Flag("server", "Nacos server addr").Short('s').Envar("NACOS_SERVER").Required().String()
	namespaceId = kingpin.Flag("namespace-id", "Nacos namespace id").Short('n').Envar("NACOS_NAMESPACE_ID").Required().String()
	username    = kingpin.Flag("username", "nacos username").Short('u').Envar("NACOS_USERNAME").Required().String()
	password    = kingpin.Flag("password", "Nacos password").Short('p').Envar("NACOS_PASSWORD").Required().String()
	dataId      = kingpin.Flag("data-id", "Nacos data id").Short('d').Envar("NACOS_DATA_ID").Required().String()
	group       = kingpin.Flag("group", "Nacos group name").Short('g').Default("DEFAULT_GROUP").Envar("NACOS_GROUP").String()
	output      = kingpin.Flag("output", "Output to config file or stdout").Short('o').Default("stdout").Envar("NACOS_OUTPUT").String()

	// BuildInfo
	version   = "dev"
	revision  = "none"
	builtDate = "unknown"
	goVersion = runtime.Version()
)

func main() {
	kingpin.Version(fmt.Sprintf(`
version: %s
goVersion: %s
revision: %s
buildDate: %s
platform: %s/%s`, version, revision, builtDate, goVersion, runtime.GOOS, runtime.GOARCH))
	kingpin.Parse()

	conf := nacos.NewNacosConfig(func(c *nacos.NacosConfig) {
		c.Username = *username
		c.Password = *password
		c.ServerAddr = *server
	})

	data, err := conf.Get(*namespaceId, *group, *dataId)

	if err != nil {
		log.Fatalf("get config failed: %s\n", err)
	}

	if *output == "stdout" {
		fmt.Printf("get config: %s\n", data)
	} else {
		err = os.WriteFile(*output, []byte(data), 0644)
		if err != nil {
			log.Fatalf("write config to file failed: %s\n", err.Error())
		}
		fmt.Println("write config to file succeed: ", *output)
	}

}
