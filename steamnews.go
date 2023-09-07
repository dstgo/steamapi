package steamapi

import (
	"github.com/246859/steamapi/types/steamnews"
)

func (c *Client) ISteamNews() *ISteamNews {
	return &ISteamNews{c}
}

type ISteamNews struct {
	c *Client
}

// GetNewsForApp see https://partner.steamgames.com/doc/webapi/ISteamNews#GetNewsForApp
func (i *ISteamNews) GetNewsForApp(query steamnews.AppNewsQueryOption, ops ...RequestOptions) (steamnews.AppNewsList, error) {
	var newsList steamnews.AppNewsList
	_, err := i.c.Get(PublicHost, steamnews.URLGetNewsForApp, &newsList, joinRequestOptions(ops, WithRequestQuery(query))...)
	if err != nil {
		return newsList, err
	}
	return newsList, nil
}

// GetNewsForAppAuthed see https://partner.steamgames.com/doc/webapi/ISteamNews#GetNewsForAppAuthed
func (i *ISteamNews) GetNewsForAppAuthed(query steamnews.AppNewsQueryOption, ops ...RequestOptions) (steamnews.AppNewsList, error) {
	return i.GetNewsForApp(query, joinRequestOptions(ops, WithRequestHost(PartnerHost), WithRequestURL(steamnews.URLGetNewsForAppAuthed))...)
}
