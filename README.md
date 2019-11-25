gofab
====
`gofab` 基于`Fabric SDK` 创建的扩展，用于更加便利的从客户端操作`Fabric`。

安装
--

文档
--
开始之前，假设你已经获得了一个 Fabric SDK 的实例，可以通过以下示例获取:
```
import (
"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

sdk, err := fabsdk.New(config.FromFile("config.yaml"))
if err != nil {
    log.Panicf("Failed to create new SDK: %s", err)
}
defer sdk.Close()
```

### channel
1. 获取 `channel` 管道的客户端
```
gofab.GetChannelClient(<管道ID>,<操作用户名>,<操作节点的组织名称>) *channel.Client
```