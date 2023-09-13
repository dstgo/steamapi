package steamapi

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

var (
	ApiKeyNotExistErr = errors.New("steam api key must be provided")
)

const (
	NopKey           = "nop-key"
	EmptyKey         = ""
	QuerySteamApiKey = "key"
	QueryLanguage    = "language"
)

type ClientOption func(client *Client)

func WithClientKey(key string) ClientOption {
	return func(c *Client) {
		c.cfg.key = key
	}
}

func WithClientResty(client *resty.Client) ClientOption {
	return func(c *Client) {
		c.client = client
	}
}

func WithHttpsClient() ClientOption {
	return func(c *Client) {
		c.cfg.https = true
	}
}

// Client
// steam api client, interact with steam api server
type Client struct {
	cfg    ClientCfg
	client *resty.Client
}

type ClientCfg struct {
	key      string
	https    bool
	language string
}

// New func create a new Client only with api key which must be provided,
// if you just want to access some interface which does not need api key
// you can pass the NopKey
func New(key string) (*Client, error) {
	return NewWith(
		WithClientKey(key),
	)
}

func NewWith(options ...ClientOption) (*Client, error) {

	client := new(Client)

	// apply options
	for _, option := range options {
		option(client)
	}

	// key must be provided
	if len(client.cfg.key) == 0 {
		return client, ApiKeyNotExistErr
	}

	if client.client == nil {
		client.client = resty.New()
	}

	return client, nil
}
