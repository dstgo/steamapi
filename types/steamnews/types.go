package steamnews

import "github.com/246859/steamapi/types/steam"

type AppNewsQueryOption struct {
	AppId     uint   `json:"appid" mapstructure:"appid" valid:"required"`
	MaxLength uint   `json:"maxlength"  mapstructure:"maxlength"`
	EndDate   uint   `json:"enddate" mapstructure:"enddate"`
	Count     uint   `json:"count" mapstructure:"count"`
	Feeds     string `json:"feeds" mapstructure:"feeds"`
}

type AppNewsList struct {
	AppNews struct {
		steam.AppId
		NewsItems []AppNews `json:"newsitems"`
	} `json:"appnews"`
}

type AppNews struct {
	steam.AppId
	Gid           string   `json:"gid"`
	Title         string   `json:"title"`
	Url           string   `json:"url"`
	IsExternalUrl bool     `json:"is_external_url"`
	Author        string   `json:"author"`
	Contents      string   `json:"contents"`
	FeedLabel     string   `json:"feed_label"`
	Date          uint     `json:"date"`
	FeedName      string   `json:"feedname"`
	FeedType      uint     `json:"feed_type"`
	Tags          []string `json:"tags"`
}
