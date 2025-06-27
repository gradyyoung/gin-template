package errs

import (
	"net/http"
)

var (
	UserNameOrPasswordError = NewSystemError(http.StatusNonAuthoritativeInfo, "用户名或密码错误！")
	UserNonAuthoritative    = NewSystemError(http.StatusNonAuthoritativeInfo, "用户未认证！")
)
