package articleHttpController

import (
	"net/http"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	articleException "github.com/infranyx/go-grpc-template/internal/article/exception"
	"github.com/labstack/echo/v4"
)

type controller struct {
	useCase articleDomain.ArticleUseCase
}

func NewController(uc articleDomain.ArticleUseCase) articleDomain.ArticleHttpController {
	return &controller{
		useCase: uc,
	}
}

func (c controller) Create(ctx echo.Context) error {
	aDto := new(articleDto.CreateArticleDto)
	if err := ctx.Bind(aDto); err != nil {
		return articleException.ArticleBindingExc()
	}

	if err := aDto.ValidateCreateArticleDto(); err != nil {
		return articleException.CreateArticleValidationExc(err)
	}

	article, err := c.useCase.Create(ctx.Request().Context(), aDto)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, article)
}
