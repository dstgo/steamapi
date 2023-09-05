package steamapi

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
	"net/http"
	"net/url"
	"path"
)

const (
	// PublicHost is the public steam server, you can interact with you own normal key
	PublicHost = "api.steampowered.com"

	// PartnerHost is the partner steam server, you should take your publisher api key in query parameters in any case
	PartnerHost = "partner.steam-api.com"
)

type Request struct {
	Https     bool
	Host      string
	Method    string
	Url       string
	QueryForm map[string]any
	Body      any
	Header    http.Header

	*resty.Request
}

func (r *Request) FullURL() string {
	if r.Https {
		return "https://" + path.Join(r.Host, r.Url)
	}
	return "http://" + path.Join(r.Host, r.Url)
}

func (r *Request) QueryFormEscaped() map[string]string {
	safeForm := make(map[string]string, len(r.QueryForm))
	for k, v := range r.QueryForm {
		safeForm[k] = url.QueryEscape(cast.ToString(v))
	}
	return safeForm
}

func (r *Request) Attach(req *resty.Request) {
	req.URL = r.FullURL()
	req.Method = r.Method
	req.Body = r.Body
	req.Header = r.Header
	req.SetQueryParams(r.QueryFormEscaped())
	r.Request = req
}

func (r *Request) Send() (*resty.Response, error) {
	rawResponse, err := r.Request.Send()
	if err != nil {
		return nil, err
	}
	if err := checkRespStatus(rawResponse); err != nil {
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

func WithRequestQuery(query map[string]any) RequestOptions {
	return func(request *Request) {
		request.QueryForm = query
	}
}

func WithRequestBody(body any) RequestOptions {
	return func(request *Request) {
		request.Body = body
	}
}

func WithRequestHeaders(header http.Header) RequestOptions {
	return func(request *Request) {
		request.Header = header
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

	newReq := c.client.R()

	r.Attach(newReq)

	if !r.Request.QueryParam.Has(QuerySteamApiKey) || len(r.Request.QueryParam.Get(QuerySteamApiKey)) == 0 {
		r.Request.SetQueryParam(QuerySteamApiKey, c.key)
	}

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
