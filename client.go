package steamapi

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
)

var (
	ApiKeyNotExistErr = errors.New("steam api key must be provided")
)

const (
	NopKey           = "nop-key"
	QuerySteamApiKey = "key"
	QueryLanguage    = "language"
)

type Options func(client *Client)

func WithKey(key string) Options {
	return func(c *Client) {
		c.key = key
	}
}

func WithLanguage(lang uint) Options {
	return func(c *Client) {
		c.language = cast.ToString(lang)
	}
}

func WithResty(client *resty.Client) Options {
	return func(c *Client) {
		c.client = client
	}
}

func WithHttps(https bool) Options {
	return func(c *Client) {
		c.https = https
	}
}

// Client
// steam api client, interact with steam api server
type Client struct {
	key      string
	language string
	https    bool
	client   *resty.Client

	baseurl string
}

// New func create a new Client only with api key which must be provided,
// if you just want to access some interface which does not need api key
// you can pass the NopKey
func New(key string) (*Client, error) {
	return NewWith(
		WithKey(key),
	)
}

func NewWith(options ...Options) (*Client, error) {

	client := new(Client)

	// apply options
	for _, option := range options {
		option(client)
	}

	// key must be provided
	if len(client.key) == 0 {
		return client, ApiKeyNotExistErr
	}

	if len(client.language) == 0 {
		client.language = "0"
	}

	if client.client == nil {
		client.client = resty.New()
	}

	return client, nil
}
