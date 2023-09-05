package test

import (
	"fmt"
	"github.com/246859/steamapi"
	"github.com/246859/steamapi/types/publishedfile"
	"github.com/246859/steamapi/types/steam"
	"testing"
)

func TestQueryFile(t *testing.T) {
	key, err := readFileKey()
	if err != nil {
		panic(err)
	}
	client, err := steamapi.New(key)
	if err != nil {
		panic(err)
	}
	files, err := client.PublishedFiles().QueryFiles(publishedfile.QueryFileRequest{
		Page:       1,
		NumPerPage: 10,
		AppID:      108600,

		SearchText:          "",
		Language:            steam.LangSimplifiedChinese,
		ReturnVoteData:      true,
		ReturnTags:          true,
		ReturnKvTags:        true,
		ReturnPreviews:      true,
		ReturnChildren:      true,
		ReturnForSaleData:   true,
		ReturnMetaData:      true,
		ReturnPlayTimeStats: 0,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(prettyJsonResp(files))
}
