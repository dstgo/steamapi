package steamapi

import (
	"github.com/dstgo/steamapi/types/community"
	"github.com/dstgo/steamapi/types/steam"
	"net/http"
)

func (c *Client) ISteamCommunity() *ISteamCommunity {
	return &ISteamCommunity{c: c}
}

// ISteamCommunity see https://partner.steamgames.com/doc/webapi/ISteamCommunity
type ISteamCommunity struct {
	c *Client
}

// ReportAbuse see https://partner.steamgames.com/doc/webapi/ISteamCommunity#ReportAbuse
func (i *ISteamCommunity) ReportAbuse(form community.ReportAbuseForm, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithFormData(form))
	return i.c.Unknown(http.MethodPost, PartnerHost, "", ops...)
}
