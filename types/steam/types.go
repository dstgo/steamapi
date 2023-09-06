package steam

// CommonResponse
// for some interface which need publisher api key are hard to test
// and no way to know their response, so use this type to replace
type CommonResponse = map[string]any

type APIKey struct {
	Key string `json:"key" mapstructure:"key"`
}

type AppId struct {
	AppId uint64 `json:"appid" mapstructure:"appid"`
}

type AppIds struct {
	AppIds string `json:"appids" mapstructure:"appdis"`
}

type SteamId struct {
	SteamId uint64 `json:"steamid" mapstructure:"steamid"`
}

type SteamIds struct {
	SteamIds string `json:"steamids" mapstructure:"steamids"`
}

type SteamIdString struct {
	SteamId string `json:"steamid" mapstructure:"steamid"`
}
