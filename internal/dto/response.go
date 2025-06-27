package dto

import "net/http"

// Response API 统一响应格式
type Response struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// NewResponse 创建响应
func NewResponse(code int, success bool, message string, data any) *Response {
	return &Response{
		Code:    code,
		Success: success,
		Message: message,
		Data:    data,
	}
}

// SuccessResponse 成功响应
func SuccessResponse(data any, message string) *Response {
	return NewResponse(http.StatusOK, true, message, data)
}

func FailedResponse(code int, message string) *Response {
	return NewResponse(code, false, message, nil)
}

// FailedMsgResponse 错误响应，附加信息
func FailedMsgResponse(code int, message string) *Response {
	return FailedResponse(code, message)
}

// FailedErrResponse 错误响应
func FailedErrResponse(code int, err error) *Response {
	return FailedResponse(code, err.Error())
}
