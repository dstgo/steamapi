package steamapi

import (
	"github.com/246859/steamapi/types/webapiutil"
)

const (
	GetServerInfoURL = "/ISteamWebAPIUtil/GetServerInfo/v1/"

	GetSupportedAPIListURL = "/ISteamWebAPIUtil/GetSupportedAPIList/v1/"
)

func (c *Client) WebApiUtil() *ISteamWebAPIUtil {
	return &ISteamWebAPIUtil{c}
}

type ISteamWebAPIUtil struct {
	c *Client
}

// GetServerInfo see https://partner.steamgames.com/doc/webapi/ISteamWebAPIUtil#GetServerInfo
func (i *ISteamWebAPIUtil) GetServerInfo(options ...RequestOptions) (webapiutil.ServerInfo, error) {
	var serverInfo webapiutil.ServerInfo
	_, err := i.c.Get(PublicHost, GetServerInfoURL, options...).SetResult(&serverInfo).Send()
	if err != nil {
		return serverInfo, err
	}

	return serverInfo, nil
}

// GetSupportedAPIList see https://partner.steamgames.com/doc/webapi/ISteamWebAPIUtil#GetSupportedAPIList
func (i *ISteamWebAPIUtil) GetSupportedAPIList(options ...RequestOptions) (webapiutil.SteamApiList, error) {
	var list webapiutil.SteamApiList
	_, err := i.c.Get(PublicHost, GetSupportedAPIListURL, options...).SetResult(&list).Send()
	if err != nil {
		return list, err
	}
	return list, nil
}
