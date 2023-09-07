package steamuser

type OwnershipQueryOption struct {
	Key     string `json:"key" mapstructure:"key"`
	SteamId uint64 `json:"steamid" mapstructure:"steamid"`
	AppId   uint32 `json:"appid" mapstructure:"appid"`
}

type AppPriceInfoQueryOption struct {
	Key     string `json:"key" mapstructure:"key"`
	SteamId uint64 `json:"steamid" mapstructure:"steamid"`
	AppIds  string `json:"appids" mapstructure:"appids"`
}

type DeletedSteamIdsQueryOption struct {
	Key        string `json:"key" mapstructure:"key"`
	RowVersion uint64 `json:"rowersion" mapstructure:"rowersion"`
}

type FriendListQueryOption struct {
	Key          string `json:"key" mapstructure:"key"`
	SteamId      uint64 `json:"steamid" mapstructure:"steamid"`
	Relationship string `json:"relationship" mapstructure:"relationship"`
}

type ResolveVanityUrlQueryOption struct {
	Key       string `json:"key" mapstructure:"key"`
	VanityUrl string `json:"vanityurl" mapstructure:"vanityurl"`
	UrlType   int    `json:"urltype" mapstructure:"urltype"`
}

type PublisherAppOwnershipChangeQueryOption struct {
	Key               string `json:"key" mapstructure:"key"`
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
	SteamId uint64 `json:"steamid" mapstructure:"steamid"`
	AppOwnership
}

type PublisherAppOwnershipList struct {
	AppOwnership struct {
		Apps []PublisherAppOwnership `json:"apps"`
	} `json:"appownership"`
}

type AppOwnershipChanges struct {
	OwnershipChanges struct {
		SteamIds          string `json:"steamids" mapstructure:"steamids"`
		PackageRowVersion string `json:"packagerowversion" mapstructure:"packagerowversion"`
		CdKeyRowVersion   string `json:"cdkeyRowversion" mapstructure:"cdkeyRowversion"`
		MoreData          bool   `json:"moredata" mapstructure:"moredata"`
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
