package steamapi

import (
	"github.com/246859/steamapi/types/steam"
	"github.com/246859/steamapi/types/user"
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
func (i *ISteamUser) CheckAppOwnership(key string, steamId uint, appId uint32, ops ...RequestOptions) (user.AppOwnershipList, error) {
	var appOwnership user.AppOwnershipList
	queryOption := user.OwnershipQueryOption{
		Key: key, AppId: appId, SteamId: steamId,
	}
	_, err := i.c.Get(PartnerHost, user.URLCheckAppOwnership, &appOwnership,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
	if err != nil {
		return appOwnership, err
	}
	return appOwnership, nil
}

// GetAppPriceInfo see https://partner.steamgames.com/doc/webapi/ISteamUser#GetAppPriceInfo
func (i *ISteamUser) GetAppPriceInfo(key string, steamId uint, appIds string, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := user.AppPriceInfoQueryOption{
		Key: key, SteamId: steamId, AppIds: appIds,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, user.URLGetAppPriceInfo,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// GetDeletedSteamIds see https://partner.steamgames.com/doc/webapi/ISteamUser#GetDeletedSteamIds
func (i *ISteamUser) GetDeletedSteamIds(key string, rowVersion uint, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := user.DeletedSteamIdsQueryOption{
		Key: key, RowVersion: rowVersion,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, user.URLGetDeletedSteamIds,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// GetFriendList see https://partner.steamgames.com/doc/webapi/ISteamUser#GetFriendList
func (i *ISteamUser) GetFriendList(key string, steamId uint, relationship string, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := user.FriendListQueryOption{
		Key: key, SteamId: steamId, Relationship: relationship,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, user.URLGetFriendList,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// GetPlayerBans see https://partner.steamgames.com/doc/webapi/ISteamUser#GetPlayerBans
func (i *ISteamUser) GetPlayerBans(key string, steamids string, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := user.SteamIdsQueryOption{
		Key: key, SteamIds: steamids,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, user.URLGetPlayerBans,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// GetPlayerSummaries see https://partner.steamgames.com/doc/webapi/ISteamUser#GetPlayerSummaries
func (i *ISteamUser) GetPlayerSummaries(key string, steamids string, ops ...RequestOptions) (user.PlayerSummaryList, error) {
	var summaryList user.PlayerSummaryList
	queryOption := user.SteamIdsQueryOption{
		Key: key, SteamIds: steamids,
	}
	_, err := i.c.Get(PublicHost, user.URLGetPlayerSummaries, &summaryList,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...,
	)
	if err != nil {
		return summaryList, err
	}
	return summaryList, nil
}

// GetPublisherAppOwnership see https://partner.steamgames.com/doc/webapi/ISteamUser#GetPublisherAppOwnership
func (i *ISteamUser) GetPublisherAppOwnership(key string, steamId uint, ops ...RequestOptions) (user.PublisherAppOwnershipList, error) {
	var publisherAppOwnership user.PublisherAppOwnershipList
	queryOption := user.SteamIdQueryOption{
		Key: key, SteamId: steamId,
	}
	_, err := i.c.Get(PublicHost, user.URLGetPublisherAppOwnership, &publisherAppOwnership,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
	if err != nil {
		return publisherAppOwnership, err
	}
	return publisherAppOwnership, nil
}

// GetPublisherAppOwnershipChanges see https://partner.steamgames.com/doc/webapi/ISteamUser#GetPublisherAppOwnershipChanges
func (i *ISteamUser) GetPublisherAppOwnershipChanges(chaneQueryOption user.PublisherAppOwnershipChangeQueryOption, ops ...RequestOptions) (user.AppOwnershipChanges, error) {
	var changes user.AppOwnershipChanges
	_, err := i.c.Get(PartnerHost, user.URLGetPublisherAppOwnershipChanges, &changes,
		joinRequestOptions(ops, WithRequestQuery(chaneQueryOption))...)
	if err != nil {
		return changes, err
	}
	return changes, nil
}

// GetUserGroupList see https://partner.steamgames.com/doc/webapi/ISteamUser#GetUserGroupList
func (i *ISteamUser) GetUserGroupList(key string, steamId uint, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := user.SteamIdQueryOption{Key: key, SteamId: steamId}
	return i.c.Unknown(http.MethodGet, PartnerHost, user.URLGetUserGroupList,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// ResolveVanityURL see https://partner.steamgames.com/doc/webapi/ISteamUser#ResolveVanityURL
func (i *ISteamUser) ResolveVanityURL(key string, vanityUrl string, urlType int, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := user.ResolveVanityUrlQueryOption{Key: key, VanityUrl: vanityUrl, UrlType: urlType}
	return i.c.Unknown(http.MethodGet, PartnerHost, user.URLResolveVanityURL,
		joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}
