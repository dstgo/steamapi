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
func (i *IPlayerService) GetRecentlyPlayedGames(key string, steamId uint, count uint, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := player.PlayedGamesQueryOption{
		ApiKeyId: steam.ApiKeyId{
			Key: key, SteamId: steamId,
		},
		Count: count,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, player.URLGetRecentlyPlayedGames, joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// GetSingleGamePlayTime see https://partner.steamgames.com/doc/webapi/IPlayerService#GetSingleGamePlaytime
func (i *IPlayerService) GetSingleGamePlayTime(key string, steamId uint, appId uint, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := steam.SteamAppKeyId{
		Key:     key,
		SteamId: steamId,
		AppId:   appId,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, player.URLGetSingleGamePlaytime, joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// GetOwnedGames see https://partner.steamgames.com/doc/webapi/IPlayerService#GetOwnedGames
func (i *IPlayerService) GetOwnedGames(option player.OwnedGamesQueryOption, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := option
	return i.c.Unknown(http.MethodGet, PartnerHost, player.URLGetOwnedGames, joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// GetSteamLevel see https://partner.steamgames.com/doc/webapi/IPlayerService#GetSteamLevel
func (i *IPlayerService) GetSteamLevel(key string, steamId uint, ops ...RequestOptions) (steam.CommonResponse, error) {
	var queryOption = steam.ApiKeyId{
		Key:     key,
		SteamId: steamId,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, player.URLGetSteamLevel, joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// GetBadges see https://partner.steamgames.com/doc/webapi/IPlayerService#GetBadges
func (i *IPlayerService) GetBadges(key string, steamId uint, ops ...RequestOptions) (steam.CommonResponse, error) {
	var queryOption = steam.ApiKeyId{
		Key:     key,
		SteamId: steamId,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, player.URLGetBadges, joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}

// GetCommunityBadgeProgress see https://partner.steamgames.com/doc/webapi/IPlayerService#GetCommunityBadgeProgress
func (i *IPlayerService) GetCommunityBadgeProgress(key string, steamId uint, badgeId int, ops ...RequestOptions) (steam.CommonResponse, error) {
	var queryOption = player.CommunityBadgeProgress{
		ApiKeyId: steam.ApiKeyId{Key: key, SteamId: steamId},
		BadgeId:  badgeId,
	}
	return i.c.Unknown(http.MethodGet, PartnerHost, player.URLGetCommunityBadgeProgress, joinRequestOptions(ops, WithRequestQuery(queryOption))...)
}
