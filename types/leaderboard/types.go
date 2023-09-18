package leaderboard

type DeleteOpt struct {
	Appid uint   `json:"appid" mapstructure:"appid" valid:"required"`
	Name  string `json:"name" mapstructure:"name"`
}

type CreateOpt struct {
	Appid             uint   `json:"appid" mapstructure:"appid" valid:"required"`
	Name              string `json:"name" mapstructure:"name"`
	SortMethod        string `json:"sortmethod" mapstructure:"sortmethod"`
	DisplayType       string `json:"displaytype" mapstructure:"displaytype"`
	CreateIfNotFound  bool   `json:"createifnotfound" mapstructure:"createifnotfound"`
	OnlyTrustedWrites bool   `json:"onlytrustedwrite" mapstructure:"onlytrustedwrite"`
	OnlyFriendSreads  bool   `json:"onlyfriendsreads" mapstructure:"onlyfriendsreads"`
}

type EntriesQueryOpt struct {
	Appid         uint `json:"appid" mapstructure:"appid" valid:"required"`
	SteamId       uint `json:"steamid" mapstructure:"steamid"`
	LeaderBoardId int  `json:"leaderboardid" mapstructure:"leaderboardid" valid:"required"`
	RangeStart    int  `json:"rangestart" mapstructure:"rangestart"`
	RangeEnd      int  `json:"rangeend" mapstructure:"rangeend"`
	DataRequest   uint `json:"datarequest" mapstructure:"datarequest" valid:"required"`
}

type ResetOpt struct {
	AppId         uint `json:"appid" mapstructure:"appid" valid:"required"`
	LeaderBoardId uint `json:"leaderboardid" mapstructure:"leaderboardid" valid:"required"`
}

type ScoreSetOpt struct {
	AppId         uint   `json:"appid" mapstructure:"appid" valid:"required"`
	SteamdId      uint   `json:"steamdid" mapstructure:"steamdid" valid:"required"`
	LeaderBoardId uint   `json:"leaderboardid" mapstructure:"leaderboardid" valid:"required"`
	Score         int    `json:"score" mapstructure:"score" valid:"required"`
	ScoreMethod   string `json:"scoremethod" mapstructure:"scoremethod" valid:"required"`
	Details       []byte `json:"details" mapstructure:"details"`
}
