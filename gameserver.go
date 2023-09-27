package steamapi

import (
	"github.com/246859/steamapi/types/gameserver"
	"github.com/246859/steamapi/types/steam"
	"net/http"
)

func (c *Client) ISteamGameServerStats() *ISteamGameServerStats {
	return &ISteamGameServerStats{c: c}
}

// ISteamGameServerStats see https://partner.steamgames.com/doc/webapi/ISteamGameServerStats
type ISteamGameServerStats struct {
	c *Client
}

// GetGameServerPlayerStatsForGame see https://partner.steamgames.com/doc/webapi/ISteamGameServerStats#GetGameServerPlayerStatsForGame
func (i *ISteamGameServerStats) GetGameServerPlayerStatsForGame(query gameserver.GameStatQueryOpt, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithQueryForm(query))
	return i.c.Unknown(http.MethodGet, PartnerHost, gameserver.URLGetGameServerPlayerStatsForGame, ops...)
}
