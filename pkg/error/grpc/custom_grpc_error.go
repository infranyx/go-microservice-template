package grpcError

import (
	"time"

	errConst "github.com/infranyx/go-grpc-template/pkg/constant/error"
	"google.golang.org/grpc/codes"
)

func NewGrpcValidationError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errConst.ErrValidationFailedTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.InvalidArgument,
		Timestamp: time.Now(),
	}
}

func NewGrpcConflictError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errConst.ErrConflictTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.AlreadyExists,
		Timestamp: time.Now(),
	}
}

func NewGrpcBadRequestError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errConst.ErrBadRequestTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.InvalidArgument,
		Timestamp: time.Now(),
	}
}

func NewGrpcNotFoundError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errConst.ErrNotFoundTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.NotFound,
		Timestamp: time.Now(),
	}
}

func NewGrpcUnAuthorizedError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errConst.ErrUnauthorizedTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.Unauthenticated,
		Timestamp: time.Now(),
	}
}

func NewGrpcForbiddenError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errConst.ErrForbiddenTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.PermissionDenied,
		Timestamp: time.Now(),
	}
}

func NewGrpcInternalServerError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errConst.ErrInternalServerErrorTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.Internal,
		Timestamp: time.Now(),
	}
}

func NewGrpcDomainError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errConst.ErrDomainTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.InvalidArgument,
		Timestamp: time.Now(),
	}
}

func NewGrpcApplicationError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errConst.ErrApplicationTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.Internal,
		Timestamp: time.Now(),
	}
}

func NewGrpcApiError(code int, message string, details map[string]string) GrpcErr {
	return &grpcErr{
		Title:     errConst.ErrApiTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    codes.Internal,
		Timestamp: time.Now(),
	}
}
