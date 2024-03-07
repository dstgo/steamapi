package gameserver

import "github.com/dstgo/steamapi/types/steam"

type GameStatQueryOpt struct {
	steam.AppId
	GameId     uint   `json:"gameid" mapstructure:"gameid"`
	RangeStart string `json:"rangestart" mapstructure:"rangestart"`
	RangeEnd   string `json:"rangeend" mapstructure:"rangeend"`
	MaxResults uint   `json:"maxresults" mapstructure:"maxresults"`
}
