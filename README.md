# nacos-initial

用于 CICD 中拉取 Nacos 中的配置到本地配置文件

docker image: daixijun1990/nacos-initial

参数:

```shell
$ ./nacos-initial --help
usage: nacos-initial [<flags>]

Flags:
      --help                   Show context-sensitive help (also try --help-long and --help-man).
  -s, --server=SERVER          Nacos server addr
  -n, --namespace-id=NAMESPACE-ID
                               Nacos namespace id
  -u, --username=USERNAME      nacos username
  -p, --password=PASSWORD      Nacos password
  -d, --data-id=DATA-ID        Nacos data id
  -g, --group="DEFAULT_GROUP"  Nacos group name
  -o, --output="stdout"        Output to config file or stdout
      --version                Show application version.

```
