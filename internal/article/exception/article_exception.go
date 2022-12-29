package articleException

import (
	customErrors "github.com/infranyx/go-grpc-template/pkg/error/custom_error"
	errorUtils "github.com/infranyx/go-grpc-template/pkg/error/error_utils"
)

func CreateArticleValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrorHandler(err)
	if ie != nil {
		return ie
	}
	return customErrors.NewValidationError("validation failed", 2000, ve)
}

func ArticleBindingExc() error {
	return customErrors.NewBadRequestError("binding failed", 3000, nil)
}
