package httpClient

import (
	"encoding/json"
	"io"
	"net/http"
)

type HttpResponse struct {
	Response *http.Response
}

func NewHttpResponse(res *http.Response) *HttpResponse {
	return &HttpResponse{Response: res}
}

func (response *HttpResponse) Status() string {
	if response.Response == nil {
		return ""
	}

	return response.Response.Status
}

func (response *HttpResponse) StatusCode() int {
	if response.Response == nil {
		return 0
	}

	return response.Response.StatusCode
}

func (response *HttpResponse) Header() http.Header {
	if response.Response == nil {
		return http.Header{}
	}

	return response.Response.Header
}

func (response *HttpResponse) Cookies() []*http.Cookie {
	if response.Response == nil {
		return make([]*http.Cookie, 0)
	}

	return response.Response.Cookies()
}

func (response *HttpResponse) Body() io.ReadCloser {
	if response.Response == nil {
		return nil
	}
	return response.Response.Body
}

func (response *HttpResponse) IsSuccess() bool {
	return response.StatusCode() > 199 && response.StatusCode() < 300
}

func (response *HttpResponse) IsError() bool {
	return response.StatusCode() > 399
}

func (response *HttpResponse) Bind(s interface{}) error {
	defer response.Response.Body.Close()

	return json.NewDecoder(response.Body()).Decode(s)
}

func (response *HttpResponse) ParseBody() (interface{}, error) {
	defer response.Response.Body.Close()

	res, err := io.ReadAll(response.Body())
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
