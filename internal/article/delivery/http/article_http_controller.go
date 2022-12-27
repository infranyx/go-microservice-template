package articleHttpController

import (
	"net/http"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	articleException "github.com/infranyx/go-grpc-template/internal/article/exception"
	"github.com/labstack/echo/v4"
)

type controller struct {
	useCase articleDomain.UseCase
}

func NewController(uc articleDomain.UseCase) articleDomain.HttpController {
	return &controller{
		useCase: uc,
	}
}

func (c controller) CreateArticle(ctx echo.Context) error {
	aDto := new(articleDto.CreateArticleRequestDto)
	if err := ctx.Bind(aDto); err != nil {
		return articleException.ArticleBindingExc()
	}

	if err := aDto.ValidateCreateArticleDto(); err != nil {
		return articleException.CreateArticleValidationExc(err)
	}

	article, err := c.useCase.CreateArticle(ctx.Request().Context(), aDto)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, article)
}
