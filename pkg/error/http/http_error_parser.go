package httpError

import (
	"net/http"

	"google.golang.org/grpc/codes"

	errorConstant "github.com/infranyx/go-grpc-template/pkg/constant/error"
	customError "github.com/infranyx/go-grpc-template/pkg/error/custom_error"
)

func ParseError(err error) HttpErr {
	customErr := customError.AsCustomError(err)
	if customErr == nil {
		err =
			customError.NewInternalServerErrorWrap(err, errorConstant.ErrInfo.InternalServerErr.Msg, errorConstant.ErrInfo.InternalServerErr.Code, nil)
		customErr = customError.AsCustomError(err)
		return NewHttpError(http.StatusInternalServerError, customErr.Code(), errorConstant.ErrInternalServerErrorTitle, customErr.Error(), customErr.Details())
	}

	if err != nil {
		switch {
		case customError.IsValidationError(err):
			return NewHttpValidationError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsBadRequestError(err):
			return NewHttpBadRequestError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsMethodNotAllowedError(err):
			return NewMethodNotAllowedError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsNotFoundError(err):
			return NewHttpNotFoundError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsInternalServerError(err):
			return NewHttpInternalServerError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsForbiddenError(err):
			return NewHttpForbiddenError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsUnAuthorizedError(err):
			return NewHttpUnAuthorizedError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsDomainError(err):
			return NewHttpDomainError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsApplicationError(err):
			return NewHttpApplicationError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsConflictError(err):
			return NewHttpConflictError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsUnMarshalingError(err):
			return NewHttpInternalServerError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsMarshalingError(err):
			return NewHttpInternalServerError(customErr.Code(), customErr.Message(), customErr.Details())

		case customError.IsCustomError(err):
			return NewHttpError(http.StatusInternalServerError, customErr.Code(), codes.Internal.String(), customErr.Message(), customErr.Details())

		// case error.Is(err, context.DeadlineExceeded):
		// 	return NewHttpError(codes.DeadlineExceeded, customErr.Code(), errorTitles.ErrRequestTimeoutTitle, err.Error(), stackTrace)

		default:
			return NewHttpInternalServerError(customErr.Code(), customErr.Message(), customErr.Details())
		}
	}

	return nil
}
