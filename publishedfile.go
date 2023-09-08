package steamapi

import (
	"github.com/246859/steamapi/types/publishedfile"
	"github.com/246859/steamapi/types/steam"
	"net/http"
)

func (c *Client) IPublishedFileService() *ISteamPublishedFileService {
	return &ISteamPublishedFileService{c: c}
}

type ISteamPublishedFileService struct {
	c *Client
}

// QueryFiles see https://partner.steamgames.com/doc/webapi/IPublishedFileService#QueryFiles
func (i *ISteamPublishedFileService) QueryFiles(queryFileOption publishedfile.FileQueryOption, options ...RequestOptions) (publishedfile.FileList, error) {
	var fileList publishedfile.FileList
	_, err := i.c.Get(PublicHost, publishedfile.URLQueryFiles, &fileList,
		joinRequestOptions(options, WithRequestQuery(queryFileOption))...)
	if err != nil {
		return fileList, err
	}
	return fileList, nil
}

// SetDeveloperMetaData see https://partner.steamgames.com/doc/webapi/IPublishedFileService#SetDeveloperMetadata
func (i *ISteamPublishedFileService) SetDeveloperMetaData(key string, setOption publishedfile.DeveloperMetaSetOption, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := steam.APIKey{Key: key}
	bodyOption := setOption
	ops = joinRequestOptions(ops, WithRequestQuery(queryOption), WithRequestBody(bodyOption))
	return i.c.Unknown(http.MethodPost, PartnerHost, publishedfile.URLSetDeveloperMetadata, ops...)
}

// UpdateUGCBan see https://partner.steamgames.com/doc/webapi/IPublishedFileService#UpdateUGCBan
func (i *ISteamPublishedFileService) UpdateUGCBan(key string, updateOption publishedfile.BanStatusUpdateOption, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := steam.APIKey{Key: key}
	bodyOption := updateOption
	ops = joinRequestOptions(ops, WithRequestQuery(queryOption), WithRequestBody(bodyOption))
	return i.c.Unknown(http.MethodPost, PartnerHost, publishedfile.URLUpdateAppUGCBan, ops...)
}

// UpdateBanStatus see https://partner.steamgames.com/doc/webapi/IPublishedFileService#UpdateBanStatus
func (i *ISteamPublishedFileService) UpdateBanStatus(key string, updateOption publishedfile.BanStatusUpdateOption, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := steam.APIKey{Key: key}
	bodyOption := updateOption
	ops = joinRequestOptions(ops, WithRequestQuery(queryOption), WithRequestBody(bodyOption))
	return i.c.Unknown(http.MethodPost, PartnerHost, publishedfile.URLUpdateBanStatus, ops...)
}

// UpdateIncompatibleStatus see https://partner.steamgames.com/doc/webapi/IPublishedFileService#UpdateIncompatibleStatus
func (i *ISteamPublishedFileService) UpdateIncompatibleStatus(key string, updateOption publishedfile.IncompatibleStatusUpdateOption, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := steam.APIKey{Key: key}
	bodyOption := updateOption
	ops = joinRequestOptions(ops, WithRequestQuery(queryOption), WithRequestBody(bodyOption))
	return i.c.Unknown(http.MethodPost, PartnerHost, publishedfile.URLUpdateIncompatibleStatus, ops...)
}

// UpdateTags see https://partner.steamgames.com/doc/webapi/IPublishedFileService#UpdateTags
func (i *ISteamPublishedFileService) UpdateTags(key string, updateOption publishedfile.TagUpdateOption, ops ...RequestOptions) (steam.CommonResponse, error) {
	queryOption := steam.APIKey{Key: key}
	bodyOption := updateOption
	ops = joinRequestOptions(ops, WithRequestQuery(queryOption), WithRequestBody(bodyOption))
	return i.c.Unknown(http.MethodPost, PartnerHost, publishedfile.URLUpdateTags, ops...)
}
