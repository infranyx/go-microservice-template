package grpcError

import (
	"github.com/infranyx/go-grpc-template/pkg/error/custom_error"
	errorUtils "github.com/infranyx/go-grpc-template/pkg/error/error_utils"
	"google.golang.org/grpc/codes"
)

func ParseError(err error) GrpcErr {
	customErr := customError.GetCustomError(err)
	stackTrace := errorUtils.RootStackTrace(err)

	if customErr == nil {
		err = customError.NewApiErrorWrap(err, "Unkown Error", 0, nil)
		customErr = customError.GetCustomError(err)
		stackTrace = errorUtils.RootStackTrace(err)
		return NewGrpcError(codes.Internal, customErr.Code(), codes.Internal.String(), customErr.Error(), customErr.Details(), stackTrace)
	}

	if err != nil {
		switch {
		case customError.IsDomainError(err):
			return NewDomainGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customError.IsApplicationError(err):
			return NewApplicationGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customError.IsApiError(err):
			return NewApiGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customError.IsBadRequestError(err):
			return NewBadRequestGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customError.IsNotFoundError(err):
			return NewNotFoundErrorGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customError.IsValidationError(err):
			return NewValidationGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customError.IsUnAuthorizedError(err):
			return NewUnAuthorizedErrorGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customError.IsForbiddenError(err):
			return NewForbiddenGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customError.IsConflictError(err):
			return NewConflictGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customError.IsInternalServerError(err):
			return NewInternalServerGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customError.IsUnMarshalingError(err):
			return NewInternalServerGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customError.IsMarshalingError(err):
			return NewInternalServerGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)

		case customError.IsCustomError(err):
			return NewGrpcError(codes.Internal, customErr.Code(), codes.Internal.String(), customErr.Message(), customErr.Details(), stackTrace)
		// case error.Is(err, context.DeadlineExceeded):
		// 	return NewGrpcError(codes.DeadlineExceeded, customErr.Code(), errorTitles.ErrRequestTimeoutTitle, err.Error(), stackTrace)
		default:
			return NewInternalServerGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		}
	}

	return nil
}
