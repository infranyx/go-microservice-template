package httpError

import (
	"net/http"
	"time"

	errConst "github.com/infranyx/go-grpc-template/pkg/constant/error"
)

func NewHttpValidationError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errConst.ErrValidationFailedTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusBadRequest,
		Timestamp: time.Now(),
	}
}

func NewHttpConflictError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errConst.ErrConflictTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusConflict,
		Timestamp: time.Now(),
	}
}

func NewHttpBadRequestError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errConst.ErrBadRequestTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusBadRequest,
		Timestamp: time.Now(),
	}
}

func NewMethodNotAllowedError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errConst.ErrMethodNotAllowed,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusMethodNotAllowed,
		Timestamp: time.Now(),
	}
}

func NewHttpNotFoundError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errConst.ErrNotFoundTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusNotFound,
		Timestamp: time.Now(),
	}
}

func NewHttpUnAuthorizedError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errConst.ErrUnauthorizedTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusUnauthorized,
		Timestamp: time.Now(),
	}
}

func NewHttpForbiddenError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errConst.ErrForbiddenTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusForbidden,
		Timestamp: time.Now(),
	}
}

func NewHttpInternalServerError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errConst.ErrInternalServerErrorTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusInternalServerError,
		Timestamp: time.Now(),
	}
}

func NewHttpDomainError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errConst.ErrDomainTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusBadRequest,
		Timestamp: time.Now(),
	}
}

func NewHttpApplicationError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errConst.ErrApplicationTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusInternalServerError,
		Timestamp: time.Now(),
	}
}

func NewHttpApiError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errConst.ErrApiTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusInternalServerError,
		Timestamp: time.Now(),
	}
}
