package grpcErrors

import (
	"time"

	errorTitles "github.com/infranyx/go-grpc-template/constant/errors"
	"google.golang.org/grpc/codes"
)

func NewValidationGrpcError(code int, detail string, stackTrace string) GrpcErr {
	validationError :=
		&grpcErr{
			Title:      errorTitles.ErrBadRequestTitle,
			Code:       code,
			Detail:     detail,
			Status:     codes.InvalidArgument,
			Timestamp:  time.Now(),
			StackTrace: stackTrace,
		}

	return validationError
}

func NewConflictGrpcError(code int, detail string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrConflictTitle,
		Code:       code,
		Detail:     detail,
		Status:     codes.AlreadyExists,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewBadRequestGrpcError(code int, detail string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrBadRequestTitle,
		Code:       code,
		Detail:     detail,
		Status:     codes.InvalidArgument,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewNotFoundErrorGrpcError(code int, detail string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrNotFoundTitle,
		Code:       code,
		Detail:     detail,
		Status:     codes.NotFound,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewUnAuthorizedErrorGrpcError(code int, detail string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrUnauthorizedTitle,
		Code:       code,
		Detail:     detail,
		Status:     codes.Unauthenticated,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewForbiddenGrpcError(code int, detail string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrForbiddenTitle,
		Code:       code,
		Detail:     detail,
		Status:     codes.PermissionDenied,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewInternalServerGrpcError(code int, detail string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrInternalServerErrorTitle,
		Code:       code,
		Detail:     detail,
		Status:     codes.Internal,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewDomainGrpcError(status codes.Code, code int, detail string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrDomainTitle,
		Code:       code,
		Detail:     detail,
		Status:     status,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewApplicationGrpcError(status codes.Code, code int, detail string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrApplicationTitle,
		Code:       code,
		Detail:     detail,
		Status:     status,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}

func NewApiGrpcError(status codes.Code, code int, detail string, stackTrace string) GrpcErr {
	return &grpcErr{
		Title:      errorTitles.ErrApiTitle,
		Code:       code,
		Detail:     detail,
		Status:     status,
		Timestamp:  time.Now(),
		StackTrace: stackTrace,
	}
}
