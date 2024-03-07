package steamapi

import (
	"github.com/dstgo/steamapi/types/remotestorage"
	"github.com/dstgo/steamapi/types/steam"
	"net/http"
)

func (c *Client) ISteamRemoteStorage() *ISteamRemoteStorage {
	return &ISteamRemoteStorage{c: c}
}

// ISteamRemoteStorage see https://partner.steamgames.com/doc/webapi/ISteamRemoteStorage
type ISteamRemoteStorage struct {
	c *Client
}

// EnumerateUserSubscribedFiles see https://partner.steamgames.com/doc/webapi/ISteamRemoteStorage#EnumerateUserSubscribedFiles
func (i *ISteamRemoteStorage) EnumerateUserSubscribedFiles(enum remotestorage.EnumerateOpt, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithFormData(enum))
	return i.c.Unknown(http.MethodPost, PartnerHost, remotestorage.URLEnumerateUserSubscribedFiles, ops...)
}

// GetCollectionDetails see https://partner.steamgames.com/doc/webapi/ISteamRemoteStorage#GetCollectionDetails
func (i *ISteamRemoteStorage) GetCollectionDetails(query remotestorage.CollectionDetailQueryOpt, ops ...RequestOption) (remotestorage.CollectionDetail, error) {
	var detail remotestorage.CollectionDetail
	ops = joinRequestOptions(ops, WithFormData(query))
	_, err := i.c.Post(PublicHost, remotestorage.URLGetCollectionDetails, &detail, ops...)
	if err != nil {
		return detail, err
	}
	return detail, nil
}

// GetPublishedFileDetails see https://partner.steamgames.com/doc/webapi/ISteamRemoteStorage#GetPublishedFileDetails
func (i *ISteamRemoteStorage) GetPublishedFileDetails(query remotestorage.PublishedFileDetailQueryOpt, ops ...RequestOption) (remotestorage.PublishedFileDetail, error) {
	var detail remotestorage.PublishedFileDetail
	ops = joinRequestOptions(ops, WithFormData(query))
	_, err := i.c.Post(PublicHost, remotestorage.URLGetPublishedFileDetails, &detail, ops...)
	if err != nil {
		return detail, err
	}
	return detail, nil
}

// GetUGCFileDetails see https://partner.steamgames.com/doc/webapi/ISteamRemoteStorage#GetUGCFileDetails
func (i *ISteamRemoteStorage) GetUGCFileDetails(query remotestorage.UgcFileDetailQueryOpt, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithQueryForm(query))
	return i.c.Unknown(http.MethodGet, PublicHost, remotestorage.URLGetUGCFileDetails, ops...)
}

// SetUGCUsedByGC see https://partner.steamgames.com/doc/webapi/ISteamRemoteStorage#SetUGCUsedByGC
func (i *ISteamRemoteStorage) SetUGCUsedByGC(ugc remotestorage.UgcFileUpdateOpt, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithFormData(ugc))
	return i.c.Unknown(http.MethodPost, PartnerHost, remotestorage.URLSetUGCUsedByGC, ops...)
}

// SubscribePublishedFile see https://partner.steamgames.com/doc/webapi/ISteamRemoteStorage#SubscribePublishedFile
func (i *ISteamRemoteStorage) SubscribePublishedFile(sub remotestorage.SubscribePublishedFileOpt, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithFormData(sub))
	return i.c.Unknown(http.MethodPost, PartnerHost, remotestorage.URLSubscribePublishedFile, ops...)
}

// UnSubscribePublishedFile see https://partner.steamgames.com/doc/webapi/ISteamRemoteStorage#UnSubscribePublishedFile
func (i *ISteamRemoteStorage) UnSubscribePublishedFile(unsub remotestorage.UnSubscribePublishedFileOpt, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithFormData(unsub))
	return i.c.Unknown(http.MethodPost, PartnerHost, remotestorage.URLUnsubscribePublishedFile, ops...)
}
