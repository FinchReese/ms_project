package errs

type ErrorCode int

type BError struct {
	Code ErrorCode
	Msg  string
}

func NewError(code ErrorCode, msg string) *BError {
	return &BError{
		Code: code,
		Msg:  msg,
	}
}