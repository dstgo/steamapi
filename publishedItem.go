package steamapi

import (
	"github.com/246859/steamapi/types/publisheditem"
	"github.com/246859/steamapi/types/steam"
	"net/http"
)

func (c *Client) ISteamPublishedSearch() *ISteamPublishedItemSearch {
	return &ISteamPublishedItemSearch{c: c}
}

// ISteamPublishedItemSearch see https://partner.steamgames.com/doc/webapi/ISteamPublishedItemSearch
type ISteamPublishedItemSearch struct {
	c *Client
}

// RankedByPublicationOrder see https://partner.steamgames.com/doc/webapi/ISteamPublishedItemSearch#RankedByPublicationOrder
func (i *ISteamPublishedItemSearch) RankedByPublicationOrder(rankOpt publisheditem.RankOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithBody(rankOpt))
	return i.c.Unknown(http.MethodPost, PartnerHost, publisheditem.URLRankedByPublicationOrder, ops...)
}

// RankedByTrend see https://partner.steamgames.com/doc/webapi/ISteamPublishedItemSearch#RankedByTrend
func (i *ISteamPublishedItemSearch) RankedByTrend(rankOpt publisheditem.RankOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithBody(rankOpt))
	return i.c.Unknown(http.MethodPost, PartnerHost, publisheditem.URLRankedByTrend, ops...)
}

// RankedByVote see https://partner.steamgames.com/doc/webapi/ISteamPublishedItemSearch#RankedByVote
func (i *ISteamPublishedItemSearch) RankedByVote(rankOpt publisheditem.RankOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithBody(rankOpt))
	return i.c.Unknown(http.MethodPost, PartnerHost, publisheditem.URLRankedByVote, ops...)
}

// ResultSetSummary see https://partner.steamgames.com/doc/webapi/ISteamPublishedItemSearch#ResultSetSummary
func (i *ISteamPublishedItemSearch) ResultSetSummary(rankOpt publisheditem.RankSummaryOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithBody(rankOpt))
	return i.c.Unknown(http.MethodPost, PartnerHost, publisheditem.URLResultSetSummary, ops...)
}

func (c *Client) ISteamPublishedItemVoting() *ISteamPublishedItemVoting {
	return &ISteamPublishedItemVoting{c: c}
}

// ISteamPublishedItemVoting see https://partner.steamgames.com/doc/webapi/ISteamPublishedItemVoting
type ISteamPublishedItemVoting struct {
	c *Client
}

// ItemVoteSummary see https://partner.steamgames.com/doc/webapi/ISteamPublishedItemVoting#ItemVoteSummary
func (i *ISteamPublishedItemVoting) ItemVoteSummary(voteOpt publisheditem.ItemVoteOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithBody(voteOpt))
	return i.c.Unknown(http.MethodPost, PartnerHost, publisheditem.URLItemVoteSummary, ops...)
}

// UserVoteSummary see https://partner.steamgames.com/doc/webapi/ISteamPublishedItemVoting#UserVoteSummary
func (i *ISteamPublishedItemVoting) UserVoteSummary(voteOpt publisheditem.UserVoteOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithBody(voteOpt))
	return i.c.Unknown(http.MethodPost, PartnerHost, publisheditem.URLUserVoteSummary, ops...)
}
