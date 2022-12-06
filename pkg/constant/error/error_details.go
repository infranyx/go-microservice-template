package errConst

var ErrDetails *errDetails

type errDetails struct {
	ValidationErr detail
}

type detail struct {
	Msg  string
	Code int
}

func init() {
	ErrDetails = &errDetails{
		ValidationErr: detail{
			Msg:  "validation failed",
			Code: 1000,
		},
	}
}
