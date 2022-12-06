package grpcError

import (
	"time"

	errConst "github.com/infranyx/go-grpc-template/pkg/constant/error"
	"google.golang.org/grpc/codes"
)

func NewValidationGrpcError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
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

func NewConflictGrpcError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
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

func NewBadRequestGrpcError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
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

func NewNotFoundErrorGrpcError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
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

func NewUnAuthorizedErrorGrpcError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
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

func NewForbiddenGrpcError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
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

func NewInternalServerGrpcError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
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

func NewDomainGrpcError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
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

func NewApplicationGrpcError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
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

func NewApiGrpcError(code int, message string, details map[string]string, stackTrace string) GrpcErr {
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
