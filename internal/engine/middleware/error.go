package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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
				// 记录异常
				logrus.Errorf("Panic: %+v\n", err)
				// 返回500错误
				ctx.JSON(http.StatusOK, dto.FailedMsgResponse(http.StatusInternalServerError, "服务器内部错误！"))
			}
		}()
		ctx.Next()
		// 处理错误
		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err
			// 按照错误类型处理
			switch e := err.(type) {
			case *errs.SystemError:
				ctx.JSON(http.StatusOK, dto.FailedResponse(e.Code, e.Error()))
			default:
				logrus.Errorf("Error: %+v\n", e)
				ctx.JSON(http.StatusOK, dto.FailedMsgResponse(http.StatusInternalServerError, "服务器内部错误！"))
			}
		}
	}
}
