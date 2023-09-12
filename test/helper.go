package test

import (
	"encoding/json"
	"os"
)

func readFileKey(file string) (string, error) {
	bytes, err := os.ReadFile(file)
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
