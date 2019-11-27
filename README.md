gofab
====
`gofab` 基于`Fabric SDK` 创建的扩展，用于更加便利的从客户端操作`Fabric`。

安装
--

文档
--
开始之前，假设你已经获得了一个 Fabric SDK 的实例，可以通过以下示例获取:
```
// 获得SDK实列
sdk := gofab.SDK("config.yaml")
context := sdk.ChannelContext(channelID, fabsdk.WithUser(user), fabsdk.WithOrg(orgName))

// 获取账本客户端实列
ledgerClient, err := ledger.New(context)
gofab.ErrFinal(err)

// 获取账本信息
cfg, err := ledgerClient.QueryConfig()
gofab.ErrLog(err)

// 通过区块的ID获得一个经过解码后的区块
block,err := gofab.GetBlock(1)
gofab.ErrLog(err)
```

### gofab
通过`channel`和`ledger`获取到的区块信息与交易是经过`protobuf`编码的，`gofab`对原方法进行封装
在获得信息后进行解码，并返回相关的标准结构体
* gofab.SDK()
> 获得一个`fabsdk.FabricSDK`实列
* gofab.GetBlock()
> 通过区块的ID获得一个经过解码后的区块
* gofab.GetBlockByHash()
> 通过区块的 hash 获得一个经过解码后的区块
* gofab.GetBlocksByTxId()
> 通过交易ID获得一个经过解码后的区块
* gofab.GetChannelClient()
> 获得一个`channel.Client`实列
* gofab.GetConfig()
> 获得区块链网络配置
* gofab.GetInfo()
> 获得区块链网络信息
* gofab.GetTransactionByTxID()
> 通过交易ID获得一个经过解码后的交易

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

