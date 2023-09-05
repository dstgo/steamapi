package steamapi

import (
	"github.com/246859/steamapi/types/webapiutil"
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
	req := i.c.Get(PublicHost, webapiutil.URLGetServerInfo, options...)
	req.SetResult(&serverInfo)
	if _, err := req.Send(); err != nil {
		return serverInfo, err
	}

	return serverInfo, nil
}

// GetSupportedAPIList see https://partner.steamgames.com/doc/webapi/ISteamWebAPIUtil#GetSupportedAPIList
func (i *ISteamWebAPIUtil) GetSupportedAPIList(options ...RequestOptions) (webapiutil.SteamApiList, error) {
	var list webapiutil.SteamApiList
	request := i.c.Get(PublicHost, webapiutil.URLGetSupportedAPIList, options...)
	request.SetResult(&list)
	if _, err := request.Send(); err != nil {
		return list, err
	}
	return list, nil
}
