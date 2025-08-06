package errs

var (
	InternalServerError = NewSystemError(500, "服务器内部错误！")
)
