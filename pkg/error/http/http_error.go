package httpError

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type httpErr struct {
	Status    int               `json:"status,omitempty"`
	Code      int               `json:"code,omitempty"`
	Title     string            `json:"title,omitempty"`
	Msg       string            `json:"msg,omitempty"`
	Details   map[string]string `json:"errorDetail,omitempty"`
	Timestamp time.Time         `json:"timestamp,omitempty"`
}

type HttpErr interface {
	GetStatus() int
	SetStatus(status int) HttpErr
	GetCode() int
	SetCode(code int) HttpErr
	GetTitle() string
	SetTitle(title string) HttpErr
	GetMsg() string
	SetMsg(msg string) HttpErr
	GetDetails() map[string]string
	SetDetails(details map[string]string) HttpErr
	GetTimestamp() time.Time
	SetTimestamp(time time.Time) HttpErr
	Error() string
	ErrBody() error
	WriteTo(w http.ResponseWriter) (int, error)
	// ToGrpcResponseErr() error
}

func NewHttpError(status int, code int, title string, messahe string, details map[string]string) HttpErr {
	httpErr := &httpErr{
		Status:    status,
		Code:      code,
		Title:     title,
		Msg:       messahe,
		Details:   details,
		Timestamp: time.Now(),
	}

	return httpErr
}

func (he *httpErr) ErrBody() error {
	return he
}

func (he *httpErr) Error() string {
	return he.Msg
}

func (he *httpErr) GetStatus() int {
	return he.Status
}

func (he *httpErr) SetStatus(status int) HttpErr {
	he.Status = status

	return he
}

func (he *httpErr) GetCode() int {
	return he.Code
}

func (he *httpErr) SetCode(code int) HttpErr {
	he.Code = code

	return he
}

func (he *httpErr) GetTitle() string {
	return he.Title
}

func (he *httpErr) SetTitle(title string) HttpErr {
	he.Title = title

	return he
}

func (he *httpErr) GetMsg() string {
	return he.Msg
}

func (he *httpErr) SetMsg(messahe string) HttpErr {
	he.Msg = messahe

	return he
}

func (he *httpErr) GetDetails() map[string]string {
	return he.Details
}

func (he *httpErr) SetDetails(detail map[string]string) HttpErr {
	he.Details = detail

	return he
}

func (he *httpErr) GetTimestamp() time.Time {
	return he.Timestamp
}

func (he *httpErr) SetTimestamp(time time.Time) HttpErr {
	he.Timestamp = time

	return he
}

func IsHttpError(err error) bool {
	var httpErr HttpErr
	return errors.As(err, &httpErr)
}

func AsHttpError(err error) HttpErr {
	var httpErr HttpErr
	if errors.As(err, &httpErr) {
		return httpErr
	}
	return nil
}

const (
	ContentTypeJSON = "application/problem+json"
)

// WriteTo writes the JSON Problem to an HTTP Response Writer
func (he *httpErr) WriteTo(w http.ResponseWriter) (int, error) {
	status := he.GetStatus()
	if status == 0 {
		status = http.StatusInternalServerError
	}
	w.Header().Set("Content-Type", ContentTypeJSON)
	w.WriteHeader(status)
	return w.Write(he.json())
}

func (he *httpErr) json() []byte {
	b, _ := json.Marshal(&he)
	return b
}

// // ToGrpcResponseErr creates a gRPC error response to send grpc engine
// func (he *httpErr) ToGrpcResponseErr() error {
// 	st := status.New(he.Status, he.Error())
// 	mappedErr := &sharedBuf.CustomError{
// 		Title:     he.Title,
// 		Code:      int64(he.Code),
// 		Msg:       he.Msg,
// 		Details:   he.Details,
// 		Timestamp: he.Timestamp.Format(time.RFC3339),
// 	}
// 	stWithDetails, _ := st.WithDetails(mappedErr)
// 	return stWithDetails.Err()
// }

// func ParseExternalHttpErr(err error) HttpErr {
// 	st := status.Convert(err)
// 	for _, detail := ranhe st.Details() {
// 		switch t := detail.(type) {
// 		case *sharedBuf.CustomError:
// 			timestamp, _ := time.Parse(time.RFC3339, t.Timestamp)
// 			return &httpErr{
// 				Status:    st.Code(),
// 				Code:      int(t.Code),
// 				Title:     t.Title,
// 				Msg:       t.Msg,
// 				Details:   t.Details,
// 				Timestamp: timestamp,
// 			}
// 		}
// 	}
// 	return nil
// }
