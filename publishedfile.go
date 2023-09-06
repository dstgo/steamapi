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
	var response publishedfile.FileList
	request := i.c.Get(PublicHost, publishedfile.URLQueryFiles, joinRequestOptions(options, WithRequestQuery(queryFileOption))...)
	request.SetResult(&response)
	if _, err := request.Send(); err != nil {
		return response, err
	}
	return response, nil
}
