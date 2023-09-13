package test

import (
	"github.com/246859/steamapi"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	client, err := steamapi.New(steamapi.NopKey)
	if err != nil {
		t.Error(err)
	}
	info, err := client.ISteamWebAPIUtil().GetServerInfo()
	if err != nil {
		t.Error(err)
	}
	t.Log(prettyJsonResp(info))
}

func TestGetSupportedAPIList(t *testing.T) {
	key, err := readFileKey("testdata/steamapi.key")
	if err != nil {
		t.Error(err)
	}
	client, err := steamapi.New(key)
	if err != nil {
		t.Error(err)
	}
	list, err := client.ISteamWebAPIUtil().GetSupportedAPIList()
	if err != nil {
		t.Error(err)
	}
	t.Log(prettyJsonResp(list))
}
