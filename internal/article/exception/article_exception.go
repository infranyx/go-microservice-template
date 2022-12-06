package articleException

import (
	customErrors "github.com/infranyx/go-grpc-template/pkg/error/custom_error"
	errorUtils "github.com/infranyx/go-grpc-template/pkg/error/error_utils"
)

func CreateArticleValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrHandler(err)
	if ie != nil {
		return ie
	}
	return customErrors.NewBadRequestError("validation failed", 2000, ve)
}
