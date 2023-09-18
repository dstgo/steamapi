package steamapi

import (
	"github.com/246859/steamapi/types/leaderboard"
	"github.com/246859/steamapi/types/steam"
	"net/http"
)

func (c *Client) ISteamLeaderBoards() *ISteamLeaderBoards {
	return &ISteamLeaderBoards{}
}

// ISteamLeaderBoards see https://partner.steamgames.com/doc/webapi/ISteamLeaderboards
type ISteamLeaderBoards struct {
	c *Client
}

// DeleteLeaderBoard see https://partner.steamgames.com/doc/webapi/ISteamLeaderboards#DeleteLeaderBoard
func (i *ISteamLeaderBoards) DeleteLeaderBoard(deleteOpt leaderboard.DeleteOpt, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithBody(deleteOpt))
	return i.c.Unknown(http.MethodPost, PartnerHost, leaderboard.URLDeleteLeaderboard, ops...)
}

// FindOrCreateLeaderboard see https://partner.steamgames.com/doc/webapi/ISteamLeaderboards#FindOrCreateLeaderboard
func (i *ISteamLeaderBoards) FindOrCreateLeaderboard(createOpt leaderboard.CreateOpt, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithBody(createOpt))
	return i.c.Unknown(http.MethodPost, PartnerHost, leaderboard.URLFindOrCreateLeaderboard, ops...)
}

// GetLeaderboardEntries see https://partner.steamgames.com/doc/webapi/ISteamLeaderboards#GetLeaderboardEntries
func (i *ISteamLeaderBoards) GetLeaderboardEntries(entryOpt leaderboard.EntriesQueryOpt, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithQueryForm(entryOpt))
	return i.c.Unknown(http.MethodGet, PartnerHost, leaderboard.URLGetLeaderboardEntries, ops...)
}

// GetLeaderboardsForGame see https://partner.steamgames.com/doc/webapi/ISteamLeaderboards#GetLeaderboardsForGame
func (i *ISteamLeaderBoards) GetLeaderboardsForGame(appid uint, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithQueryForm(steam.AppId{AppId: appid}))
	return i.c.Unknown(http.MethodGet, PartnerHost, leaderboard.URLGetLeaderboardsForGame, ops...)
}

// ResetLeaderboard see https://partner.steamgames.com/doc/webapi/ISteamLeaderboards#ResetLeaderboard
func (i *ISteamLeaderBoards) ResetLeaderboard(opt leaderboard.ResetOpt, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithBody(opt))
	return i.c.Unknown(http.MethodPost, PartnerHost, leaderboard.URLResetLeaderboard, ops...)
}

// SetLeaderboardScore see https://partner.steamgames.com/doc/webapi/ISteamLeaderboards#SetLeaderboardScore
func (i *ISteamLeaderBoards) SetLeaderboardScore(opt leaderboard.ScoreSetOpt, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithBody(opt))
	return i.c.Unknown(http.MethodPost, PartnerHost, leaderboard.URLSetLeaderboardScore, ops...)
}
