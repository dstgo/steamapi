package player

import "github.com/246859/steamapi/types/steam"

type PlayedGamesQueryOption struct {
	steam.ApiKeyId
	Count uint `json:"count" mapstructure:"count"`
}
type OwnedGamesQueryOption struct {
	steam.ApiKeyId
	IncludeAppInfo         bool `json:"include_app_info" mapstructure:"include_app_info"`
	IncludePlayedFreeGames bool `json:"include_played_free_games" mapstructure:"include_played_free_games"`
	AppIdsFilter           uint `json:"app_ids_filter" mapstructure:"app_ids_filter"`
}

type CommunityBadgeProgress struct {
	steam.ApiKeyId
	BadgeId int `json:"badgeid" mapstructure:"badgeid"`
}
