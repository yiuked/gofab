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
获取 `channel` 管道的客户端
```
gofab.GetChannelClient(<管道ID>,<操作用户名>,<操作节点的组织名称>) *channel.Client
```
channel.Client支持以下操作：
* channel.Client.Query()
> 执行查询
* channel.Client.Execute() 
> 执行交易类型操作
* channel.Client.InvokeHandler()
> 执行交易类型操作
* channel.Client.RegisterChaincodeEvent()
> 注册链码事件
* channel.Client.UnregisterChaincodeEvent()
> 注销链码事件

### ledger
获取`ledger`客户端
```
ledger.New(context)
```
ledger.Client 支持以下操作：
* ledger.Client.QueryInfo()
> 获取当前区块链网络信息
* ledger.Client.QueryBlock()
> 获取区块信息
* ledger.Client.QueryConfig()
> 获取配置信息
* ledger.Client.QueryBlockByHash()
> 通过Hash获取交易信息
* ledger.Client.QueryBlockByTxID()
> 通过交易ID获取Block信息
* ledger.Client.QueryTransaction()
> 通过交易ID获取交易信息