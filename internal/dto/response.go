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

// ErrorResponse 错误响应
func ErrorResponse(message string) *Response {
	return NewResponse(http.StatusInternalServerError, false, message, nil)
}
