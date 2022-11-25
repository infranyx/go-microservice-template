package customErrors

import (
	"github.com/pkg/errors"
)

func NewDomainError(message string, code int) error {
	de := &domainError{
		CustomError: NewCustomError(nil, code, message),
	}
	stackErr := errors.WithStack(de)

	return stackErr
}

func NewDomainErrorWrap(err error, code int, message string) error {
	de := &domainError{
		CustomError: NewCustomError(err, code, message),
	}
	stackErr := errors.WithStack(de)

	return stackErr
}

type domainError struct {
	CustomError
}

type DomainError interface {
	CustomError
	IsDomainError() bool
}

func (d *domainError) IsDomainError() bool {
	return true
}

func IsDomainError(err error) bool {
	var domainErr DomainError
	//us, ok := grpc_errors.Cause(err).(DomainError)
	if errors.As(err, &domainErr) {
		return domainErr.IsDomainError()
	}
	return false
}
