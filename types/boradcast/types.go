package boradcast

type DataFrame struct {
	AppId       uint   `json:"appid" mapstructure:"appid" valid:"required"`
	SteamId     uint   `json:"steamid" mapstructure:"steamid" valid:"required"`
	BroadcastId uint   `json:"broadcast_id" mapstructure:"broadcast_id" valid:"required"`
	FrameData   string `json:"frame_data" mapstructure:"frame_data" valid:"required"`
}
