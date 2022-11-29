package errors

var ErrDetails *errDetails

type errDetails struct {
	ErrBadRequest     detaill
	ErrConflict       detaill
	ErrNotFound       detaill
	ErrUnauthorized   detaill
	ErrForbidden      detaill
	ErrRequestTimeout detaill
	ErrInternal       detaill
	ErrDomain         detaill
	ErrApplication    detaill
	ErrApi            detaill
}

type detaill struct {
	Title string
	Code  int
}

// init is invoked before main()
func init() {
	ErrDetails = &errDetails{
		ErrBadRequest: detaill{
			Title: "Bad Request",
			Code:  1000,
		},
		ErrConflict: detaill{
			Title: "Conflict Error",
			Code:  1001,
		},
		ErrNotFound: detaill{
			Title: "Not Found",
			Code:  1002,
		},
		ErrUnauthorized: detaill{
			Title: "Unauthorized",
			Code:  1003,
		},
		ErrForbidden: detaill{
			Title: "Forbidden",
			Code:  1004,
		},
		ErrRequestTimeout: detaill{
			Title: "Request Timeout",
			Code:  1005,
		},
		ErrInternal: detaill{
			Title: "Internal Server Error",
			Code:  1006,
		},
		ErrDomain: detaill{
			Title: "Request Timeout",
			Code:  1007,
		},
		ErrApplication: detaill{
			Title: "Application Service Error",
			Code:  1008,
		},
		ErrApi: detaill{
			Title: "Api Error",
			Code:  1009,
		},
	}
}
