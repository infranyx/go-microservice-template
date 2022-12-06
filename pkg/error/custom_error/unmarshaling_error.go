package customError

import (
	"github.com/pkg/errors"
)

type unMarshalingError struct {
	CustomError
}

func (ue *unMarshalingError) IsUnMarshalingError() bool {
	return true
}

type UnMarshalingError interface {
	CustomError
	IsUnMarshalingError() bool
}

func IsUnMarshalingError(err error) bool {
	var unMarshalingError UnMarshalingError

	if errors.As(err, &unMarshalingError) {
		return unMarshalingError.IsUnMarshalingError()
	}

	return false
}

func NewUnMarshalingError(message string, code int, details map[string]string) error {
	ue := &unMarshalingError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return ue
}

func NewUnMarshalingErrorWrap(err error, message string, code int, details map[string]string) error {
	ue := &unMarshalingError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(ue)

	return stackErr
}
