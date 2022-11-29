package grpc_errors

import (
	"encoding/json"
	"fmt"
	"time"

	customErrors "github.com/infranyx/go-grpc-template/shared/error/custom_error"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
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
	// x := status.Error(p.GetStatus(), p.ToJson())

	st := status.New(codes.Unknown, "unknown error occurred")
	br := &errdetails.DebugInfo{
		Detail: "detail reason of err",
	}

	// byts, _ := proto.Marshal(br)
	// if err != nil {
	// 	return nil, err
	// }
	stWithDetails, _ := st.WithDetails(br)
	// if err != nil {
	// 	return nil, st.Err()
	// }
	fmt.Println(stWithDetails.Err())
	return stWithDetails.Err()
}

func (p *grpcErr) ToJson() string {
	return string(p.json())
}

func (p *grpcErr) json() []byte {
	b, _ := json.Marshal(&p)
	return b
}
