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
func (i *ISteamPublishedFileService) QueryFiles(queryFileOption publishedfile.FileQueryOption, ops ...RequestOption) (publishedfile.FileList, error) {
	var fileList publishedfile.FileList
	ops = joinRequestOptions(ops, WithQueryForm(queryFileOption))
	_, err := i.c.Get(PublicHost, publishedfile.URLQueryFiles, &fileList, ops...)
	if err != nil {
		return fileList, err
	}
	return fileList, nil
}

// SetDeveloperMetaData see https://partner.steamgames.com/doc/webapi/IPublishedFileService#SetDeveloperMetadata
func (i *ISteamPublishedFileService) SetDeveloperMetaData(devSetOpt publishedfile.DeveloperMetaSetOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithInputJson(devSetOpt))
	return i.c.Unknown(http.MethodPost, PartnerHost, publishedfile.URLSetDeveloperMetadata, ops...)
}

// UpdateUGCBan see https://partner.steamgames.com/doc/webapi/IPublishedFileService#UpdateUGCBan
func (i *ISteamPublishedFileService) UpdateUGCBan(banUpdateOpt publishedfile.BanStatusUpdateOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithInputJson(banUpdateOpt))
	return i.c.Unknown(http.MethodPost, PartnerHost, publishedfile.URLUpdateAppUGCBan, ops...)
}

// UpdateBanStatus see https://partner.steamgames.com/doc/webapi/IPublishedFileService#UpdateBanStatus
func (i *ISteamPublishedFileService) UpdateBanStatus(banStatusUpdateOpt publishedfile.BanStatusUpdateOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithInputJson(banStatusUpdateOpt))
	return i.c.Unknown(http.MethodPost, PartnerHost, publishedfile.URLUpdateBanStatus, ops...)
}

// UpdateIncompatibleStatus see https://partner.steamgames.com/doc/webapi/IPublishedFileService#UpdateIncompatibleStatus
func (i *ISteamPublishedFileService) UpdateIncompatibleStatus(updateOpt publishedfile.IncompatibleStatusUpdateOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithInputJson(updateOpt))
	return i.c.Unknown(http.MethodPost, PartnerHost, publishedfile.URLUpdateIncompatibleStatus, ops...)
}

// UpdateTags see https://partner.steamgames.com/doc/webapi/IPublishedFileService#UpdateTags
func (i *ISteamPublishedFileService) UpdateTags(tagUpdateOpt publishedfile.TagUpdateOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithInputJson(tagUpdateOpt))
	return i.c.Unknown(http.MethodPost, PartnerHost, publishedfile.URLUpdateTags, ops...)
}
