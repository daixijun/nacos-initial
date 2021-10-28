package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/verystar/nacos-go-sdk"
)

var (
	server      string
	namespaceId string
	username    string
	password    string
	dataId      string
	group       string
	output      string
)

func init() {
	flag.StringVar(&server, "server", "", "Nacos server addr")
	flag.StringVar(&namespaceId, "namespace", "", "Nacos namespace id")
	flag.StringVar(&username, "username", "", "Nacos Username")
	flag.StringVar(&password, "password", "", "Nacos Password")
	flag.StringVar(&dataId, "dataId", "", "Nacos dataId")
	flag.StringVar(&group, "group", "DEFAULT_GROUP", "Nacos Group")
	flag.StringVar(&output, "o", "stdout", "output to config file or stdout")
	flag.Parse()
}

func main() {
	conf := nacos.NewNacosConfig(func(c *nacos.NacosConfig) {
		c.Username = username
		c.Password = password
		c.ServerAddr = server
	})

	data, err := conf.Get(namespaceId, group, dataId)

	if err != nil {
		log.Fatalf("get config failed: %s\n", err)
	}

	if output == "stdout" {
		fmt.Printf("get config: %s\n", data)
	} else {
		err = ioutil.WriteFile(output, []byte(data), 0644)
		if err != nil {
			log.Fatalf("write config to file failed: %s\n", err.Error())
		}
		fmt.Println("write config to file succeed: ", output)
	}

}
