package types

type ServerInfo struct {
	ServerTime       uint64 `json:"servertime"`
	ServerTimeString string `json:"servertimestring"`
}

type SteamApiList struct {
	ApiList struct {
		Interfaces []SteamInterfaceGroup `json:"interfaces"`
	} `json:"apiList"`
}

type SteamInterfaceGroup struct {
	Name    string           `json:"name"`
	Methods []SteamInterface `json:"methods"`
}

type SteamInterfaceParam struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Optional    bool   `json:"optional"`
	Description string `json:"description"`
}

type SteamInterface struct {
	Name       string                `json:"name"`
	Version    int                   `json:"version"`
	HttpMethod string                `json:"httpmethod"`
	Parameters []SteamInterfaceParam `json:"parameters"`
}
