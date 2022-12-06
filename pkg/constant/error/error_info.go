package errConst

var ErrInfo *errInfo

type errInfo struct {
	ValidationErr     eInfo
	InternalServerErr eInfo
}

type eInfo struct {
	Msg  string
	Code int
}

func init() {
	ErrInfo = &errInfo{
		// 1000 - 1999 : Boiler-Plate Err
		// 2000 - 2999 : Custom Err Per Service
		// .
		// .
		// .
		// 8000 - 8999 : Third-party
		// 9000 - 9999 : FATAL

		ValidationErr: eInfo{
			Msg:  "request validation failed",
			Code: 1000,
		},

		InternalServerErr: eInfo{
			Msg:  "internal server error",
			Code: 1001,
		},
	}
}
