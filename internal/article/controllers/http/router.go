package articleHttp

import (
	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	"github.com/labstack/echo/v4"
)

type Router struct {
	articleCtrl articleDomain.ArticleHttpController
}

func NewArticleAPI(ahc articleDomain.ArticleHttpController) *Router {
	return &Router{
		articleCtrl: ahc,
	}
}

func (r *Router) Register(e *echo.Group) {
	e.POST("/article", r.articleCtrl.Create)
}
