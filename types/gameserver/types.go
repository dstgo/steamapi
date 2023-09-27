package gameserver

import "github.com/246859/steamapi/types/steam"

type GameStatQueryOpt struct {
	steam.AppId
	GameId     uint   `json:"gameid" mapstructure:"gameid"`
	RangeStart string `json:"rangestart" mapstructure:"rangestart"`
	RangeEnd   string `json:"rangeend" mapstructure:"rangeend"`
	MaxResults uint   `json:"maxresults" mapstructure:"maxresults"`
}
