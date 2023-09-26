package remotestorage

import (
	"github.com/246859/steamapi/types/publishedfile"
	"github.com/246859/steamapi/types/steam"
)

type EnumerateOpt struct {
	steam.SteamId
	steam.AppId
	ListType uint `json:"listtype" mapstructure:"listtype"`
}

type CollectionDetailQueryOpt struct {
	CollectionCount   uint `json:"collectioncount" mapstructure:"collectioncount" valid:"required"`
	PublishedFileIds0 uint `json:"publishedfileids[0]" mapstructure:"publishedfileids" valid:"required"`
}

type PublishedFileDetailQueryOpt struct {
	ItemCount         uint `json:"itemcount" mapstructure:"itemcount" valid:"required"`
	PublishedFileIds0 uint `json:"publishedfileids[0]" mapstructure:"publishedfileids" valid:"required"`
}

type UgcFileDetailQueryOpt struct {
	steam.SteamId
	steam.AppId
	UgcId uint `json:"ugcid" mapstructure:"ugcid" valid:"required"`
}

type UgcFileUpdateOpt struct {
	steam.SteamId
	steam.AppId
	UgcId uint `json:"ugcid" mapstructure:"ugcid" valid:"required"`
	Used  bool `json:"used" mapstructure:"used"`
}

type SubscribePublishedFileOpt struct {
	steam.SteamId
	steam.AppId
	PublishedFileId uint `json:"publishedfileid" mapstructure:"publishedfileid"`
}

type UnSubscribePublishedFileOpt struct {
	steam.SteamId
	steam.AppId
	PublishedFileId uint `json:"publishedfileid" mapstructure:"publishedfileid"`
}

type PublishedFileDetail struct {
	Response struct {
		Result               int                  `json:"result"`
		ResultCount          int                  `json:"resultcount"`
		PublishedFileDetails []publishedfile.File `json:"publishedfiledetails"`
	} `json:"response"`
}

type Collection struct {
	PublishedFileId string `json:"publishedfileid"`
	Result          int    `json:"result"`
}

type CollectionDetail struct {
	Response struct {
		Result               int          `json:"result"`
		ResultCount          int          `json:"resultcount"`
		PublishedFileDetails []Collection `json:"publishedfiledetails"`
	} `json:"response"`
}
