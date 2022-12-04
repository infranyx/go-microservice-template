package grpcError

import (
	errorTitles "github.com/infranyx/go-grpc-template/pkg/constant/error"
	"time"

	customErrors "github.com/infranyx/go-grpc-template/shared/error/custom_error"
	"google.golang.org/grpc/codes"
)

func NewValidationGrpcError(code int, message string, details []customErrors.ErrorDetail, stackTrace string) GrpcErr {
	validationError :=
		&grpcErr{
			Title:      errorTitles.ErrBadRequestTitle,
			Code:       code,
			Msg:        message,
			Details:    details,
			Status:     codes.InvalidArgument,
			Timestamp:  time.Now(),
			StackTrace: stackTrace,
		}

	return validationError
}

func NewConflictGrpcError(code int, message string, details []customErrors.ErrorDetail, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrConflictTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.AlreadyExists,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewBadRequestGrpcError(code int, message string, details []customErrors.ErrorDetail, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrBadRequestTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.InvalidArgument,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewNotFoundErrorGrpcError(code int, message string, details []customErrors.ErrorDetail, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrNotFoundTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.NotFound,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewUnAuthorizedErrorGrpcError(code int, message string, details []customErrors.ErrorDetail, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrUnauthorizedTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.Unauthenticated,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewForbiddenGrpcError(code int, message string, details []customErrors.ErrorDetail, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrForbiddenTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.PermissionDenied,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewInternalServerGrpcError(code int, message string, details []customErrors.ErrorDetail, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrInternalServerErrorTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.Internal,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewDomainGrpcError(code int, message string, details []customErrors.ErrorDetail, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrDomainTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.InvalidArgument,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewApplicationGrpcError(code int, message string, details []customErrors.ErrorDetail, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrApplicationTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.Internal,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewApiGrpcError(code int, message string, details []customErrors.ErrorDetail, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrApiTitle,
		Code:       code,
		Msg:        message,
		Details:    details,
		Status:     codes.Internal,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}
