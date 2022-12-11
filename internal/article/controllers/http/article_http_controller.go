package articleHttp

import (
	"net/http"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	articleException "github.com/infranyx/go-grpc-template/internal/article/exception"
	"github.com/labstack/echo/v4"
)

type articleHttpController struct {
	articleUC articleDomain.ArticleUseCase
}

func NewArticleHttpController(uc articleDomain.ArticleUseCase) articleDomain.ArticleHttpController {
	return &articleHttpController{
		articleUC: uc,
	}
}

func (ac articleHttpController) Create(c echo.Context) error {
	aDto := new(articleDto.CreateArticle)
	if err := c.Bind(aDto); err != nil {
		return articleException.ArticleBindingExc()
	}

	if err := aDto.ValidateCreateArticleDto(); err != nil {
		return articleException.CreateArticleValidationExc(err)
	}

	article, err := ac.articleUC.Create(c.Request().Context(), aDto)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, article)
}
