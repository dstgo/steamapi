package steamapi

import (
	"github.com/246859/steamapi/types/news"
)

func (c *Client) ISteamNews() *ISteamNews {
	return &ISteamNews{c}
}

type ISteamNews struct {
	c *Client
}

// GetNewsForApp see https://partner.steamgames.com/doc/webapi/ISteamNews#GetNewsForApp
func (i *ISteamNews) GetNewsForApp(query news.AppNewsQueryOption, ops ...RequestOption) (news.AppNewsList, error) {
	var newsList news.AppNewsList
	ops = joinRequestOptions(ops, WithQueryForm(query))
	_, err := i.c.Get(PublicHost, news.URLGetNewsForApp, &newsList, ops...)
	if err != nil {
		return newsList, err
	}
	return newsList, nil
}

// GetNewsForAppAuthed see https://partner.steamgames.com/doc/webapi/ISteamNews#GetNewsForAppAuthed
func (i *ISteamNews) GetNewsForAppAuthed(query news.AppNewsQueryOption, ops ...RequestOption) (news.AppNewsList, error) {
	return i.GetNewsForApp(query, joinRequestOptions(ops, WithHost(PartnerHost), WithURL(news.URLGetNewsForAppAuthed))...)
}
