package grpcErrors

import (
	"context"
	"database/sql"

	validator "github.com/go-ozzo/ozzo-validation"
	errorTitles "github.com/infranyx/go-grpc-template/constant/errors"
	customErrors "github.com/infranyx/go-grpc-template/shared/error/custom_error"
	errorUtils "github.com/infranyx/go-grpc-template/utils/error_utils"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

//https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md
//https://github.com/grpc/grpc/blob/master/doc/statuscodes.md

func ParseError(err error) GrpcErr {
	customErr := customErrors.GetCustomError(err)
	var validatorErr validator.Errors
	stackTrace := errorUtils.ErrorsWithStack(err)

	if err != nil {
		switch {
		case customErrors.IsDomainError(err):
			return NewDomainGrpcError(codes.InvalidArgument, customErr.Code(), customErr.Error(), stackTrace)
		case customErrors.IsApplicationError(err):
			return NewApplicationGrpcError(codes.Internal, customErr.Code(), customErr.Error(), stackTrace)
		case customErrors.IsApiError(err):
			return NewApiGrpcError(codes.Internal, customErr.Code(), customErr.Error(), stackTrace)
		case customErrors.IsBadRequestError(err):
			return NewBadRequestGrpcError(customErr.Code(), customErr.Error(), stackTrace)
		case customErrors.IsNotFoundError(err):
			return NewNotFoundErrorGrpcError(customErr.Code(), customErr.Error(), stackTrace)
		case customErrors.IsValidationError(err):
			return NewValidationGrpcError(customErr.Code(), customErr.Error(), stackTrace)
		case customErrors.IsUnAuthorizedError(err):
			return NewUnAuthorizedErrorGrpcError(customErr.Code(), customErr.Error(), stackTrace)
		case customErrors.IsForbiddenError(err):
			return NewForbiddenGrpcError(customErr.Code(), customErr.Error(), stackTrace)
		case customErrors.IsConflictError(err):
			return NewConflictGrpcError(customErr.Code(), customErr.Error(), stackTrace)
		case customErrors.IsInternalServerError(err):
			return NewInternalServerGrpcError(customErr.Code(), customErr.Error(), stackTrace)
		case customErrors.IsCustomError(err):
			return NewGrpcError(codes.Internal, customErr.Code(), codes.Internal.String(), customErr.Error(), stackTrace)
		case customErrors.IsUnMarshalingError(err):
			return NewInternalServerGrpcError(customErr.Code(), customErr.Error(), stackTrace)
		case customErrors.IsMarshalingError(err):
			return NewInternalServerGrpcError(customErr.Code(), customErr.Error(), stackTrace)
		case errors.Is(err, sql.ErrNoRows):
			return NewNotFoundErrorGrpcError(customErr.Code(), err.Error(), stackTrace)
		case errors.Is(err, context.DeadlineExceeded):
			return NewGrpcError(codes.DeadlineExceeded, customErr.Code(), errorTitles.ErrRequestTimeoutTitle, err.Error(), stackTrace)
		case errors.As(err, &validatorErr):
			return NewValidationGrpcError(customErr.Code(), validatorErr.Error(), stackTrace)
		default:
			return NewInternalServerGrpcError(customErr.Code(), err.Error(), stackTrace)
		}
	}

	return nil
}
