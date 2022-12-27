package articleHttpController

import (
	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	"github.com/labstack/echo/v4"
)

type Router struct {
	controller articleDomain.HttpController
}

func NewRouter(controller articleDomain.HttpController) *Router {
	return &Router{
		controller: controller,
	}
}

func (r *Router) Register(e *echo.Group) {
	e.POST("/article", r.controller.Create)
}
