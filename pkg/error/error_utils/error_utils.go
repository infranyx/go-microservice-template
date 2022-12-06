package errorUtils

import (
	"fmt"
	"github.com/infranyx/go-grpc-template/pkg/error/custom_error"
	"strings"

	validator "github.com/go-ozzo/ozzo-validation"
	errorContract "github.com/infranyx/go-grpc-template/pkg/error/contracts"
	"github.com/pkg/errors"
)

// CheckErrMessages check for specific messages contains in the error
func CheckErrMessages(err error, messages ...string) bool {
	for _, message := range messages {
		if strings.Contains(strings.TrimSpace(strings.ToLower(err.Error())), strings.TrimSpace(strings.ToLower(message))) {
			return true
		}
	}
	return false
}

// RootStackTrace returns root stack trace with a string contains just stack trace levels for the given error
func RootStackTrace(err error) string {
	var stackStr string
	for {
		st, ok := err.(errorContract.StackTracer)
		if ok {
			stackStr = fmt.Sprintf("%+v\n", st.StackTrace())

			if !ok {
				break
			}
		}
		err = errors.Unwrap(err)
		if err == nil {
			break
		}
	}

	return stackStr
}

func ValidationErrHandler(err error) (map[string]string, error) {
	var customErr validator.Errors
	if errors.As(err, &customErr) {
		details := make(map[string]string)
		for k, v := range customErr {
			details[k] = v.Error()
		}
		return details, nil
	}
	// TODO : get internal error from constant.
	return nil, customError.NewInternalServerErrorWrap(err, "internal error", 8585, nil)
}
