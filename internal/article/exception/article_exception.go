package articleException

import (
	customErrors "github.com/infranyx/go-grpc-template/pkg/error/custom_error"
	errorCodes "github.com/infranyx/go-grpc-template/pkg/error/error_codes"
	errorUtils "github.com/infranyx/go-grpc-template/pkg/error/error_utils"
)

func CreateArticleValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrorHandler(err)
	if ie != nil {
		return ie
	}

	validationErrorCode := errorCodes.InternalErrorCodes.ValidationError
	return customErrors.NewValidationError(validationErrorCode.Msg, validationErrorCode.Code, ve)
}

func ArticleBindingExc() error {
	articleBindingError := errorCodes.InternalErrorCodes.ArticleExceptions.BindingError
	return customErrors.NewBadRequestError(articleBindingError.Msg, articleBindingError.Code, nil)
}
