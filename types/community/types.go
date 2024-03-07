package community

import "github.com/dstgo/steamapi/types/steam"

type ReportAbuseForm struct {
	steam.AppId
	SteamIdActor  uint   `json:"steamidActor" mapstructure:"steamidActor"`
	SteamIdTarget uint   `json:"steamidTarget" mapstructure:"steamidTarget"`
	AbuseType     uint   `json:"abuseType" mapstructure:"abuseType"`
	ContentType   uint   `json:"contentType" mapstructure:"contentType"`
	Description   string `json:"description" mapstructure:"description"`
	Gid           uint   `json:"gid" mapstructure:"gid"`
}
