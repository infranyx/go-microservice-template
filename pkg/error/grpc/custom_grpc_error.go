package grpcError

import (
	"time"

	"google.golang.org/grpc/codes"

	errorConstant "github.com/infranyx/go-microservice-template/pkg/constant/error"
)

func NewGrpcValidationError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errorConstant.ErrValidationFailedTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.InvalidArgument,
		Timestamp: time.Now(),
	}
}

func NewGrpcConflictError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errorConstant.ErrConflictTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.AlreadyExists,
		Timestamp: time.Now(),
	}
}

func NewGrpcBadRequestError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errorConstant.ErrBadRequestTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.InvalidArgument,
		Timestamp: time.Now(),
	}
}

func NewGrpcNotFoundError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errorConstant.ErrNotFoundTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.NotFound,
		Timestamp: time.Now(),
	}
}

func NewGrpcUnAuthorizedError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errorConstant.ErrUnauthorizedTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.Unauthenticated,
		Timestamp: time.Now(),
	}
}

func NewGrpcForbiddenError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errorConstant.ErrForbiddenTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.PermissionDenied,
		Timestamp: time.Now(),
	}
}

func NewGrpcInternalServerError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errorConstant.ErrInternalServerErrorTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.Internal,
		Timestamp: time.Now(),
	}
}

func NewGrpcDomainError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errorConstant.ErrDomainTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.InvalidArgument,
		Timestamp: time.Now(),
	}
}

func NewGrpcApplicationError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errorConstant.ErrApplicationTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.Internal,
		Timestamp: time.Now(),
	}
}

func NewGrpcApiError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errorConstant.ErrApiTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.Internal,
		Timestamp: time.Now(),
	}
}
