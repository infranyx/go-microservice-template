package httpClient

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

const (
	GetMethod     = "GET"
	PostMethod    = "POST"
	PutMethod     = "PUT"
	DeleteMethod  = "DELETE"
	PatchMethod   = "PATCH"
	HeadMethod    = "HEAD"
	OptionsMethod = "OPTIONS"
)

type Config struct{}

func NewHttpClient(config *Config) *http.Client {
	return &http.Client{
		Transport: http.DefaultTransport,
	}
}

type HttpRequest struct {
	request *http.Request
	client  *http.Client
	context context.Context
	URL     string
	Method  string
	Query   url.Values
	Header  http.Header
	Body    io.Reader
}

func BuildReq() *HttpRequest {
	return &HttpRequest{
		client:  NewHttpClient(&Config{}),
		context: context.Background(),
		request: &http.Request{},
		Header:  http.Header{},
		Query:   url.Values{},
		Method:  GetMethod,
		URL:     "",
		Body:    http.NoBody,
	}
}

func (request *HttpRequest) Get(url string) *HttpRequest {
	request.Method = GetMethod
	request.URL = url

	return request
}

func (request *HttpRequest) Head(url string) *HttpRequest {
	request.Method = HeadMethod
	request.URL = url

	return request
}

func (request *HttpRequest) Post(url string) *HttpRequest {
	request.Method = PostMethod
	request.URL = url

	return request
}

func (request *HttpRequest) Put(url string) *HttpRequest {
	request.Method = PutMethod
	request.URL = url

	return request
}

func (request *HttpRequest) Delete(url string) *HttpRequest {
	request.Method = DeleteMethod
	request.URL = url

	return request
}

func (request *HttpRequest) Options(url string) *HttpRequest {
	request.Method = OptionsMethod
	request.URL = url

	return request
}

func (request *HttpRequest) Patch(url string) *HttpRequest {
	request.Method = PatchMethod
	request.URL = url

	return request
}

func (request *HttpRequest) Execute() (*HttpResponse, error) {
	newRequestWithCtx, err := http.NewRequestWithContext(request.context, request.Method, request.URL, request.Body)
	if err != nil {
		return nil, err
	}

	newRequestWithCtx.Header = request.Header
	newRequestWithCtx.URL.RawQuery = request.Query.Encode()
	request.request = newRequestWithCtx

	res, err := request.client.Do(request.request)
	if err != nil {
		return nil, err
	}

	return NewHttpResponse(res), nil
}

func (request *HttpRequest) Client() *http.Client {
	return request.client
}

func (request *HttpRequest) SetClient(config *Config) *HttpRequest {
	request.client = NewHttpClient(config)

	return request
}

func (request *HttpRequest) SetContext(context context.Context) *HttpRequest {
	request.context = context

	return request
}

func (request *HttpRequest) SetHeader(header, value string) *HttpRequest {
	request.Header.Set(header, value)

	return request
}

func (request *HttpRequest) SetHeaders(headers map[string]string) *HttpRequest {
	for h, v := range headers {
		request.SetHeader(h, v)
	}

	return request
}

func (request *HttpRequest) SetQueryParam(param, value string) *HttpRequest {
	query := request.Query
	query.Add(param, value)

	return request
}

func (request *HttpRequest) SetQueryParams(params map[string]string) *HttpRequest {
	query := request.Query
	for p, v := range params {
		query.Add(p, v)
	}

	return request
}

func (request *HttpRequest) SetBody(body io.Reader) *HttpRequest {
	rc := io.NopCloser(body)
	request.Body = rc

	return request
}
