package customEcho

import (
	"context"
	"strings"

	"github.com/infranyx/go-grpc-template/pkg/constant"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type echoHttpServer struct {
	echo   *echo.Echo
	config *EchoHttpConfig
}

type EchoHttpServer interface {
	RunHttpServer(ctx context.Context, configEcho func(echoServer *echo.Echo)) error
	GracefulShutdown(ctx context.Context) error
	ApplyVersioningFromHeader()
	GetEchoInstance() *echo.Echo
	SetupDefaultMiddlewares()
	AddMiddlewares(middlewares ...echo.MiddlewareFunc)
	// ConfigGroup(groupName string, groupFunc func(group *echo.Group))
}

type EchoHttpConfig struct {
	Port        string
	Development bool
	Timeout     int
	Host        string
}

func NewEchoHttpServer(config *EchoHttpConfig) *echoHttpServer {
	return &echoHttpServer{echo: echo.New()}
}

func (s *echoHttpServer) RunHttpServer(ctx context.Context, configEcho func(echo *echo.Echo)) error {
	s.echo.Server.ReadTimeout = constant.ReadTimeout
	s.echo.Server.WriteTimeout = constant.WriteTimeout
	s.echo.Server.MaxHeaderBytes = constant.MaxHeaderBytes

	if configEcho != nil {
		configEcho(s.echo)
	}

	go func() {
		for {
			<-ctx.Done()
			logger.Zap.Sugar().Infof("Http server is shutting down PORT: {%s}", s.config.Port)
			if err := s.GracefulShutdown(ctx); err != nil {
				logger.Zap.Sugar().Warnf("(Shutdown) err: {%v}", err)
			}
		}
	}()
	return s.echo.Start(s.config.Port)
}

func (s *echoHttpServer) AddMiddlewares(middlewares ...echo.MiddlewareFunc) {
	if len(middlewares) > 0 {
		s.echo.Use(middlewares...)
	}
}

func (s *echoHttpServer) GracefulShutdown(ctx context.Context) error {
	err := s.echo.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *echoHttpServer) SetupDefaultMiddlewares() {
	// handling internal echo middlewares logs with our log provider
	// s.echo.Use(echozap.ZapLogger(logger.Zap))
	s.echo.HideBanner = false
	// s.echo.HTTPErrorHandler = customHadnlers.ProblemHandler

	// s.echo.Use(otelTracer.Middleware(s.config.Name))
	s.echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogContentLength: true,
		LogLatency:       true,
		LogError:         false,
		LogMethod:        true,
		LogRequestID:     true,
		LogURI:           true,
		LogResponseSize:  true,
		LogURIPath:       true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			// logger.Zap.Sugar().Infow(fmt.Sprintf("[Request Middleware] REQUEST: uri: %v, status: %v\n", v.URI, v.Status), logger.Fields{"URI": v.URI, "Status": v.Status})
			return nil
		},
	}))
	s.echo.Use(middleware.BodyLimit(constant.BodyLimit))
	s.echo.Use(middleware.RequestID())
	s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: constant.GzipLevel,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
}

func (s *echoHttpServer) GetEchoInstance() *echo.Echo {
	return s.echo
}
