package grpcError

import (
	"time"

	errConst "github.com/infranyx/go-grpc-template/pkg/constant/error"
	"google.golang.org/grpc/codes"
)

func NewGrpcValidationError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
	validationError :=
		&grpcErr{
			Title:      errConst.ErrBadRequestTitle,
			Code:       code,
			Msg:        message,
			Details:    details,
			Status:     codes.InvalidArgument,
			Timestamp:  time.Now(),
			StackTrace: stackTrace,
		}

	return validationError
}

func NewGrpcConflictError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errConst.ErrConflictTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.AlreadyExists,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewGrpcBadRequestError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errConst.ErrBadRequestTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.InvalidArgument,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewGrpcNotFoundError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errConst.ErrNotFoundTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.NotFound,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewGrpcUnAuthorizedError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errConst.ErrUnauthorizedTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.Unauthenticated,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewGrpcForbiddenError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errConst.ErrForbiddenTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.PermissionDenied,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewGrpcInternalServerError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errConst.ErrInternalServerErrorTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.Internal,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewGrpcDomainError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errConst.ErrDomainTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.InvalidArgument,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewGrpcApplicationError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errConst.ErrApplicationTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.Internal,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewGrpcApiError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errConst.ErrApiTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.Internal,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}
