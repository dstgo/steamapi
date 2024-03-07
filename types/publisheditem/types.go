package publisheditem

import "github.com/dstgo/steamapi/types/publishedfile"

type RankOption struct {
	StartIndex        uint `json:"startidx" mapstructure:"startidx"`
	HasAppAdminAccess bool `json:"hasappadminaccess" mapstructure:"hasappadminaccess"`
	RankSummaryOption
}

type RankSummaryOption struct {
	SteamId uint `json:"steamid" mapstructure:"steamid" valid:"required"`
	AppId   uint `json:"appid" mapstructure:"appid" valid:"required"`

	Count        uint `json:"count" mapstructure:"count"`
	TagCount     uint `json:"tagcount" mapstructure:"tagcount"`
	UserTagCount uint `json:"usertagcount" mapstructure:"usertagCount"`

	FileType publishedfile.EPublishedFileInfoMatchingFileType `json:"filetype" mapstructure:"filetype"`
	Tag0     string                                           `json:"tag[0]" mapstructure:"tag[0]"`
	UserTag0 string                                           `json:"usertag[0]" mapstructure:"usertag[0]"`
}

type ItemVoteOption struct {
	SteamId        uint `json:"steamid" mapstructure:"steamid" valid:"required"`
	AppId          uint `json:"appid" mapstructure:"appid" valid:"required"`
	Count          uint `json:"count" mapstructure:"count"`
	PublishedFile0 uint `json:"publishedfile[0]" mapstructure:"publishedfile[0]"`
}

type UserVoteOption struct {
	SteamId        uint `json:"steamid" mapstructure:"steamid" valid:"required"`
	Count          uint `json:"count" mapstructure:"count"`
	PublishedFile0 uint `json:"publishedfile[0]" mapstructure:"publishedfile[0]"`
}
