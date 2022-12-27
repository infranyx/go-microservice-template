package healthCheckHttp

import (
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	"github.com/labstack/echo/v4"
)

type Router struct {
	controller healthCheckDomain.HttpController
}

func NewRouter(controller healthCheckDomain.HttpController) *Router {
	return &Router{
		controller: controller,
	}
}

func (r *Router) Register(e *echo.Group) {
	e.GET("/health", r.controller.Check)
}
