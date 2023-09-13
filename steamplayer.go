package steamapi

import (
	"github.com/246859/steamapi/types/player"
	"github.com/246859/steamapi/types/steam"
	"net/http"
)

func (c *Client) IPlayerService() *IPlayerService {
	return &IPlayerService{c}
}

// IPlayerService see https://partner.steamgames.com/doc/webapi/IPlayerService
type IPlayerService struct {
	c *Client
}

// GetRecentlyPlayedGames see https://partner.steamgames.com/doc/webapi/IPlayerService#GetRecentlyPlayedGames
func (i *IPlayerService) GetRecentlyPlayedGames(steamId uint, count uint, ops ...RequestOption) (steam.CommonResponse, error) {
	queryOption := player.PlayedGamesQueryOption{
		SteamId: steam.SteamId{SteamId: steamId},
		Count:   count,
	}
	ops = joinRequestOptions(ops, WithQueryForm(queryOption))
	return i.c.Unknown(http.MethodGet, PartnerHost, player.URLGetRecentlyPlayedGames, ops...)
}

// GetSingleGamePlayTime see https://partner.steamgames.com/doc/webapi/IPlayerService#GetSingleGamePlaytime
func (i *IPlayerService) GetSingleGamePlayTime(steamId uint, appId uint, ops ...RequestOption) (steam.CommonResponse, error) {
	queryOption := steam.SteamAppId{
		SteamId: steamId,
		AppId:   appId,
	}
	ops = joinRequestOptions(ops, WithQueryForm(queryOption))
	return i.c.Unknown(http.MethodGet, PartnerHost, player.URLGetSingleGamePlaytime, ops...)
}

// GetOwnedGames see https://partner.steamgames.com/doc/webapi/IPlayerService#GetOwnedGames
func (i *IPlayerService) GetOwnedGames(ownGameOpt player.OwnedGamesQueryOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithQueryForm(ownGameOpt))
	return i.c.Unknown(http.MethodGet, PartnerHost, player.URLGetOwnedGames, ops...)
}

// GetSteamLevel see https://partner.steamgames.com/doc/webapi/IPlayerService#GetSteamLevel
func (i *IPlayerService) GetSteamLevel(steamId uint, ops ...RequestOption) (steam.CommonResponse, error) {
	queryOption := steam.SteamId{SteamId: steamId}
	ops = joinRequestOptions(ops, WithQueryForm(queryOption))
	return i.c.Unknown(http.MethodGet, PartnerHost, player.URLGetSteamLevel, ops...)
}

// GetBadges see https://partner.steamgames.com/doc/webapi/IPlayerService#GetBadges
func (i *IPlayerService) GetBadges(steamId uint, ops ...RequestOption) (steam.CommonResponse, error) {
	queryOption := steam.SteamId{SteamId: steamId}
	ops = joinRequestOptions(ops, WithQueryForm(queryOption))
	return i.c.Unknown(http.MethodGet, PartnerHost, player.URLGetBadges, ops...)
}

// GetCommunityBadgeProgress see https://partner.steamgames.com/doc/webapi/IPlayerService#GetCommunityBadgeProgress
func (i *IPlayerService) GetCommunityBadgeProgress(steamId uint, badgeId int, ops ...RequestOption) (steam.CommonResponse, error) {
	var queryOption = player.CommunityBadgeProgress{
		SteamId: steam.SteamId{SteamId: steamId},
		BadgeId: badgeId,
	}
	ops = joinRequestOptions(ops, WithQueryForm(queryOption))
	return i.c.Unknown(http.MethodGet, PartnerHost, player.URLGetCommunityBadgeProgress, ops...)
}
