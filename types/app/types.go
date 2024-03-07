package app

import "github.com/dstgo/steamapi/types/steam"

type BuildsQueryOption struct {
	steam.AppId
	Count uint `json:"count" mapstructure:"count" valid:"required"`
}

type PartnerAppQueryFilter struct {
	TypeFilter string `json:"type_filter" mapstructure:"type_filter"`
}

type ServerListQueryFilter struct {
	Filter string `json:"filter" mapstructure:"filter"`
	Limit  uint   `json:"limit" mapstructure:"limit"`
}

type ServerAddressQueryOption struct {
	Address string `json:"addr" mapstructure:"addr" valid:"required,ip"`
}

type PublicApp struct {
	AppId uint   `json:"appid"`
	Name  string `json:"name"`
}

type PartnerApp struct {
	AppId   uint   `json:"appid"`
	AppName string `json:"app_name"`
	AppType string `json:"app_type"`
}

type PublicAppList struct {
	AppList struct {
		Apps []PublicApp `json:"apps"`
	} `json:"applist"`
}

type PartnerAppList struct {
	AppList struct {
		Apps struct {
			App []PartnerApp `json:"app"`
		} `json:"apps"`
	} `json:"applist"`
}

type ServerAddress struct {
	Address  string `json:"addr"`
	GmsIndex int    `json:"gmsindex"`
	SteamId  string `json:"steamid"`
	Message  string `json:"message"`
	AppId    uint   `json:"appid"`
	GameDir  string `json:"gamedir"`
	Region   uint   `json:"region"`
	Secure   bool   `json:"secure"`
	Lan      bool   `json:"lan"`
	GamePort uint   `json:"gameport"`
	SpecPort uint   `json:"specport"`
}

type ServerAddressList struct {
	Response struct {
		Success bool            `json:"success"`
		Servers []ServerAddress `json:"servers"`
	} `json:"response"`
}

type BuildLiveUpdateOption struct {
	AppId       uint   `json:"appid" valid:"required"`
	BuildId     uint   `json:"buildid" valid:"required"`
	BetaKey     string `json:"betakey" valid:"required"`
	Description string `json:"description"`
}

type UpToDateCheckQueryOption struct {
	AppId   uint `json:"appid" mapstructure:"appid" valid:"required"`
	Version uint `json:"version" mapstructure:"version" valid:"required"`
}

type UpToDateCheck struct {
	Response struct {
		Success           bool   `json:"success"`
		UpToDate          bool   `json:"up_to_date"`
		VersionIsListable bool   `json:"version_is_listable"`
		RequiredVersion   uint   `json:"required_version"`
		Message           string `json:"message"`
	} `json:"response"`
}
