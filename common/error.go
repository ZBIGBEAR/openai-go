package common

import "fmt"

const (
	UnknowErrCode = 1

	NoOpenAPIKeyErrCode = 2
	NoOpenAPIKeyErrMsg  = "没有配置open_api_key"
)

type Error struct {
	Code int
	Msg  string
	Info []string
}

func (e *Error) Error() string {
	return fmt.Sprintf(`{"code":%d, "msg":%s}`, e.Code, e.Msg)
}

var (
	NoOpenAPIKeyErr = &Error{
		Code: NoOpenAPIKeyErrCode,
		Msg:  NoOpenAPIKeyErrMsg,
	}
)

func NewError(format string, args ...any) error {
	return &Error{
		Code: UnknowErrCode,
		Msg:  fmt.Sprintf(format, args...),
	}
}
