package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"

	"github.com/verystar/nacos-go-sdk"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	server      = kingpin.Flag("server", "Nacos server addr").Short('s').Envar("NACOS_SERVER").String()
	namespaceId = kingpin.Flag("namespace-id", "Nacos namespace id").Short('n').Envar("NACOS_NAMESPACE_ID").String()
	username    = kingpin.Flag("username", "nacos username").Short('u').Envar("NACOS_USERNAME").String()
	password    = kingpin.Flag("password", "Nacos password").Short('p').Envar("NACOS_PASSWORD").String()
	dataId      = kingpin.Flag("data-id", "Nacos data id").Short('d').Envar("NACOS_DATA_ID").String()
	group       = kingpin.Flag("group", "Nacos group name").Short('g').Default("DEFAULT_GROUP").Envar("NACOS_GROUP").String()
	output      = kingpin.Flag("output", "Output to config file or stdout").Short('o').Default("stdout").Envar("NACOS_OUTPUT").String()
	// showVersion = kingpin.Flag("version", "show version").Short('v').Bool()

	// BuildInfo
	version   = "dev"
	revision  = "none"
	builtDate = "unknown"
	builtUser = "unknown"
	goVersion = runtime.Version()
)

func main() {
	kingpin.Version(fmt.Sprintf(`
version: %s
revision: %s
buildUser: %s
buildDate: %s
goVersion: %s
platform: %s/%s`, version, revision, builtUser, builtDate, goVersion, runtime.GOOS, runtime.GOARCH))
	kingpin.Parse()

	// if *showVersion {
	// 	fmt.Printf("version: %s", version)
	// 	os.Exit(0)
	// }

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
		err = ioutil.WriteFile(*output, []byte(data), 0644)
		if err != nil {
			log.Fatalf("write config to file failed: %s\n", err.Error())
		}
		fmt.Println("write config to file succeed: ", *output)
	}

}
