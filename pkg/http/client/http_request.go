package httpClient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
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
}

func BuildReq() *HttpRequest {
	return &HttpRequest{
		client:  NewHttpClient(&HttpClientConfig{}),
		context: context.Background(),
		req:     &http.Request{},
	}
}

func (hr *HttpRequest) Get(url string) (*HttpRequest, error) {
	return hr.handleReq(GetMethod, url)
}

func (hr *HttpRequest) Head(url string) (*HttpRequest, error) {
	return hr.handleReq(HeadMethod, url)
}

func (hr *HttpRequest) Post(url string) (*HttpRequest, error) {
	return hr.handleReq(PostMethod, url)
}

func (hr *HttpRequest) Put(url string) (*HttpRequest, error) {
	return hr.handleReq(PutMethod, url)
}

func (hr *HttpRequest) Delete(url string) (*HttpRequest, error) {
	return hr.handleReq(DeleteMethod, url)
}

func (hr *HttpRequest) Options(url string) (*HttpRequest, error) {
	return hr.handleReq(OptionsMethod, url)
}

func (hr *HttpRequest) Patch(url string) (*HttpRequest, error) {
	return hr.handleReq(PatchMethod, url)
}

func (hr *HttpRequest) handleReq(method string, url string) (*HttpRequest, error) {
	nr, err := http.NewRequestWithContext(hr.context, method, url, hr.req.Body)
	hr.req = nr
	return hr, err
}

func (hr *HttpRequest) Execute() (*HttpResponse, error) {
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
	hr.req.Header.Set(header, value)
	return hr
}

func (hr *HttpRequest) SetHeaders(headers map[string]string) *HttpRequest {
	for h, v := range headers {
		hr.SetHeader(h, v)
	}
	return hr
}

func (hr *HttpRequest) SetQueryParam(param, value string) *HttpRequest {
	query := hr.req.URL.Query()
	query.Add(param, value)
	hr.req.URL.RawQuery = query.Encode()
	return hr
}

func (hr *HttpRequest) SetQueryParams(params map[string]string) *HttpRequest {
	query := hr.req.URL.Query()
	for p, v := range params {
		query.Add(p, v)
	}
	hr.req.URL.RawQuery = query.Encode()
	return hr
}

func (hr *HttpRequest) SetBody(body interface{}) (*HttpRequest, error) {
	mbody, err := json.Marshal(body)
	if err != nil {
		return hr, err
	}

	reader := bytes.NewReader(mbody)
	rc := io.NopCloser(reader)

	hr.req.Body = rc
	return hr, err
}

func (hr *HttpRequest) SetCookie(hc *http.Cookie) *HttpRequest {
	hr.req.AddCookie(hc)
	return hr
}

func (hr *HttpRequest) SetCookies(hcs []*http.Cookie) *HttpRequest {
	for _, hc := range hcs {
		hr.SetCookie(hc)
	}
	return hr
}
