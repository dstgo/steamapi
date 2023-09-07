package steamapi

import (
	"github.com/246859/steamapi/types/publishedfile"
)

func (c *Client) IPublishedFileService() *ISteamPublishedFileService {
	return &ISteamPublishedFileService{c: c}
}

type ISteamPublishedFileService struct {
	c *Client
}

// QueryFiles see https://partner.steamgames.com/doc/webapi/IPublishedFileService#QueryFiles
func (i ISteamPublishedFileService) QueryFiles(queryFileOption publishedfile.FileQueryOption, options ...RequestOptions) (publishedfile.FileList, error) {
	var fileList publishedfile.FileList
	_, err := i.c.Get(PublicHost, publishedfile.URLQueryFiles, &fileList,
		joinRequestOptions(options, WithRequestQuery(queryFileOption))...)
	if err != nil {
		return fileList, err
	}
	return fileList, nil
}
