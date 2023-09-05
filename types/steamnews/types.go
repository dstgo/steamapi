package steamnews

type AppNewsQuery struct {
	AppId     uint   `json:"appid" mapstructure:"appid"`
	MaxLength uint   `json:"maxlength"  mapstructure:"maxlength"`
	EndDate   uint   `json:"enddate" mapstructure:"enddate"`
	Count     uint   `json:"count" mapstructure:"count"`
	Feeds     string `json:"feeds" mapstructure:"feeds"`
}

type AppNewsQueryList struct {
	AppNews struct {
		AppId     uint      `json:"appid"`
		NewsItems []AppNews `json:"newsitems"`
	} `json:"appnews"`
}

type AppNews struct {
	Gid           string   `json:"gid"`
	AppId         uint     `json:"appid"`
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
