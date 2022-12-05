package articleException

import (
	errorUtils "github.com/infranyx/go-grpc-template/pkg/error/error_utils"
	customErrors "github.com/infranyx/go-grpc-template/shared/error/custom_error"
)

func CreateArticleValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrHandler(err)
	if ie != nil {
		return ie
	}
	return customErrors.NewBadRequestError("validation failed", 2000, ve)
}
