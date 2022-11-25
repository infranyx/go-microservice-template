package grpc_errors

import (
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcErr struct {
	Status     codes.Code `json:"status,omitempty"`
	Code       int        `json:"code,omitempty"`
	Title      string     `json:"title,omitempty"`
	Detail     string     `json:"detail,omitempty"`
	Timestamp  time.Time  `json:"timestamp,omitempty"`
	StackTrace string     `json:"stackTrace,omitempty"`
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
	GetDetail() string
	SetDetail(detail string) GrpcErr
	GetTimestamp() time.Time
	SetTimestamp(time time.Time) GrpcErr
	Error() string
	ErrBody() error
	ToJson() string
	ToGrpcResponseErr() error
}

func NewGrpcError(status codes.Code, code int, title string, detail string, stackTrace string) GrpcErr {
	grpcErr := &grpcErr{
		Status:     status,
		Code:       code,
		Title:      title,
		Timestamp:  time.Now(),
		Detail:     detail,
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
	return fmt.Sprintf("Error Title: %s - Error Status: %d - Error Detail: %s", p.Title, p.Status, p.Detail)
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

func (p *grpcErr) GetDetail() string {
	return p.Detail
}

func (p *grpcErr) SetDetail(detail string) GrpcErr {
	p.Detail = detail

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
	return status.Error(p.GetStatus(), p.ToJson())
}

func (p *grpcErr) ToJson() string {
	return string(p.json())
}

func (p *grpcErr) json() []byte {
	b, _ := json.Marshal(&p)
	return b
}
