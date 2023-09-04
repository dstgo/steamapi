package steamapi

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"path"
	"slices"
)

const (
	PublicHost = "api.steampowered.com"

	PartnerHost = "partner.steam-api.com"
)

type Request struct {
	Https  bool
	Host   string
	Method string
	Url    string
	lang   string
	*resty.Request
}

func (r *Request) FullURL() string {
	if r.Https {
		return "https://" + path.Join(r.Host, r.Url)
	}
	return "http://" + path.Join(r.Host, r.Url)
}

func (r *Request) Attach(req *resty.Request) {
	req.URL = r.FullURL()
	req.Method = r.Method
	r.Request = req
}

func (r *Request) Send() (*resty.Response, error) {
	rawResponse, err := r.Request.Send()
	if err != nil {
		return nil, err
	}
	if err := requireOk(rawResponse); err != nil {
		return nil, err
	}
	return rawResponse, nil
}

// RequestOptions
// use options to override some request settings
type RequestOptions func(request *Request)

func WithRequestMethod(method string) RequestOptions {
	return func(request *Request) {
		request.Method = method
	}
}

func WithRequestHost(host string) RequestOptions {
	return func(request *Request) {
		request.Host = host
	}
}

func WithRequestHttps(https bool) RequestOptions {
	return func(request *Request) {
		request.Https = https
	}
}

func WithRequestURL(url string) RequestOptions {
	return func(request *Request) {
		request.Url = url
	}
}

func WithRequestFn(fn func(r *resty.Request)) RequestOptions {
	return func(request *Request) {
		fn(request.Request)
	}
}

// NewRequest creates a new request, you can use options to customize request configuration
func (c *Client) NewRequest(method, host, url string, options ...RequestOptions) *Request {

	r := &Request{}

	// apply options
	for _, apply := range options {
		apply(r)
	}

	if len(r.Host) == 0 {
		r.Host = host
	}

	if len(r.Url) == 0 {
		r.Url = url
	}

	if len(r.Method) == 0 {
		r.Method = method
	}

	if len(r.lang) == 0 {
		r.lang = c.language
	}

	newReq := c.client.R().SetQueryParams(map[string]string{
		QuerySteamApiKey: c.key,
		QueryLanguage:    r.lang,
	})

	r.Attach(newReq)

	return r
}

func (c *Client) Get(host, url string, options ...RequestOptions) *Request {
	return c.NewRequest(http.MethodGet, host, url, options...)
}

func (c *Client) Post(host, url string, options ...RequestOptions) *Request {
	return c.NewRequest(http.MethodPost, host, url, options...)
}

func (c *Client) Delete(host, url string, options ...RequestOptions) *Request {
	return c.NewRequest(http.MethodDelete, host, url, options...)
}

func (c *Client) Put(host, url string, options ...RequestOptions) *Request {
	return c.NewRequest(http.MethodPut, host, url, options...)
}

func (c *Client) Options(host, url string, options ...RequestOptions) *Request {
	return c.NewRequest(http.MethodOptions, host, url, options...)
}

func requireOk(resp *resty.Response) error {
	return requireHttpCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted)
}

func requireHttpCode(resp *resty.Response, httpCode ...int) error {
	if resp == nil {
		return errors.New("nil response")
	}

	if !slices.Contains(httpCode, resp.StatusCode()) {
		return fmt.Errorf("unexpected status code: %d, expected %v", resp.StatusCode(), httpCode)
	}
	return nil
}
