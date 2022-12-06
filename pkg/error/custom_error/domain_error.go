package customError

import (
	"github.com/pkg/errors"
)

type domainError struct {
	CustomError
}

func (de *domainError) IsDomainError() bool {
	return true
}

type DomainError interface {
	CustomError
	IsDomainError() bool
}

func IsDomainError(err error) bool {
	var domainErr DomainError

	if errors.As(err, &domainErr) {
		return domainErr.IsDomainError()
	}
	return false
}

func NewDomainError(message string, code int, details map[string]string) error {
	de := &domainError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return de
}

func NewDomainErrorWrap(err error, message string, code int, details map[string]string) error {
	de := &domainError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(de)

	return stackErr
}
