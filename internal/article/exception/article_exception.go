package articleException

import (
	errorList "github.com/infranyx/go-grpc-template/pkg/constant/error/error_list"
	customErrors "github.com/infranyx/go-grpc-template/pkg/error/custom_error"
	errorUtils "github.com/infranyx/go-grpc-template/pkg/error/error_utils"
)

func CreateArticleValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrorHandler(err)
	if ie != nil {
		return ie
	}

	validationErrorCode := errorList.InternalErrorList.ValidationError
	return customErrors.NewValidationError(validationErrorCode.Msg, validationErrorCode.Code, ve)
}

func ArticleBindingExc() error {
	articleBindingError := errorList.InternalErrorList.ArticleExceptions.BindingError
	return customErrors.NewBadRequestError(articleBindingError.Msg, articleBindingError.Code, nil)
}
