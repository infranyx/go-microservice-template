package echoErrorHandler

import (
	"net/http"

	loggerConst "github.com/infranyx/go-grpc-template/pkg/constant/logger"
	errorUtils "github.com/infranyx/go-grpc-template/pkg/error/error_utils"
	httpError "github.com/infranyx/go-grpc-template/pkg/error/http"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func ErrorHandler(err error, c echo.Context) {
	he := httpError.ParseError(err)
	if !c.Response().Committed {
		if _, err := he.WriteTo(c.Response()); err != nil {
			logger.Zap.Error(
				err.Error(),
				zap.String(loggerConst.TYPE, loggerConst.GRPC),
				zap.String(loggerConst.TITILE, he.GetTitle()),
				zap.Int(loggerConst.CODE, he.GetCode()),
				zap.String(loggerConst.STATUS, http.StatusText(he.GetStatus())),
				zap.Time(loggerConst.TIME, he.GetTimestamp()),
				zap.Any(loggerConst.DETAILS, he.GetDetails()),
				zap.String(loggerConst.STACK_TRACE, errorUtils.RootStackTrace(err)),
			)
		}
	}
}
