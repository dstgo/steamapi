[English](./README.md)|简体中文

# steamapi
![Static Badge](https://img.shields.io/badge/go-%3E%3D1.21-blue)
![GitHub](https://img.shields.io/github/license/246859/steamapi?color=red)

steamapi 是一个由go编写的 Steam Web API 客户端库，通过它你可以很轻松的与Steam Web API进行交互。
原Steam API的地址在[Steam Web API](https://partner.steamgames.com/doc/webapi)


## 安装
```sh
go get github.com/246859/steamapi
```

## 支持的接口
由于host为`partner.steam-api.com`的接口需要发行商密钥才能够调用，
而发行商密钥必须要100$才能注册，所以我们没法去获取它的具体响应是什么，
于是这类接口统一返回的响应类型为`map[string]any`

- IPublishedFileService
- ISteamNews
- ISteamUser
- IPlayerService
- ISteamWebAPIUtil
- ISteamApps
- ISteamPublishedItemVoting
- ISteamPublishedItemSearch



## 示例

**GetServerInfo** 接口即便不需要api key也能使用，但你仍然需要传递一个可以给client，
为此，你可以使用`steamapi.NopKey`来代替，它只是一个无意义的字符串。该接口可以获取服务的简单信息，通常是用来确认steam服务器能否正常访问。

```go
// initialize client with NopKey
client, err := steamapi.New(steamapi.NopKey)
if err != nil {
panic(err)
}
// call the webapiutil interface
info, err := client.ISteamWebAPIUtil().GetServerInfo()
if err != nil {
panic(err)
}
fmt.Printf("%+v", info)
```
输出
```
{ ServerTime: 1693832041 ServerTimeString: Mon Sep  4 05:54:01 2023 }
```



**GetSupportedAPIList** 接口需要使用你自己注册的api key，否则无法使用。该接口可以获取你的key所能支持使用的所有API列表

```go

// pass your own key
client, err := steamapi.New(key)
if err != nil {
panic(err)
}
list, err := client.ISteamWebAPIUtil().GetSupportedAPIList()
if err != nil {
panic(err)
}
fmt.Printf("%+v", list)
```
输出
```
{ApiList:{Interfaces:[{Name:IClientStats_1046930 Methods:[{Name:ReportEvent Version:1 HttpMethod:POST Parameters:[]}]} {Name:ICSGOPlayers_730...
...
...
...
```



## 贡献

1. fork本仓库
2. 创建你自己的分支
3. 提交修改
4. 向本仓库发起pr
5. 等待合并