package publishedfile

import "github.com/246859/steamapi/types/steam"

type FileQueryOption struct {
	Key                    string                  `json:"key" mapstructure:"key"`
	Language               steam.LanguageCode      `json:"language" mapstructure:"language"`
	QueryType              EPublishedFileQueryType `json:"query_type" mapstructure:"query_type"`
	Page                   uint                    `json:"page" mapstructure:"page"`
	Cursor                 string                  `json:"cursor" mapstructure:"cursor"`
	NumPerPage             uint                    `json:"numperpage" mapstructure:"numperpage"`
	CreatorAppId           uint                    `json:"creator_appid" mapstructure:"creator_appid"`
	AppID                  uint                    `json:"appid" mapstructure:"appid"`
	RequiredTags           string                  `json:"requiredtags" mapstructure:"requiredtags"`
	ExcludedTags           string                  `json:"excludedtags" mapstructure:"excludedtags"`
	MatchAllTags           bool                    `json:"match_all_tags" mapstructure:"match_all_tags"`
	RequiredFlags          string                  `json:"required_flags" mapstructure:"required_flags"`
	OmittedFlags           string                  `json:"omitted_flags" mapstructure:"omitted_flags"`
	SearchText             string                  `json:"search_text" mapstructure:"search_text"`
	FileType               uint                    `json:"filetype" mapstructure:"filetype"`
	ChildPublishedFileId   uint                    `json:"child_publishedfileid" mapstructure:"child_publishedfileid"`
	Days                   uint                    `json:"days" mapstructure:"days"`
	IncludeRecentVotesOnly bool                    `json:"include_recent_votes_only" mapstructure:"include_recent_votes_only"`
	CacheMaxAgeSeconds     uint                    `json:"cache_max_age_seconds" mapstructure:"cache_max_age_seconds"`
	TotalOnly              bool                    `json:"totalonly" mapstructure:"totalonly"`
	IdsOnly                bool                    `json:"ids_only" mapstructure:"ids_only"`
	ReturnVoteData         bool                    `json:"return_vote_data" mapstructure:"return_vote_data"`
	ReturnTags             bool                    `json:"return_tags" mapstructure:"return_tags"`
	ReturnKvTags           bool                    `json:"return_kv_tags" mapstructure:"return_kv_tags"`
	ReturnPreviews         bool                    `json:"return_previews" mapstructure:"return_previews"`
	ReturnChildren         bool                    `json:"return_children" mapstructure:"return_children"`
	ReturnShortDescription bool                    `json:"return_short_description" mapstructure:"return_short_description"`
	ReturnForSaleData      bool                    `json:"return_for_sale_data" mapstructure:"return_for_sale_data"`
	ReturnMetaData         bool                    `json:"return_metadata" mapstructure:"return_metadata"`
	ReturnPlayTimeStats    uint                    `json:"return_playtime_stats" mapstructure:"return_playtime_stats"`
}

type FileList struct {
	Response struct {
		Total                int    `json:"total"`
		PublishedFileDetails []File `json:"publishedfiledetails"`
	} `json:"response"`
}

type File struct {
	Result                     uint          `json:"result"`
	PublishedFileId            string        `json:"publishedfileid"`
	Creator                    string        `json:"creator"`
	CreatorAppId               uint          `json:"creator_appid"`
	ConsumerAppId              uint          `json:"consume_appid"`
	ConsumerShortcutId         uint          `json:"consumer_shortcutid"`
	Filename                   string        `json:"filename"`
	FileSize                   string        `json:"file_size"`
	PreviewFileSize            string        `json:"preview_file_size"`
	PreviewUrl                 string        `json:"preview_url"`
	Url                        string        `json:"url"`
	HContentFile               string        `json:"hcontent_file"`
	HContentPreview            string        `json:"hcontent_preview"`
	Title                      string        `json:"title"`
	FileDescription            string        `json:"file_description"`
	Language                   int           `json:"language"`
	TimeCreated                uint64        `json:"time_created"`
	TimeUpdated                uint64        `json:"time_updated"`
	Visibility                 uint          `json:"visibility"`
	Flags                      uint          `json:"flags"`
	WorkshopFile               bool          `json:"workshop_file"`
	WorkShopAccepted           bool          `json:"workshop_accepted"`
	ShowSubscribeAll           bool          `json:"show_subscribe_all"`
	NumCommentsPublic          uint          `json:"num_comments_public"`
	Banned                     bool          `json:"banned"`
	BanReason                  string        `json:"ban_reason"`
	Banner                     string        `json:"banner"`
	CanBeDeleted               bool          `json:"can_be_deleted"`
	AppName                    string        `json:"app_name"`
	FileType                   uint          `json:"file_type"`
	CanSubscribe               bool          `json:"can_subscribe"`
	Subscriptions              uint          `json:"subscriptions"`
	Favorited                  uint          `json:"favorited"`
	Followers                  uint          `json:"followers"`
	LifeTimeSubscriptions      uint          `json:"lifetime_subscriptions"`
	LifeTimeFavorited          uint          `json:"lifetime_favorited"`
	LifeTimeFollowers          uint          `json:"lifetime_followers"`
	LifeTimePlayTime           string        `json:"lifetime_playtime"`
	LifeTimePlayTimeSessions   string        `json:"lifetime_playtime_sessions"`
	Views                      uint          `json:"views"`
	NumChildren                uint          `json:"num_children"`
	NumReports                 uint          `json:"num_reports"`
	Previews                   []FilePreview `json:"previews"`
	Tags                       []FileTag     `json:"tags"`
	KVTags                     []FileKvTag   `json:"kvtags"`
	VoteData                   FileVoteData  `json:"vote_data"`
	MaybeInappropriateSex      bool          `json:"maybe_inappropriate_sex"`
	MaybeInappropriateViolence bool          `json:"maybe_inappropriate_violence"`
	RevisionChangeNumber       string        `json:"revision_change_number"`
	Revision                   uint          `json:"revision"`
	BanTextCheckResult         uint          `json:"ban_text_check_result"`
}

type FilePreview struct {
	PreviewId   string `json:"previewid"`
	SortOrder   int    `json:"sortorder"`
	Url         string `json:"url"`
	Size        int    `json:"size"`
	Filename    string `json:"filename"`
	PreviewType int    `json:"preview_type"`
}

type FileTag struct {
	Tag         string `json:"tag"`
	DisplayName string `json:"display_name"`
}

type FileKvTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type FileVoteData struct {
	Score    float64 `json:"score"`
	VoteUp   int     `json:"vote_up"`
	VoteDown int     `json:"vote_down"`
}
