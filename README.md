English|[简体中文](./README.zh.md)

# steamapi
steamapi is a Steam Web API Go package,By using it, you can easily interact with the Steam web API.



## Install
```sh
go get github.com/246859/steamapi
```

## api doc
The original Steam API address is [Steam Web API](https://partner.steamgames.com/doc/webapi), but you can also go to [steamapi doc](https://apifox.com/apidoc/shared-1a2822b1-0e88-4df1-b7ad-08acfd783cf2)
View interface documents, support preview interface response and online debugging.

## Supported Interface

- ISteamWebAPIUtil



## Example

**GetServerInfo** interface has no apikey required, but you should pass a key still.In this case, 
you can pass a `steamapi.NopKey` to New func which just a meaningless string.

```go
// initialize client with NopKey
client, err := steamapi.New(steamapi.NopKey)
if err != nil {
    panic(err)
}
// call the webapiutil interface
info, err := client.WebApiUtil().GetServerInfo()
if err != nil {
    panic(err)
}
fmt.Printf("%+v", info)
```
output
```
{ ServerTime: 1693832041 ServerTimeString: Mon Sep  4 05:54:01 2023 }
```



**GetSupportedAPIList** interface needs to pass your own apikey

```go

// pass your own key
client, err := steamapi.New(key)
if err != nil {
    panic(err)
}
list, err := client.WebApiUtil().GetSupportedAPIList()
if err != nil {
    panic(err)
}
fmt.Printf("%+v", list)
```
output
```
{ApiList:{Interfaces:[{Name:IClientStats_1046930 Methods:[{Name:ReportEvent Version:1 HttpMethod:POST Parameters:[]}]} {Name:ICSGOPlayers_730...
...
...
...
```



## Contribute

1. fork this repository
2. create you own feature branch
3. commit your changes
4. create a pull request to this repository
5. waiting pr to be merged