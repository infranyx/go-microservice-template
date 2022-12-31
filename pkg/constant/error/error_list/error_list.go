package errorList

var InternalErrorList *internalErrorList

type internalErrorList struct {
	ValidationError     ErrorList
	InternalServerError ErrorList
	NotFoundError       ErrorList
	ArticleExceptions   ArticleErrorList
}

type ErrorList struct {
	Msg  string
	Code int
}

type ArticleErrorList struct {
	BindingError ErrorList
}

func init() {
	InternalErrorList = &internalErrorList{
		// 1000 - 1999 : BoilerPlate Err
		// 2000 - 2999 : Custom Err Per Service
		// .
		// .
		// .
		// 8000 - 8999 : Third-party
		// 9000 - 9999 : FATAL

		InternalServerError: ErrorList{
			Msg:  "internal server error",
			Code: 1000,
		},

		ValidationError: ErrorList{
			Msg:  "request validation failed",
			Code: 1001,
		},

		NotFoundError: ErrorList{
			Msg:  "not found",
			Code: 1002,
		},

		ArticleExceptions: ArticleErrorList{
			BindingError: ErrorList{
				Msg:  "binding failed",
				Code: 3002,
			},
		},
	}
}
