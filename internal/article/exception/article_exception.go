package articleException

import (
	errorList "github.com/infranyx/go-microservice-template/pkg/constant/error/error_list"
	customErrors "github.com/infranyx/go-microservice-template/pkg/error/custom_error"
	errorUtils "github.com/infranyx/go-microservice-template/pkg/error/error_utils"
)

func CreateArticleValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrorHandler(err)
	if ie != nil {
		return ie
	}

	validationError := errorList.InternalErrorList.ValidationError
	return customErrors.NewValidationError(validationError.Msg, validationError.Code, ve)
}

func ArticleBindingExc() error {
	articleBindingError := errorList.InternalErrorList.ArticleExceptions.BindingError
	return customErrors.NewBadRequestError(articleBindingError.Msg, articleBindingError.Code, nil)
}
