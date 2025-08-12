package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"ygang.top/gin-template/internal/dto"
	"ygang.top/gin-template/internal/errs"
)

type ErrorHandleMiddleware struct {
}

func NewErrorHandleMiddleware() *ErrorHandleMiddleware {
	return &ErrorHandleMiddleware{}
}

func (e *ErrorHandleMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 设置一个recover，防止应用崩溃
		defer func() {
			if err := recover(); err != nil {
				// 记录异常，并打印堆栈
				logrus.Errorf("Panic: %+v\n%s\n", err, debug.Stack())
				// 返回500错误
				ctx.JSON(http.StatusOK, dto.FailedErrResponse(http.StatusInternalServerError, errs.InternalServerError))
			}
		}()
		ctx.Next()
		// 处理错误
		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err
			// 按照错误类型处理
			switch er := errors.Cause(err).(type) {
			case *errs.SystemError:
				ctx.JSON(http.StatusOK, dto.FailedResponse(er.Code, er.Error()))
			default:
				logrus.Errorf("Error: %+v\n", err)
				ctx.JSON(http.StatusOK, dto.FailedErrResponse(http.StatusInternalServerError, errs.InternalServerError))
			}
		}
	}
}
