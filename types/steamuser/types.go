package steamuser

import "github.com/246859/steamapi/types/steam"

type SteamApiIdOption struct {
	steam.APIKey
	steam.SteamId
}

type SteamApiIdsOption struct {
	steam.APIKey
	steam.SteamIds
}

type AppOwnershipOption struct {
	SteamApiIdOption
	steam.AppId
}

type AppsOwnershipOption struct {
	SteamApiIdOption
	steam.AppIds
}

type DeletedSteamIdsQueryOption struct {
	steam.APIKey
	RowVersion uint64 `json:"rowersion" mapstructure:"rowersion"`
}

type FriendListQueryOption struct {
	steam.APIKey
	steam.SteamId
	Relationship string `json:"relationship" mapstructure:"relationship"`
}

type ResolveVanityUrlQueryOption struct {
	steam.APIKey
	VanityUrl string `json:"vanityurl" mapstructure:"vanityurl"`
	UrlType   int    `json:"urltype" mapstructure:"urltype"`
}

type PublisherAppOwnershipChangeQueryOption struct {
	steam.APIKey
	PackageRowVersion string `json:"packagerowversion" mapstructure:"packagerowversion"`
	CdKeyRowVersion   string `json:"cdkeyRowversion" mapstructure:"cdkeyRowversion"`
}

type PlayerSummaryList struct {
	Response struct {
		Players []PlayerSummary `json:"players"`
	} `json:"response"`
}

type AppOwnership struct {
	Ownsapp       bool   `json:"ownsapp"`
	Permanent     bool   `json:"permanent"`
	Timestamp     string `json:"timestamp"`
	OwnersSteamId uint64 `json:"ownerssteamid"`
	SiteLicense   bool   `json:"sitelicense"`
}

type PublisherAppOwnership struct {
	steam.AppId
	AppOwnership
}

type AppOwnershipChanges struct {
	OwnershipChanges struct {
		SteamIds          steam.SteamIdString `json:"steamids"`
		PackageRowVersion string              `json:"packagerowversion" mapstructure:"packagerowversion"`
		CdKeyRowVersion   string              `json:"cdkeyRowversion" mapstructure:"cdkeyRowversion"`
		MoreData          bool                `json:"moredata" mapstructure:"moredata"`
	}
}

type PlayerSummary struct {
	SteamId                  string `json:"steamid"`
	CommunityVisibilityState uint   `json:"communityvisibilitystate"`
	ProfileState             uint   `json:"profilestate"`
	PersonName               string `json:"person_name"`
	LastLogoff               uint64 `json:"lastlogoff"`
	ProfileUrl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	AvatarMedium             string `json:"avatar_medium"`
	AvatarFull               string `json:"avatar_full"`
}
