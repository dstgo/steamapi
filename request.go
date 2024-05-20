package steamapi

import (
	"encoding/json"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/dstgo/steamapi/types/steam"
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

	// InputJson service interface need "input_json" query parameter to pass post params
	InputJson = "input_json"
)

var (
	PartnerEnforceHttpsErr = errors.New("partner.steam-api.com must be requested with https")
)

type Request struct {
	Host      string
	Method    string
	Url       string
	QueryForm map[string]any
	FormData  map[string]any
	Body      any
	Header    http.Header
	err       error

	cfg ClientCfg
	*resty.Request
}

func (r *Request) FullURL() string {
	if r.cfg.https {
		return "https://" + path.Join(r.Host, r.Url)
	}
	return "http://" + path.Join(r.Host, r.Url)
}

func (r *Request) FormEscaped(form map[string]any) map[string]string {
	safeForm := make(map[string]string, len(form))
	for k, v := range form {
		safeForm[k] = url.QueryEscape(cast.ToString(v))
	}
	return safeForm
}

func (r *Request) Attach(req *resty.Request) {
	req.URL = r.FullURL()
	req.Method = r.Method
	req.Body = r.Body
	for k, v := range r.Header {
		req.Header[k] = v
	}
	req.SetQueryParams(r.FormEscaped(r.QueryForm))
	req.SetFormData(r.FormEscaped(r.FormData))
	r.Request = req
}

func (r *Request) Send() (*resty.Response, error) {
	if r.err != nil {
		return nil, r.err
	}
	rawResponse, err := r.Request.Send()
	if err != nil {
		return nil, err
	}
	if err := checkRespStatus(rawResponse); err != nil {
		return nil, errors.Join(err, errors.New(string(rawResponse.Body())))
	}
	return rawResponse, nil
}

// RequestOption
// use options to override some request settings
type RequestOption func(request *Request)

func WithAPIKey(key string) RequestOption {
	return func(request *Request) {
		request.Request.QueryParam.Set(QuerySteamApiKey, key)
	}
}

func WithLanguage(language string) RequestOption {
	return func(request *Request) {
		request.Request.QueryParam.Set(QueryLanguage, language)
	}
}

func WithMethod(method string) RequestOption {
	return func(request *Request) {
		request.Method = method
	}
}

func WithHost(host string) RequestOption {
	return func(request *Request) {
		request.Host = host
	}
}

func WithURL(url string) RequestOption {
	return func(request *Request) {
		request.Url = url
	}
}

func WithRequest(fn func(r *resty.Request)) RequestOption {
	return func(request *Request) {
		fn(request.Request)
	}
}

func WithQueryMap(query map[string]any) RequestOption {
	return func(request *Request) {
		request.QueryForm = query
	}
}

func WithQueryForm(query any) RequestOption {
	return func(request *Request) {
		form, err := toMap(query)
		if err != nil {
			request.err = err
			return
		}
		ok, err := govalidator.ValidateStruct(query)
		if !ok {
			request.err = err
			return
		}
		request.QueryForm = form
	}
}

func WithFormData(formdata any) RequestOption {
	return func(request *Request) {
		formdata, err := toMap(formdata)
		if err != nil {
			request.err = err
			return
		}
		ok, err := govalidator.ValidateStruct(formdata)
		if !ok {
			request.err = err
			return
		}
		request.FormData = formdata
	}
}

func WithInputJson(input any) RequestOption {
	return func(request *Request) {
		ok, err := govalidator.ValidateStruct(input)
		if !ok {
			request.err = err
			return
		}
		inputjson, err := json.Marshal(input)
		if err != nil {
			request.err = err
			return
		}
		request.SetQueryParam(InputJson, string(inputjson))
	}
}

func WithBody(body any) RequestOption {
	return func(request *Request) {
		ok, err := govalidator.ValidateStruct(body)
		if !ok {
			request.err = err
			return
		}
		request.Body = body
	}
}

func WithHeader(header http.Header) RequestOption {
	return func(request *Request) {
		request.Header = header
	}
}

// NewRequest creates a new request, you can use options to customize request configuration
func (c *Client) NewRequest(method, host, url string, options ...RequestOption) *Request {

	r := &Request{
		cfg: c.cfg,
	}

	if err := checkMethod(method); err != nil {
		r.err = err
		return r
	}

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

	// PartnerHost must be request by https
	if r.Host == PartnerHost && !r.cfg.https {
		r.err = PartnerEnforceHttpsErr
		return r
	}

	newReq := c.client.R()

	r.Attach(newReq)

	if !r.Request.QueryParam.Has(QuerySteamApiKey) || len(r.Request.QueryParam.Get(QuerySteamApiKey)) == 0 {
		r.Request.SetQueryParam(QuerySteamApiKey, c.cfg.key)
	}

	if !r.Request.QueryParam.Has(QueryLanguage) || len(r.Request.QueryParam.Get(QueryLanguage)) == 0 {
		r.Request.SetQueryParam(QueryLanguage, r.cfg.language)
	}

	return r
}

// SendRequest this is a helper func to send a request, resultPtr must be a pointer.
func (c *Client) SendRequest(method, host, url string, resultPtr any, opts ...RequestOption) (*resty.Response, error) {
	if c == nil {
		return nil, errors.New("client must be not nil")
	}
	// build new request
	newRequest := c.NewRequest(method, host, url, opts...)
	// bind result
	newRequest.SetResult(&resultPtr)
	// send request
	rawResponse, err := newRequest.Send()
	if err != nil {
		return nil, err
	} else if rawResponse.StatusCode() >= 400 {
		return nil, errors.New(http.StatusText(rawResponse.StatusCode()))
	}
	return rawResponse, nil
}

func (c *Client) Get(host, url string, resultPtr any, opts ...RequestOption) (*resty.Response, error) {
	return c.SendRequest(http.MethodGet, host, url, resultPtr, opts...)
}

func (c *Client) Post(host, url string, resultPtr any, opts ...RequestOption) (*resty.Response, error) {
	return c.SendRequest(http.MethodPost, host, url, resultPtr, opts...)
}

// Unknown some interfaces response are unknown, so return steam.CommonResponse that is alias of map[string]any
func (c *Client) Unknown(method, host, url string, opts ...RequestOption) (steam.CommonResponse, error) {
	commonResponse := steam.CommonResponse{}
	_, err := c.SendRequest(method, host, url, &commonResponse, opts...)
	return commonResponse, err
}
