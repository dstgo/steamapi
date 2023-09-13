package user

import "github.com/246859/steamapi/types/steam"

type AppPriceInfoQueryOption struct {
	steam.SteamId
	AppIds string `json:"appids" mapstructure:"appids" valid:"required"`
}

type DeletedSteamIdsQueryOption struct {
	RowVersion uint `json:"rowersion" mapstructure:"rowersion" valid:"required"`
}

type FriendListQueryOption struct {
	steam.SteamId
	Relationship string `json:"relationship" mapstructure:"relationship"`
}

type ResolveVanityUrlQueryOption struct {
	VanityUrl string `json:"vanityurl" mapstructure:"vanityurl" valid:"required"`
	UrlType   int    `json:"urltype" mapstructure:"urltype"`
}

type PublisherAppOwnershipChangeQueryOption struct {
	PackageRowVersion string `json:"packagerowversion" mapstructure:"packagerowversion" valid:"required"`
	CdKeyRowVersion   string `json:"cdkeyRowversion" mapstructure:"cdkeyRowversion" valid:"required"`
}

type PlayerSummaryList struct {
	Response struct {
		Players []PlayerSummary `json:"players"`
	} `json:"response"`
}

type PlayerSummary struct {
	SteamId                  string `json:"steamid"`
	CommunityVisibilityState uint   `json:"communityvisibilitystate"`
	ProfileState             uint   `json:"profilestate"`
	PersonName               string `json:"person_name"`
	LastLogoff               uint   `json:"lastlogoff"`
	ProfileUrl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	AvatarMedium             string `json:"avatar_medium"`
	AvatarFull               string `json:"avatar_full"`
}

type AppOwnership struct {
	Ownsapp       bool   `json:"ownsapp"`
	Permanent     bool   `json:"permanent"`
	Timestamp     string `json:"timestamp"`
	OwnersSteamId uint   `json:"ownerssteamid"`
	SiteLicense   bool   `json:"sitelicense"`
}

type PublisherAppOwnership struct {
	SteamId uint `json:"steamid" mapstructure:"steamid"`
	AppOwnership
}

type AppOwnershipList struct {
	AppOwnership struct {
		Apps []AppOwnership `json:"apps"`
	} `json:"appownership"`
}

type PublisherAppOwnershipList struct {
	AppOwnership struct {
		Apps []PublisherAppOwnership `json:"apps"`
	} `json:"appownership"`
}

type AppOwnershipChanges struct {
	OwnershipChanges struct {
		SteamIds          []steam.SteamIdString `json:"steamids" mapstructure:"steamids"`
		PackageRowVersion string                `json:"packagerowversion" mapstructure:"packagerowversion"`
		CdKeyRowVersion   string                `json:"cdkeyRowversion" mapstructure:"cdkeyRowversion"`
		MoreData          bool                  `json:"moredata" mapstructure:"moredata"`
	}
}
