package steam

import (
	"github.com/spf13/cast"
	"strings"
)

// CommonResponse
// for some interface which need publisher api key are hard to test
// and no way to know their response, so use this type to replace
type CommonResponse = map[string]any

type APIKey struct {
	Key string `json:"key" mapstructure:"key"`
}

type AppId struct {
	AppId uint `json:"appid" mapstructure:"appid"`
}

type SteamId struct {
	SteamId uint `json:"steamid" mapstructure:"steamid"`
}

type SteamIdString struct {
	SteamId string `json:"steamid" mapstructure:"steamid"`
}

type SteamIds struct {
	SteamIds string `json:"steamids" mapstructure:"steaids" valid:"required"`
}

type AppIds struct {
	AppIds string `json:"appids" mapstructure:"appids" valid:"required"`
}

type SteamAppId struct {
	SteamId uint `json:"steamid" mapstructure:"steamid" valid:"required"`
	AppId   uint `json:"appid" mapstructure:"appid" valid:"required"`
}

// Ids join several uint id with the ","
func Ids(ids ...uint64) string {
	var res []string
	for _, id := range ids {
		res = append(res, cast.ToString(id))
	}
	return strings.Join(res, ",")
}
