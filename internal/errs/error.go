package errs

import "fmt"

type SystemError struct {
	Code int
	Msg  string
}

func (e *SystemError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Msg)
}

func NewSystemError(code int, msg string) *SystemError {
	return &SystemError{
		Code: code,
		Msg:  msg,
	}
}
