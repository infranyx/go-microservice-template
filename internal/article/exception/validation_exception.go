package article_exception

import (
	customErrors "github.com/infranyx/go-grpc-template/shared/error/custom_error"
)

func NewCreateArticleValidationErr(err error) error {
	bad := customErrors.NewBadRequestErrorWrap(err, 2000, "Article validation failed")
	return bad
}
