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
func (i *ISteamUser) CheckAppOwnership(ids steam.SteamAppId, ops ...RequestOption) (user.AppOwnershipList, error) {
	var appOwnership user.AppOwnershipList
	ops = joinRequestOptions(ops, WithQueryForm(ids))
	_, err := i.c.Get(PartnerHost, user.URLCheckAppOwnership, &appOwnership, ops...)
	if err != nil {
		return appOwnership, err
	}
	return appOwnership, nil
}

// GetAppPriceInfo see https://partner.steamgames.com/doc/webapi/ISteamUser#GetAppPriceInfo
func (i *ISteamUser) GetAppPriceInfo(steamId uint, appIds string, ops ...RequestOption) (steam.CommonResponse, error) {
	queryOption := user.AppPriceInfoQueryOption{SteamId: steam.SteamId{SteamId: steamId}, AppIds: appIds}
	ops = joinRequestOptions(ops, WithQueryForm(queryOption))
	return i.c.Unknown(http.MethodGet, PartnerHost, user.URLGetAppPriceInfo, ops...)
}

// GetDeletedSteamIds see https://partner.steamgames.com/doc/webapi/ISteamUser#GetDeletedSteamIds
func (i *ISteamUser) GetDeletedSteamIds(rowVersion uint, ops ...RequestOption) (steam.CommonResponse, error) {
	queryOption := user.DeletedSteamIdsQueryOption{RowVersion: rowVersion}
	ops = joinRequestOptions(ops, WithQueryForm(queryOption))
	return i.c.Unknown(http.MethodGet, PartnerHost, user.URLGetDeletedSteamIds, ops...)
}

// GetFriendList see https://partner.steamgames.com/doc/webapi/ISteamUser#GetFriendList
func (i *ISteamUser) GetFriendList(relation user.FriendListQueryOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithQueryForm(relation))
	return i.c.Unknown(http.MethodGet, PartnerHost, user.URLGetFriendList, ops...)
}

// GetPlayerBans see https://partner.steamgames.com/doc/webapi/ISteamUser#GetPlayerBans
func (i *ISteamUser) GetPlayerBans(steamids string, ops ...RequestOption) (steam.CommonResponse, error) {
	queryOption := steam.SteamIds{SteamIds: steamids}
	ops = joinRequestOptions(ops, WithQueryForm(queryOption))
	return i.c.Unknown(http.MethodGet, PartnerHost, user.URLGetPlayerBans, ops...)
}

// GetPlayerSummaries see https://partner.steamgames.com/doc/webapi/ISteamUser#GetPlayerSummaries
func (i *ISteamUser) GetPlayerSummaries(steamids string, ops ...RequestOption) (user.PlayerSummaryList, error) {
	var summaryList user.PlayerSummaryList
	queryOption := steam.SteamIds{SteamIds: steamids}
	ops = joinRequestOptions(ops, WithQueryForm(queryOption))
	_, err := i.c.Get(PublicHost, user.URLGetPlayerSummaries, &summaryList, ops...)
	if err != nil {
		return summaryList, err
	}
	return summaryList, nil
}

// GetPublisherAppOwnership see https://partner.steamgames.com/doc/webapi/ISteamUser#GetPublisherAppOwnership
func (i *ISteamUser) GetPublisherAppOwnership(steamId uint, ops ...RequestOption) (user.PublisherAppOwnershipList, error) {
	var publisherAppOwnership user.PublisherAppOwnershipList
	queryOption := steam.SteamId{SteamId: steamId}
	ops = joinRequestOptions(ops, WithQueryForm(queryOption))
	_, err := i.c.Get(PublicHost, user.URLGetPublisherAppOwnership, &publisherAppOwnership, ops...)
	if err != nil {
		return publisherAppOwnership, err
	}
	return publisherAppOwnership, nil
}

// GetPublisherAppOwnershipChanges see https://partner.steamgames.com/doc/webapi/ISteamUser#GetPublisherAppOwnershipChanges
func (i *ISteamUser) GetPublisherAppOwnershipChanges(chaneQueryOption user.PublisherAppOwnershipChangeQueryOption, ops ...RequestOption) (user.AppOwnershipChanges, error) {
	var changes user.AppOwnershipChanges
	ops = joinRequestOptions(ops, WithQueryForm(chaneQueryOption))
	_, err := i.c.Get(PartnerHost, user.URLGetPublisherAppOwnershipChanges, &changes, ops...)
	if err != nil {
		return changes, err
	}
	return changes, nil
}

// GetUserGroupList see https://partner.steamgames.com/doc/webapi/ISteamUser#GetUserGroupList
func (i *ISteamUser) GetUserGroupList(steamId uint, ops ...RequestOption) (steam.CommonResponse, error) {
	queryOption := steam.SteamId{SteamId: steamId}
	ops = joinRequestOptions(ops, WithQueryForm(queryOption))
	return i.c.Unknown(http.MethodGet, PartnerHost, user.URLGetUserGroupList, ops...)
}

// ResolveVanityURL see https://partner.steamgames.com/doc/webapi/ISteamUser#ResolveVanityURL
func (i *ISteamUser) ResolveVanityURL(urlResolve user.ResolveVanityUrlQueryOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithQueryForm(urlResolve))
	return i.c.Unknown(http.MethodGet, PartnerHost, user.URLResolveVanityURL, ops...)
}

func (c *Client) ISteamUserAuth() *ISteamUserAuth {
	return &ISteamUserAuth{c: c}
}

// ISteamUserAuth see https://partner.steamgames.com/doc/webapi/ISteamUserAuth
type ISteamUserAuth struct {
	c *Client
}

// AuthenticateUser see https://partner.steamgames.com/doc/webapi/ISteamUserAuth#AuthenticateUser
func (i *ISteamUserAuth) AuthenticateUser(authenticateOpt user.AuthenticateOpt, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithFormData(authenticateOpt))
	return i.c.Unknown(http.MethodPost, PartnerHost, user.URLAuthenticateUser, ops...)
}

// AuthenticateUserTicket see https://partner.steamgames.com/doc/webapi/ISteamUserAuth#AuthenticateUserTicket
func (i *ISteamUserAuth) AuthenticateUserTicket(ticketAuthenticateOpt user.TicketAuthenticateOpt, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithQueryForm(ticketAuthenticateOpt))
	return i.c.Unknown(http.MethodGet, PartnerHost, user.URLAuthenticateUserTicket, ops...)
}
