package customError

import (
	"github.com/pkg/errors"
)

type apiError struct {
	CustomError
}

func (ae *apiError) IsApiError() bool {
	return true
}

type ApiError interface {
	CustomError
	IsApiError() bool
}

func IsApiError(err error) bool {
	var apiError ApiError

	if errors.As(err, &apiError) {
		return apiError.IsApiError()
	}

	return false
}

func NewApiError(message string, code int, details map[string]string) error {
	ae := &apiError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return ae
}

func NewApiErrorWrap(err error, message string, code int, details map[string]string) error {
	ae := &apiError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(ae)

	return stackErr
}
