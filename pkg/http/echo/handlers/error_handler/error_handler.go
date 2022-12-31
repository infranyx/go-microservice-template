package echoErrorHandler

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	sentryEcho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	loggerConstant "github.com/infranyx/go-grpc-template/pkg/constant/logger"
	customErrors "github.com/infranyx/go-grpc-template/pkg/error/custom_error"
	errorUtils "github.com/infranyx/go-grpc-template/pkg/error/error_utils"
	httpError "github.com/infranyx/go-grpc-template/pkg/error/http"
	"github.com/infranyx/go-grpc-template/pkg/logger"
)

func ErrorHandler(err error, c echo.Context) {
	// default echo errors
	echoHttpError, ok := err.(*echo.HTTPError)
	var httpResponseError httpError.HttpErr

	if ok {
		httpResponseError = httpError.NewHttpError(echoHttpError.Code, echoHttpError.Code, http.StatusText(echoHttpError.Code), http.StatusText(echoHttpError.Code), nil)
	} else {
		// parse as a custom error
		httpResponseError = httpError.ParseError(err)
	}

	if customErrors.IsInternalServerError(err) {
		hub := sentryEcho.GetHubFromContext(c)
		if hub != nil {
			hub.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetLevel(sentry.LevelError)
				scope.SetContext("HttpErr", map[string]interface{}{
					loggerConstant.TYPE:        loggerConstant.HTTP,
					loggerConstant.TITILE:      httpResponseError.GetTitle(),
					loggerConstant.CODE:        httpResponseError.GetCode(),
					loggerConstant.STATUS:      http.StatusText(httpResponseError.GetStatus()),
					loggerConstant.TIME:        httpResponseError.GetTimestamp(),
					loggerConstant.DETAILS:     httpResponseError.GetDetails(),
					loggerConstant.STACK_TRACE: errorUtils.RootStackTrace(err),
				})
			})
			hub.CaptureException(err)
		}
	}

	if !c.Response().Committed {
		if _, err := httpResponseError.WriteTo(c.Response()); err != nil {
			logger.Zap.Sugar().Error(`error while writing http error response: %v`, err)
		}
		logger.Zap.Error(
			err.Error(),
			zap.String(loggerConstant.TYPE, loggerConstant.HTTP),
			zap.String(loggerConstant.TITILE, httpResponseError.GetTitle()),
			zap.Int(loggerConstant.CODE, httpResponseError.GetCode()),
			zap.String(loggerConstant.STATUS, http.StatusText(httpResponseError.GetStatus())),
			zap.Time(loggerConstant.TIME, httpResponseError.GetTimestamp()),
			zap.Any(loggerConstant.DETAILS, httpResponseError.GetDetails()),
			zap.String(loggerConstant.STACK_TRACE, errorUtils.RootStackTrace(err)),
		)
	}
}
