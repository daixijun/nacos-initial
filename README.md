# nacos-initial

用于CICD中拉取ACM/Nacos中的配置到本地

docker image: daixijun1990/nacos-initial

> 目前暂不支持ACM中使用`KMS AES-128`加密过的配置

参数:

```shell
$ ./nacos-initial --help
Usage of ./nacos-initial:
  -ak string
      阿里云RAM子账号 accessKey
  -dataId string
      Nacos/ACM dataId
  -endpoint string
      Nacos/ACM 服务器地址 (default "acm.aliyun.com")
  -group string
      配置组 (default "DEFAULT_GROUP")
  -kms
      开启KMS自动解密配置内容
  -ns string
      Nacos/ACM 命名空间ID
  -o string
      输出配置到终端或者文件 (default "stdout")
  -regionId string
      KMS所在区域 (default "cn-hangzhou")
  -sk string
      阿里云RAM子账号 secretKey
```
