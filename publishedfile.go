package steamapi

import (
	"github.com/246859/steamapi/types/publishedfile"
)

func (c *Client) PublishedFiles() *ISteamPublishedFilesService {
	return &ISteamPublishedFilesService{c: c}
}

type ISteamPublishedFilesService struct {
	c *Client
}

// QueryFiles see https://partner.steamgames.com/doc/webapi/IPublishedFileService#QueryFiles
func (i ISteamPublishedFilesService) QueryFiles(queryFileRequest publishedfile.FileQueryOption, options ...RequestOptions) (publishedfile.FileList, error) {
	var response publishedfile.FileList
	queryform, err := structToMap(queryFileRequest)
	if err != nil {
		return response, err
	}
	options = joinRequestOptions(options, WithRequestQuery(queryform))
	request := i.c.Get(PublicHost, publishedfile.URLQueryFiles, options...)
	request.SetResult(&response)
	if _, err = request.Send(); err != nil {
		return response, err
	}
	return response, nil
}
