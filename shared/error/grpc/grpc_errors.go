package grpc_errors

import (
	"encoding/json"
	"fmt"
	"time"

	customErrors "github.com/infranyx/go-grpc-template/shared/error/custom_error"
	sharedBuf "github.com/infranyx/protobuf-template-go/shared/error"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcErr struct {
	Status     codes.Code                 `json:"status,omitempty"`
	Code       int                        `json:"code,omitempty"`
	Title      string                     `json:"title,omitempty"`
	Msg        string                     `json:"msg,omitempty"`
	Details    []customErrors.ErrorDetail `json:"errorDetail,omitempty"`
	Timestamp  time.Time                  `json:"timestamp,omitempty"`
	StackTrace string                     `json:"stackTrace,omitempty"`
}

type GrpcErr interface {
	GetStatus() codes.Code
	SetStatus(status codes.Code) GrpcErr
	GetCode() int
	SetCode(code int) GrpcErr
	GetTitle() string
	SetTitle(title string) GrpcErr
	GetStackTrace() string
	SetStackTrace(stackTrace string) GrpcErr
	GetMsg() string
	SetMsg(msg string) GrpcErr
	GetDetails() []customErrors.ErrorDetail
	SetDetails(details []customErrors.ErrorDetail) GrpcErr
	GetTimestamp() time.Time
	SetTimestamp(time time.Time) GrpcErr
	Error() string
	ErrBody() error
	ToJson() string
	ToGrpcResponseErr() error
}

func NewGrpcError(status codes.Code, code int, title string, message string, details []customErrors.ErrorDetail, stackTrace string) GrpcErr {
	grpcErr := &grpcErr{
		Status:     status,
		Code:       code,
		Title:      title,
		Msg:        message,
		Details:    details,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}

	return grpcErr
}

// ErrBody Error body
func (p *grpcErr) ErrBody() error {
	return p
}

// Error  Error() interface method
func (p *grpcErr) Error() string {
	return fmt.Sprintf("Error Title: %s - Error Status: %d - Error Detail: %s", p.Title, p.Status, p.Msg)
}

func (p *grpcErr) GetStatus() codes.Code {
	return p.Status
}

func (p *grpcErr) SetStatus(status codes.Code) GrpcErr {
	p.Status = status

	return p
}

func (p *grpcErr) GetCode() int {
	return p.Code
}

func (p *grpcErr) SetCode(code int) GrpcErr {
	p.Code = code

	return p
}

func (p *grpcErr) GetTitle() string {
	return p.Title
}

func (p *grpcErr) SetTitle(title string) GrpcErr {
	p.Title = title

	return p
}

func (p *grpcErr) GetMsg() string {
	return p.Msg
}

func (p *grpcErr) SetMsg(message string) GrpcErr {
	p.Msg = message

	return p
}

func (p *grpcErr) GetDetails() []customErrors.ErrorDetail {
	return p.Details
}

func (p *grpcErr) SetDetails(detail []customErrors.ErrorDetail) GrpcErr {
	p.Details = detail

	return p
}

func (p *grpcErr) GetTimestamp() time.Time {
	return p.Timestamp
}

func (p *grpcErr) SetTimestamp(time time.Time) GrpcErr {
	p.Timestamp = time

	return p
}

func (p *grpcErr) GetStackTrace() string {
	return p.StackTrace
}

func (p *grpcErr) SetStackTrace(stackTrace string) GrpcErr {
	p.StackTrace = stackTrace

	return p
}

// ToGrpcResponseErr creates a gRPC error response to send grpc engine
func (p *grpcErr) ToGrpcResponseErr() error {
	st := status.New(codes.Code(p.Code), p.Error())
	mappedErr := &sharedBuf.CustomError{
		Title:      p.Title,
		Code:       "1",
		Detail:     p.Msg,
		Timestamp:  p.Timestamp.Format(time.RFC3339),
		StackTrace: &p.StackTrace,
	}
	// byts, _ := proto.Marshal(mappedErr)
	stWithDetails, _ := st.WithDetails(mappedErr)
	return stWithDetails.Err()
}

func (p *grpcErr) ToJson() string {
	return string(p.json())
}

func (p *grpcErr) json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

func ParseExternalGrpcErr(err error) GrpcErr {
	st := status.Convert(err)
	for _, detail := range st.Details() {
		switch t := detail.(type) {
		case *sharedBuf.CustomError:
			timestamp, _ := time.Parse(time.RFC3339, t.Timestamp)
			return &grpcErr{
				Status: st.Code(),
				Code:   1,
				Title:  t.Title,
				// Msg:        message,
				// Details:    details,
				Timestamp: timestamp,
				// StackTrace: stackTrace,
			}
		}
	}
	return nil
}
