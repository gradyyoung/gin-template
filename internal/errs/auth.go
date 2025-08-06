package errs

import (
	"net/http"
)

var (
	UserNameOrPasswordError = NewSystemError(http.StatusNonAuthoritativeInfo, "用户名或密码错误！")
	UserNonAuthoritative    = NewSystemError(http.StatusNonAuthoritativeInfo, "用户未认证！")
	NoAccessPermission      = NewSystemError(http.StatusForbidden, "无访问权限！")
	UserNotExist            = NewSystemError(http.StatusNonAuthoritativeInfo, "用户不存在！")
)
