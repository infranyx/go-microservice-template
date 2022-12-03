package customErrors

import (
	"github.com/pkg/errors"
)

func NewValidationError(message string, code int, details []ErrorDetail) error {
	bad := NewBadRequestError(message, code, details)
	customErr := GetCustomError(bad)
	ve := &validationError{
		BadRequestError: customErr.(BadRequestError),
	}
	// stackErr := errors.WithStack(ue)

	return ve
}

func NewValidationErrorWrap(err error, message string, code int, details []ErrorDetail) error {
	bad := NewBadRequestErrorWrap(err, message, code, details)
	customErr := GetCustomError(bad)
	ue := &validationError{
		BadRequestError: customErr.(BadRequestError),
	}
	stackErr := errors.WithStack(ue)

	return stackErr
}

type validationError struct {
	BadRequestError
}

type ValidationError interface {
	BadRequestError
	IsValidationError() bool
}

func (v *validationError) IsValidationError() bool {
	return true
}

func IsValidationError(err error) bool {
	var validationError ValidationError

	if errors.As(err, &validationError) {
		return validationError.IsValidationError()
	}

	return false
}
