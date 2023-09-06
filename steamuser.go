package steamapi

import (
	"github.com/246859/steamapi/types/steam"
	"github.com/246859/steamapi/types/steamuser"
)

func (c *Client) ISteamUser() *ISteamUser {
	return &ISteamUser{c}
}

type ISteamUser struct {
	c *Client
}

// CheckAppOwnership see https://partner.steamgames.com/doc/webapi/ISteamUser#CheckAppOwnership
func (i *ISteamUser) CheckAppOwnership(key string, steamId, appId uint64, ops ...RequestOptions) (steamuser.AppOwnership, error) {
	var appOwnership steamuser.AppOwnership
	queryOption := steamuser.AppOwnershipOption{
		SteamApiIdOption: steamuser.SteamApiIdOption{
			APIKey:  steam.APIKey{Key: key},
			SteamId: steam.SteamId{SteamId: steamId},
		},
		AppId: steam.AppId{AppId: appId},
	}
	request := i.c.Get(PartnerHost, steamuser.URLCheckAppOwnership, joinRequestOptions(ops, WithRequestQuery(queryOption))...)
	request.SetResult(&appOwnership)
	if _, err := request.Send(); err != nil {
		return appOwnership, err
	}
	return appOwnership, nil
}
