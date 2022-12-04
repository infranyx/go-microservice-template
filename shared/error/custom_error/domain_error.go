package customErrors

import (
	"github.com/pkg/errors"
)

func NewDomainError(message string, code int, details []ErrorDetail) error {
	de := &domainError{
		CustomError: NewCustomError(nil, code, message, details),
	}
	// stackErr := error.WithStack(de)

	return de
}

func NewDomainErrorWrap(err error, message string, code int, details []ErrorDetail) error {
	de := &domainError{
		CustomError: NewCustomError(err, code, message, details),
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

	if errors.As(err, &domainErr) {
		return domainErr.IsDomainError()
	}
	return false
}
