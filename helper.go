package steamapi

import (
	"errors"
	"github.com/dstgo/steamapi/types/steam"
	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func checkRespStatus(resp *resty.Response) error {
	if resp == nil {
		return errors.New("nil response")
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	switch resp.StatusCode() {
	case 400:
		return steam.ErrBadRequest
	case 401:
		return steam.ErrUnauthorized
	case 403:
		return steam.ErrForbidden
	case 404:
		return steam.ErrNotFound
	case 405:
		return steam.ErrNoMethod
	case 429:
		return steam.ErrExcessive
	case 500:
		return steam.ErrInternal
	case 503:
		return steam.ErrUnavailable
	}

	return nil
}

func checkMethod(method string) error {
	switch method {
	case http.MethodPost:
		return nil
	case http.MethodGet:
		return nil
	}
	return errors.New("method not supported")
}

func toMap(in any) (map[string]any, error) {
	out := make(map[string]any, 10)
	if err := mapstructure.Decode(in, &out); err != nil {
		return out, err
	}
	return out, nil
}

func joinRequestOptions(options []RequestOption, ops ...RequestOption) []RequestOption {
	return append(ops, options...)
}
