package user

import "github.com/246859/steamapi/types/steam"

type OwnershipQueryOption struct {
	Key     string `json:"key" mapstructure:"key"`
	SteamId uint   `json:"steamid" mapstructure:"steamid" valid:"required"`
	AppId   uint32 `json:"appid" mapstructure:"appid" valid:"required"`
}

type AppPriceInfoQueryOption struct {
	Key     string `json:"key" mapstructure:"key"`
	SteamId uint   `json:"steamid" mapstructure:"steamid" valid:"required"`
	AppIds  string `json:"appids" mapstructure:"appids" valid:"required"`
}

type DeletedSteamIdsQueryOption struct {
	Key        string `json:"key" mapstructure:"key"`
	RowVersion uint   `json:"rowersion" mapstructure:"rowersion" valid:"required"`
}

type FriendListQueryOption struct {
	Key          string `json:"key" mapstructure:"key"`
	SteamId      uint   `json:"steamid" mapstructure:"steamid" valid:"required"`
	Relationship string `json:"relationship" mapstructure:"relationship"`
}

type SteamIdsQueryOption struct {
	Key      string `json:"key" mapstructure:"key"`
	SteamIds string `json:"steamids" mapstructure:"steamids" valid:"required"`
}

type SteamIdQueryOption struct {
	Key     string `json:"key" mapstructure:"key"`
	SteamId uint   `json:"steamid" mapstructure:"steamid" valid:"required"`
}

type ResolveVanityUrlQueryOption struct {
	Key       string `json:"key" mapstructure:"key"`
	VanityUrl string `json:"vanityurl" mapstructure:"vanityurl" valid:"required"`
	UrlType   int    `json:"urltype" mapstructure:"urltype"`
}

type PublisherAppOwnershipChangeQueryOption struct {
	Key               string `json:"key" mapstructure:"key"`
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
