package customErrors

import (
	"github.com/pkg/errors"
)

func NewUnMarshalingError(message string, code int, details map[string]string) error {
	ume := &unMarshalingError{
		CustomError: NewCustomError(nil, code, message, details),
	}
	// stackErr := error.WithStack(ne)

	return ume
}

func NewUnMarshalingErrorWrap(err error, message string, code int, details map[string]string) error {
	ume := &unMarshalingError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(ume)

	return stackErr
}

type unMarshalingError struct {
	CustomError
}

type UnMarshalingError interface {
	CustomError
	IsUnMarshalingError() bool
}

func (u *unMarshalingError) IsUnMarshalingError() bool {
	return true
}

func IsUnMarshalingError(err error) bool {
	var unMarshalingError UnMarshalingError

	if errors.As(err, &unMarshalingError) {
		return unMarshalingError.IsUnMarshalingError()
	}

	return false
}
