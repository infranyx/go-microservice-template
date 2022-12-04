package customErrors

import (
	"github.com/pkg/errors"
)

func NewApiError(message string, code int, details []ErrorDetail) error {
	ae := &apiError{
		CustomError: NewCustomError(nil, code, message, details),
	}
	// stackErr := error.WithStack(ae)

	return ae
}

func NewApiErrorWrap(err error, message string, code int, details []ErrorDetail) error {
	ae := &apiError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(ae)

	return stackErr
}

type apiError struct {
	CustomError
}

type ApiError interface {
	CustomError
	IsApiError() bool
}

func (a *apiError) IsApiError() bool {
	return true
}

func IsApiError(err error) bool {
	var apiError ApiError

	if errors.As(err, &apiError) {
		return apiError.IsApiError()
	}

	return false
}
