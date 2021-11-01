package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/verystar/nacos-go-sdk"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	server      = kingpin.Flag("server", "Nacos server addr").Short('s').Required().Envar("NACOS_SERVER").String()
	namespaceId = kingpin.Flag("namespace-id", "Nacos namespace id").Short('n').Required().Envar("NACOS_NAMESPACE_ID").String()
	username    = kingpin.Flag("username", "nacos username").Short('u').Required().Envar("NACOS_USERNAME").String()
	password    = kingpin.Flag("password", "Nacos password").Short('p').Required().Envar("NACOS_PASSWORD").String()
	dataId      = kingpin.Flag("data-id", "Nacos data id").Short('d').Required().Envar("NACOS_DATA_ID").String()
	group       = kingpin.Flag("group", "Nacos group name").Short('g').Default("DEFAULT_GROUP").Envar("NACOS_GROUP").String()
	output      = kingpin.Flag("output", "Output to config file or stdout").Short('o').Default("stdout").Envar("NACOS_OUTPUT").String()
)

// func init() {

// 	flag.StringVar(&server, "server", "", "Nacos server addr")
// 	flag.StringVar(&namespaceId, "namespace", "", "Nacos namespace id")
// 	flag.StringVar(&username, "username", "", "Nacos Username")
// 	flag.StringVar(&password, "password", "", "Nacos Password")
// 	flag.StringVar(&dataId, "dataId", "", "Nacos dataId")
// 	flag.StringVar(&group, "group", "DEFAULT_GROUP", "Nacos Group")
// 	flag.StringVar(&output, "o", "stdout", "output to config file or stdout")
// }

func main() {
	// flag.Parse()
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
		err = ioutil.WriteFile(*output, []byte(data), 0644)
		if err != nil {
			log.Fatalf("write config to file failed: %s\n", err.Error())
		}
		fmt.Println("write config to file succeed: ", *output)
	}

}
