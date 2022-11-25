package customErrors

import (
	"github.com/pkg/errors"
)

func NewValidationError(message string, code int) error {
	bad := NewBadRequestError(message, code)
	customErr := GetCustomError(bad)
	ue := &validationError{
		BadRequestError: customErr.(BadRequestError),
	}
	stackErr := errors.WithStack(ue)

	return stackErr
}

func NewValidationErrorWrap(err error, code int, message string) error {
	bad := NewBadRequestErrorWrap(err, code, message)
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
	//us, ok := grpc_errors.Cause(err).(ValidationError)
	if errors.As(err, &validationError) {
		return validationError.IsValidationError()
	}

	return false
}
