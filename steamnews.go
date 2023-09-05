package steamapi

import (
	"github.com/246859/steamapi/types/steamnews"
)

type ISteamNews struct {
	c *Client
}

// GetNewsForApp see https://partner.steamgames.com/doc/webapi/ISteamNews#GetNewsForApp
func (i *ISteamNews) GetNewsForApp(query steamnews.AppNewsQuery, ops ...RequestOptions) (steamnews.AppNewsQueryList, error) {
	var newsList steamnews.AppNewsQueryList
	queryForm, err := structToMap(query)
	if err != nil {
		return newsList, err
	}
	request := i.c.Get(PublicHost, steamnews.URLGetNewsForApp, joinRequestOptions(ops, WithRequestQuery(queryForm))...)
	request.SetResult(&newsList)
	_, err = request.Send()
	if err != nil {
		return newsList, err
	}
	return newsList, nil
}

// GetNewsForAppAuthed see https://partner.steamgames.com/doc/webapi/ISteamNews#GetNewsForAppAuthed
func (i *ISteamNews) GetNewsForAppAuthed(query steamnews.AppNewsQuery, ops ...RequestOptions) (steamnews.AppNewsQueryList, error) {
	return i.GetNewsForApp(query, joinRequestOptions(ops, WithRequestHost(PartnerHost), WithRequestURL(steamnews.URLGetNewsForAppAuthed))...)
}
