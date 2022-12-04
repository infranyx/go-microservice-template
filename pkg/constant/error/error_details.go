package errConst

var ErrDetails *errDetails

type errDetails struct {
	ValidationErr detaill
}

type detaill struct {
	Msg  string
	Code int
}

func init() {
	ErrDetails = &errDetails{
		ValidationErr: detaill{
			Msg:  "validation failed",
			Code: 1000,
		},
	}
}
