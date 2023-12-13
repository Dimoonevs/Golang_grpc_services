package error

type ErrorResp struct {
	StatusResp int
	ErrMsg     string
}

func (e *ErrorResp) Error() string {
	return e.ErrMsg
}

var _ error = (*ErrorResp)(nil)
