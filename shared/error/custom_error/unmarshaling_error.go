package customErrors

import (
	"github.com/pkg/errors"
)

func NewUnMarshalingError(message string, code int) error {
	internal := NewInternalServerError(message, code)
	customErr := GetCustomError(internal)
	ue := &unMarshalingError{
		InternalServerError: customErr.(InternalServerError),
	}
	stackErr := errors.WithStack(ue)

	return stackErr
}

func NewUnMarshalingErrorWrap(err error, code int, message string) error {
	internal := NewInternalServerErrorWrap(err, code, message)
	customErr := GetCustomError(internal)
	ue := &unMarshalingError{
		InternalServerError: customErr.(InternalServerError),
	}
	stackErr := errors.WithStack(ue)

	return stackErr
}

type unMarshalingError struct {
	InternalServerError
}

type UnMarshalingError interface {
	InternalServerError
	IsUnMarshalingError() bool
}

func (u *unMarshalingError) IsUnMarshalingError() bool {
	return true
}

func IsUnMarshalingError(err error) bool {
	var unMarshalingError UnMarshalingError
	//us, ok := grpc_errors.Cause(err).(UnMarshalingError)
	if errors.As(err, &unMarshalingError) {
		return unMarshalingError.IsUnMarshalingError()
	}

	return false
}
