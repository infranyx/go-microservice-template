package grpcError

import (
	"google.golang.org/grpc/codes"

	customError "github.com/infranyx/go-grpc-template/pkg/error/custom_error"
	errorCodes "github.com/infranyx/go-grpc-template/pkg/error/error_codes"
)

func ParseError(err error) GrpcErr {
	customErr := customError.AsCustomError(err)
	if customErr == nil {
		internalServerErrorCode := errorCodes.InternalErrorCodes.InternalServerError
		err =
			customError.NewInternalServerErrorWrap(err, internalServerErrorCode.Msg, internalServerErrorCode.Code, nil)
		customErr = customError.AsCustomError(err)
		return NewGrpcError(codes.Internal, customErr.Code(), codes.Internal.String(), customErr.Error(), customErr.Details())
	}

	if err != nil {
		switch {
		case customError.IsValidationError(err):
			return NewGrpcValidationError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsBadRequestError(err):
			return NewGrpcBadRequestError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsNotFoundError(err):
			return NewGrpcNotFoundError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsInternalServerError(err):
			return NewGrpcInternalServerError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsForbiddenError(err):
			return NewGrpcForbiddenError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsUnAuthorizedError(err):
			return NewGrpcUnAuthorizedError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsDomainError(err):
			return NewGrpcDomainError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsApplicationError(err):
			return NewGrpcApplicationError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsConflictError(err):
			return NewGrpcConflictError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsUnMarshalingError(err):
			return NewGrpcInternalServerError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsMarshalingError(err):
			return NewGrpcInternalServerError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsCustomError(err):
			return NewGrpcError(codes.Internal, customErr.Code(), codes.Internal.String(), customErr.Message(), customErr.Details())

		// case error.Is(err, context.DeadlineExceeded):
		// 	return NewGrpcError(codes.DeadlineExceeded, customErr.Code(), errorTitles.ErrRequestTimeoutTitle, err.Error(), stackTrace)

		default:
			return NewGrpcInternalServerError(customErr.Code(), customErr.Message(), customErr.Details())
		}
	}

	return nil
}
