package httpClient

import (
	"encoding/json"
	"io"
	"net/http"
)

type HttpResponse struct {
	Res *http.Response
}

func NewHttpResponse(res *http.Response) *HttpResponse {
	return &HttpResponse{Res: res}
}

func (hr *HttpResponse) Status() string {
	if hr.Res == nil {
		return ""
	}
	return hr.Res.Status
}

func (hr *HttpResponse) StatusCode() int {
	if hr.Res == nil {
		return 0
	}
	return hr.Res.StatusCode
}

func (hr *HttpResponse) Header() http.Header {
	if hr.Res == nil {
		return http.Header{}
	}
	return hr.Res.Header
}

func (hr *HttpResponse) Cookies() []*http.Cookie {
	if hr.Res == nil {
		return make([]*http.Cookie, 0)
	}
	return hr.Res.Cookies()
}

func (hr *HttpResponse) Body() io.ReadCloser {
	if hr.Res == nil {
		return nil
	}
	return hr.Res.Body
}

func (hr *HttpResponse) IsSuccess() bool {
	return hr.StatusCode() > 199 && hr.StatusCode() < 300
}

func (hr *HttpResponse) IsError() bool {
	return hr.StatusCode() > 399
}

func (hr *HttpResponse) Bind(s interface{}) error {
	defer hr.Res.Body.Close()
	return json.NewDecoder(hr.Body()).Decode(s)
}

func (hr *HttpResponse) ParseBody() (interface{}, error) {
	defer hr.Res.Body.Close()
	var data interface{}
	res, err := io.ReadAll(hr.Body())
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
