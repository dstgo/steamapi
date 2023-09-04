package test

import (
	"encoding/json"
	"github.com/246859/steamapi"
	"os"
	"testing"
)

func readFileKey() (string, error) {
	bytes, err := os.ReadFile("testdata/steamapi.key")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func prettyJsonResp(v any) string {
	bytes, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return ""
	}
	return string(bytes)
}

func TestGetUserInfo(t *testing.T) {
	client, err := steamapi.New(steamapi.NopKey)
	if err != nil {
		t.Error(err)
	}
	info, err := client.WebApiUtil().GetServerInfo()
	if err != nil {
		t.Error(err)
	}
	t.Log(prettyJsonResp(info))
}

func TestGetSupportedAPIList(t *testing.T) {
	key, err := readFileKey()
	if err != nil {
		t.Error(err)
	}
	client, err := steamapi.New(key)
	if err != nil {
		t.Error(err)
	}
	list, err := client.WebApiUtil().GetSupportedAPIList()
	if err != nil {
		t.Error(err)
	}
	t.Log(prettyJsonResp(list))
}
