package articleHttpController

import (
	"github.com/labstack/echo/v4"

	articleDomain "github.com/infranyx/go-microservice-template/internal/article/domain"
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
	e.POST("/article", r.controller.CreateArticle)
}
