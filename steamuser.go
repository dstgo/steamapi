package steamapi

import (
	"github.com/246859/steamapi/types/steam"
	"github.com/246859/steamapi/types/steamuser"
	"net/http"
)

// ISteamUser see https://partner.steamgames.com/doc/webapi/ISteamUser
func (c *Client) ISteamUser() *ISteamUser {
	return &ISteamUser{c}
}

type ISteamUser struct {
	c *Client
}

// CheckAppOwnership see https://partner.steamgames.com/doc/webapi/ISteamUser#CheckAppOwnership
func (i *ISteamUser) CheckAppOwnership(key string, steamId uint64, appId uint32, ops ...RequestOptions) (steamuser.AppOwnership, error) {
	var appOwnership steamuser.AppOwnership
	queryOption := steamuser.OwnershipQueryOption{
		Key: key, AppId: appId, SteamId: steamId,
	}
	_, err := i.c.Get(PartnerHost, steamuser.URLCheckAppOwnership, &appOwnership,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
	if err != nil {
		return appOwnership, err
	}
	return appOwnership, nil
}

// GetAppPriceInfo see https://partner.steamgames.com/doc/webapi/ISteamUser#GetAppPriceInfo
func (i *ISteamUser) GetAppPriceInfo(key string, steamId uint64, appIds string, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := steamuser.AppPriceInfoQueryOption{
		Key: key, SteamId: steamId, AppIds: appIds,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, steamuser.URLGetAppPriceInfo,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// GetDeletedSteamIds see https://partner.steamgames.com/doc/webapi/ISteamUser#GetDeletedSteamIds
func (i *ISteamUser) GetDeletedSteamIds(key string, rowVersion uint64, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := steamuser.DeletedSteamIdsQueryOption{
		Key: key, RowVersion: rowVersion,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, steamuser.URLGetDeletedSteamIds,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// GetFriendList see https://partner.steamgames.com/doc/webapi/ISteamUser#GetFriendList
func (i *ISteamUser) GetFriendList(key string, steamId uint64, relationship string, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := steamuser.FriendListQueryOption{
		Key: key, SteamId: steamId, Relationship: relationship,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, steamuser.URLGetFriendList,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}
