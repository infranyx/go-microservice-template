package customError

import (
	"github.com/pkg/errors"
)

type domainError struct {
	CustomError
}

func (e *domainError) IsDomainError() bool {
	return true
}

type DomainError interface {
	CustomError
	IsDomainError() bool
}

func IsDomainError(e error) bool {
	var domainError DomainError

	if errors.As(e, &domainError) {
		return domainError.IsDomainError()
	}
	return false
}

func NewDomainError(message string, code int, details map[string]string) error {
	e := &domainError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return e
}

func NewDomainErrorWrap(err error, message string, code int, details map[string]string) error {
	e := &domainError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(e)

	return stackErr
}
