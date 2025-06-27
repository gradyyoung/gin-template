package errs

import "fmt"

type SystemError struct {
	Code int
	Msg  string
}

func (e *SystemError) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}

func NewSystemError(code int, msg string) error {
	return &SystemError{
		Code: code,
		Msg:  msg,
	}
}
