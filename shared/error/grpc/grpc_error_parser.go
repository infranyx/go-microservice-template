package grpc_errors

import (
	customErrors "github.com/infranyx/go-grpc-template/shared/error/custom_error"
	errorUtils "github.com/infranyx/go-grpc-template/utils/error_utils"
	"google.golang.org/grpc/codes"
)

//https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md
//https://github.com/grpc/grpc/blob/master/doc/statuscodes.md

func ParseError(err error) GrpcErr {
	customErr := customErrors.GetCustomError(err)
	stackTrace := errorUtils.RootStackTrace(err)

	if customErr == nil {
		err = customErrors.NewApiErrorWrap(err, "Unkown Error", 0, nil)
		customErr = customErrors.GetCustomError(err)
		stackTrace = errorUtils.RootStackTrace(err)
		return NewGrpcError(codes.Internal, customErr.Code(), codes.Internal.String(), customErr.Error(), customErr.Details(), stackTrace)
	}

	if err != nil {
		switch {
		case customErrors.IsDomainError(err):
			return NewDomainGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customErrors.IsApplicationError(err):
			return NewApplicationGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customErrors.IsApiError(err):
			return NewApiGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customErrors.IsBadRequestError(err):
			return NewBadRequestGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customErrors.IsNotFoundError(err):
			return NewNotFoundErrorGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customErrors.IsValidationError(err):
			return NewValidationGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customErrors.IsUnAuthorizedError(err):
			return NewUnAuthorizedErrorGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customErrors.IsForbiddenError(err):
			return NewForbiddenGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customErrors.IsConflictError(err):
			return NewConflictGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customErrors.IsInternalServerError(err):
			return NewInternalServerGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customErrors.IsUnMarshalingError(err):
			return NewInternalServerGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		case customErrors.IsMarshalingError(err):
			return NewInternalServerGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)

		case customErrors.IsCustomError(err):
			return NewGrpcError(codes.Internal, customErr.Code(), codes.Internal.String(), customErr.Message(), customErr.Details(), stackTrace)
		// case errors.Is(err, context.DeadlineExceeded):
		// 	return NewGrpcError(codes.DeadlineExceeded, customErr.Code(), errorTitles.ErrRequestTimeoutTitle, err.Error(), stackTrace)
		default:
			return NewInternalServerGrpcError(customErr.Code(), customErr.Message(), customErr.Details(), stackTrace)
		}
	}

	return nil
}
