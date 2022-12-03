package article_exception

import (
	customErrors "github.com/infranyx/go-grpc-template/shared/error/custom_error"
)

func NewCreateArticleValidationErr(err error) error {
	bad := customErrors.NewBadRequestError("validation failed", 2000, nil)
	return bad
}
