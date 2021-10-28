# nacos-initial

用于 CICD 中拉取 Nacos 中的配置到本地配置文件

docker image: daixijun1990/nacos-initial

参数:

```shell
$ ./nacos-initial --help
Usage of ./nacos-initial:
  -dataId string
        Nacos dataId
  -group string
        Nacos Group (default "DEFAULT_GROUP")
  -namespace string
        Nacos namespace id
  -o string
        output to config file or stdout (default "stdout")
  -password string
        Nacos Password
  -server string
        Nacos server addr
  -username string
        Nacos Username

```
