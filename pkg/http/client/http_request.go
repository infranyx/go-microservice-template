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

type HttpClientConfig struct {
}

func NewHttpClient(config *HttpClientConfig) *http.Client {
	client := &http.Client{
		Transport: http.DefaultTransport,
	}
	return client
}

type HttpRequest struct {
	req     *http.Request
	client  *http.Client
	context context.Context

	URL    string
	Method string
	Query  url.Values
	Header http.Header
	Body   io.Reader
}

func BuildReq() *HttpRequest {
	return &HttpRequest{
		client:  NewHttpClient(&HttpClientConfig{}),
		context: context.Background(),
		req:     &http.Request{},
		Header:  http.Header{},
		Query:   url.Values{},
		Method:  GetMethod,
		URL:     "",
		Body:    http.NoBody,
	}
}

func (hr *HttpRequest) Get(url string) *HttpRequest {
	hr.Method = GetMethod
	hr.URL = url
	return hr
}

func (hr *HttpRequest) Head(url string) *HttpRequest {
	hr.Method = HeadMethod
	hr.URL = url
	return hr
}

func (hr *HttpRequest) Post(url string) *HttpRequest {
	hr.Method = PostMethod
	hr.URL = url
	return hr
}

func (hr *HttpRequest) Put(url string) *HttpRequest {
	hr.Method = PutMethod
	hr.URL = url
	return hr
}

func (hr *HttpRequest) Delete(url string) *HttpRequest {
	hr.Method = DeleteMethod
	hr.URL = url
	return hr
}

func (hr *HttpRequest) Options(url string) *HttpRequest {
	hr.Method = OptionsMethod
	hr.URL = url
	return hr
}

func (hr *HttpRequest) Patch(url string) *HttpRequest {
	hr.Method = PatchMethod
	hr.URL = url
	return hr
}

func (hr *HttpRequest) Execute() (*HttpResponse, error) {
	nr, err := http.NewRequestWithContext(hr.context, hr.Method, hr.URL, hr.Body)
	if err != nil {
		return nil, err
	}

	nr.Header = hr.Header
	nr.URL.RawQuery = hr.Query.Encode()
	hr.req = nr

	res, err := hr.client.Do(hr.req)
	if err != nil {
		return nil, err
	}
	return NewHttpResponse(res), nil
}

func (hr *HttpRequest) Client() *http.Client {
	return hr.client
}

func (hr *HttpRequest) SetClient(config *HttpClientConfig) *HttpRequest {
	hr.client = NewHttpClient(config)
	return hr
}

func (hr *HttpRequest) SetContext(context context.Context) *HttpRequest {
	hr.context = context
	return hr
}

func (hr *HttpRequest) SetHeader(header, value string) *HttpRequest {
	hr.Header.Set(header, value)
	return hr
}

func (hr *HttpRequest) SetHeaders(headers map[string]string) *HttpRequest {
	for h, v := range headers {
		hr.SetHeader(h, v)
	}
	return hr
}

func (hr *HttpRequest) SetQueryParam(param, value string) *HttpRequest {
	query := hr.Query
	query.Add(param, value)
	return hr
}

func (hr *HttpRequest) SetQueryParams(params map[string]string) *HttpRequest {
	query := hr.Query
	for p, v := range params {
		query.Add(p, v)
	}
	return hr
}

func (hr *HttpRequest) SetBody(body io.Reader) *HttpRequest {
	rc := io.NopCloser(body)
	hr.Body = rc
	return hr
}
