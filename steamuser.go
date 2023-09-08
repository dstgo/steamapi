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
func (i *ISteamUser) CheckAppOwnership(key string, steamId uint, appId uint32, ops ...RequestOptions) (steamuser.AppOwnershipList, error) {
	var appOwnership steamuser.AppOwnershipList
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
func (i *ISteamUser) GetAppPriceInfo(key string, steamId uint, appIds string, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := steamuser.AppPriceInfoQueryOption{
		Key: key, SteamId: steamId, AppIds: appIds,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, steamuser.URLGetAppPriceInfo,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// GetDeletedSteamIds see https://partner.steamgames.com/doc/webapi/ISteamUser#GetDeletedSteamIds
func (i *ISteamUser) GetDeletedSteamIds(key string, rowVersion uint, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := steamuser.DeletedSteamIdsQueryOption{
		Key: key, RowVersion: rowVersion,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, steamuser.URLGetDeletedSteamIds,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// GetFriendList see https://partner.steamgames.com/doc/webapi/ISteamUser#GetFriendList
func (i *ISteamUser) GetFriendList(key string, steamId uint, relationship string, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := steamuser.FriendListQueryOption{
		Key: key, SteamId: steamId, Relationship: relationship,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, steamuser.URLGetFriendList,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// GetPlayerBans see https://partner.steamgames.com/doc/webapi/ISteamUser#GetPlayerBans
func (i *ISteamUser) GetPlayerBans(key string, steamids string, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := steamuser.SteamIdsQueryOption{
		Key: key, SteamIds: steamids,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, steamuser.URLGetPlayerBans,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// GetPlayerSummaries see https://partner.steamgames.com/doc/webapi/ISteamUser#GetPlayerSummaries
func (i *ISteamUser) GetPlayerSummaries(key string, steamids string, ops ...RequestOptions) (steamuser.PlayerSummaryList, error) {
	var summaryList steamuser.PlayerSummaryList
	queryOption := steamuser.SteamIdsQueryOption{
		Key: key, SteamIds: steamids,
	}
	_, err := i.c.Get(PublicHost, steamuser.URLGetPlayerSummaries, &summaryList,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...,
	)
	if err != nil {
		return summaryList, err
	}
	return summaryList, nil
}

// GetPublisherAppOwnership see https://partner.steamgames.com/doc/webapi/ISteamUser#GetPublisherAppOwnership
func (i *ISteamUser) GetPublisherAppOwnership(key string, steamId uint, ops ...RequestOptions) (steamuser.PublisherAppOwnershipList, error) {
	var publisherAppOwnership steamuser.PublisherAppOwnershipList
	queryOption := steamuser.SteamIdQueryOption{
		Key: key, SteamId: steamId,
	}
	_, err := i.c.Get(PublicHost, steamuser.URLGetPublisherAppOwnership, &publisherAppOwnership,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
	if err != nil {
		return publisherAppOwnership, err
	}
	return publisherAppOwnership, nil
}

// GetPublisherAppOwnershipChanges see https://partner.steamgames.com/doc/webapi/ISteamUser#GetPublisherAppOwnershipChanges
func (i *ISteamUser) GetPublisherAppOwnershipChanges(chaneQueryOption steamuser.PublisherAppOwnershipChangeQueryOption, ops ...RequestOptions) (steamuser.AppOwnershipChanges, error) {
	var changes steamuser.AppOwnershipChanges
	_, err := i.c.Get(PartnerHost, steamuser.URLGetPublisherAppOwnershipChanges, &changes,
		joinRequestOptions(ops, WithRequestQuery(chaneQueryOption))...)
	if err != nil {
		return changes, err
	}
	return changes, nil
}

// GetUserGroupList see https://partner.steamgames.com/doc/webapi/ISteamUser#GetUserGroupList
func (i *ISteamUser) GetUserGroupList(key string, steamId uint, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := steamuser.SteamIdQueryOption{Key: key, SteamId: steamId}
	return i.c.Unknown(http.MethodGet, PartnerHost, steamuser.URLGetUserGroupList,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// ResolveVanityURL see https://partner.steamgames.com/doc/webapi/ISteamUser#ResolveVanityURL
func (i *ISteamUser) ResolveVanityURL(key string, vanityUrl string, urlType int, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := steamuser.ResolveVanityUrlQueryOption{Key: key, VanityUrl: vanityUrl, UrlType: urlType}
	return i.c.Unknown(http.MethodGet, PartnerHost, steamuser.URLGetUserGroupList,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}
