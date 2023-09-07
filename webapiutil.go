package steamapi

import (
	"github.com/246859/steamapi/types/steam"
	"github.com/246859/steamapi/types/webapiutil"
)

func (c *Client) ISteamWebAPIUtil() *ISteamWebAPIUtil {
	return &ISteamWebAPIUtil{c}
}

type ISteamWebAPIUtil struct {
	c *Client
}

// GetServerInfo see https://partner.steamgames.com/doc/webapi/ISteamWebAPIUtil#GetServerInfo
func (i *ISteamWebAPIUtil) GetServerInfo(options ...RequestOptions) (webapiutil.ServerInfo, error) {
	var serverInfo webapiutil.ServerInfo
	_, err := i.c.Get(PublicHost, webapiutil.URLGetServerInfo, &serverInfo, options...)
	if err != nil {
		return serverInfo, err
	}
	return serverInfo, nil
}

// GetSupportedAPIList see https://partner.steamgames.com/doc/webapi/ISteamWebAPIUtil#GetSupportedAPIList
func (i *ISteamWebAPIUtil) GetSupportedAPIList(key string, options ...RequestOptions) (webapiutil.SteamApiList, error) {
	var list webapiutil.SteamApiList
	_, err := i.c.Get(PublicHost, webapiutil.URLGetSupportedAPIList, &list,
		joinRequestOptions(options, WithRequestQuery(steam.APIKey{Key: key}))...)
	if err != nil {
		return list, err
	}
	return list, nil
}
