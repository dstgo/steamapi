package steamapi

import (
	"github.com/246859/steamapi/types/boradcast"
	"github.com/246859/steamapi/types/steam"
	"net/http"
)

func (c *Client) ISteamBroadcastService() *ISteamBroadcastService {
	return &ISteamBroadcastService{c: c}
}

// ISteamBroadcastService see https://partner.steamgames.com/doc/webapi/IBroadcastService
type ISteamBroadcastService struct {
	c *Client
}

// PostGameDataFrame see https://partner.steamgames.com/doc/webapi/IBroadcastService#PostGameDataFrame
func (i *ISteamBroadcastService) PostGameDataFrame(frame boradcast.DataFrame, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithInputJson(frame))
	return i.c.Unknown(http.MethodPost, PartnerHost, boradcast.URLPostGameDataFrame, ops...)
}
