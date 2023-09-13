package steamapi

import (
	"github.com/246859/steamapi/types/webapiutil"
)

func (c *Client) ISteamWebAPIUtil() *ISteamWebAPIUtil {
	return &ISteamWebAPIUtil{c}
}

type ISteamWebAPIUtil struct {
	c *Client
}

// GetServerInfo see https://partner.steamgames.com/doc/webapi/ISteamWebAPIUtil#GetServerInfo
func (i *ISteamWebAPIUtil) GetServerInfo(ops ...RequestOption) (webapiutil.ServerInfo, error) {
	var serverInfo webapiutil.ServerInfo
	_, err := i.c.Get(PublicHost, webapiutil.URLGetServerInfo, &serverInfo, ops...)
	if err != nil {
		return serverInfo, err
	}
	return serverInfo, nil
}

// GetSupportedAPIList see https://partner.steamgames.com/doc/webapi/ISteamWebAPIUtil#GetSupportedAPIList
func (i *ISteamWebAPIUtil) GetSupportedAPIList(ops ...RequestOption) (webapiutil.SteamApiList, error) {
	var list webapiutil.SteamApiList
	_, err := i.c.Get(PublicHost, webapiutil.URLGetSupportedAPIList, &list, ops...)
	if err != nil {
		return list, err
	}
	return list, nil
}
