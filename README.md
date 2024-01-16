# nacos-initial

用于 CICD 中拉取 Nacos 中的配置到本地配置文件

docker image: daixijun1990/nacos-initial

参数:

```shell
$ ./nacos-initial --help
usage: nacos-initial --server=SERVER --namespace-id=NAMESPACE-ID --username=USERNAME --password=PASSWORD --data-id=DATA-ID [<flags>]


Flags:
      --[no-]help              Show context-sensitive help (also try --help-long and --help-man).
  -s, --server=SERVER          Nacos server addr ($NACOS_SERVER)
  -n, --namespace-id=NAMESPACE-ID
                               Nacos namespace id ($NACOS_NAMESPACE_ID)
  -u, --username=USERNAME      nacos username ($NACOS_USERNAME)
  -p, --password=PASSWORD      Nacos password ($NACOS_PASSWORD)
  -d, --data-id=DATA-ID        Nacos data id ($NACOS_DATA_ID)
  -g, --group="DEFAULT_GROUP"  Nacos group name ($NACOS_GROUP)
  -o, --output="stdout"        Output to config file or stdout ($NACOS_OUTPUT)
      --[no-]version           Show application version.
```
