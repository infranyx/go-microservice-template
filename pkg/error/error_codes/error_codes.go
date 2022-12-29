package errorCodes

var InternalErrorCodes *internalErrorCodes

type internalErrorCodes struct {
	ValidationError       ErrorCode
	InternalServerError   ErrorCode
	NotFoundError         ErrorCode
	MethodNotAllowedError ErrorCode
}

type ErrorCode struct {
	Msg  string
	Code int
}

func init() {
	InternalErrorCodes = &internalErrorCodes{
		// 1000 - 1999 : BoilerPlate Err
		// 2000 - 2999 : Custom Err Per Service
		// .
		// .
		// .
		// 8000 - 8999 : Third-party
		// 9000 - 9999 : FATAL

		ValidationError: ErrorCode{
			Msg:  "request validation failed",
			Code: 1000,
		},

		InternalServerError: ErrorCode{
			Msg:  "internal server error",
			Code: 1001,
		},

		NotFoundError: ErrorCode{
			Msg:  "not found",
			Code: 1002,
		},

		MethodNotAllowedError: ErrorCode{
			Msg:  "method not allowed",
			Code: 1003,
		},
	}
}
