package apperror

import "net/http"

// AppError 统一业务错误结构
type AppError struct {
	Code       int         `json:"code"`               // 业务错误码
	HTTPStatus int         `json:"-"`                  // HTTP 状态码
	Message    string      `json:"message"`            // 用户友好提示
	RequestID  string      `json:"request_id,omitempty"` // 链路追踪 ID
	Details    interface{} `json:"details,omitempty"`  // 额外调试信息
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code=%d message=%s", e.Code, e.Message)
}

// 预定义错误码
var (
	// 4xx 客户端错误
	ErrInvalidParams = &AppError{Code: 1001, HTTPStatus: http.StatusBadRequest, Message: "参数错误"}
	ErrUnauthorized  = &AppError{Code: 1002, HTTPStatus: http.StatusUnauthorized, Message: "未授权"}
	ErrForbidden     = &AppError{Code: 1003, HTTPStatus: http.StatusForbidden, Message: "禁止访问"}
	ErrNotFound      = &AppError{Code: 1004, HTTPStatus: http.StatusNotFound, Message: "资源不存在"}
	ErrConflict      = &AppError{Code: 1005, HTTPStatus: http.StatusConflict, Message: "资源冲突"}
	ErrTooMany       = &AppError{Code: 1006, HTTPStatus: http.StatusTooManyRequests, Message: "请求过于频繁"}

	// 5xx 服务端错误
	ErrInternal = &AppError{Code: 5000, HTTPStatus: http.StatusInternalServerError, Message: "内部错误"}
	ErrDatabase = &AppError{Code: 5001, HTTPStatus: http.StatusInternalServerError, Message: "数据库错误"}
	ErrCache    = &AppError{Code: 5002, HTTPStatus: http.StatusInternalServerError, Message: "缓存错误"}
	ErrMQ       = &AppError{Code: 5003, HTTPStatus: http.StatusInternalServerError, Message: "消息队列错误"}
)

// Wrap 包装普通 error 为 AppError
func Wrap(err error, code int, msg string) *AppError {
	if appErr, ok := err.(*AppError); ok {
		return appErr
	}
	return &AppError{
		Code:       code,
		HTTPStatus: http.StatusInternalServerError,
		Message:    msg,
		Details:    err.Error(),
	}
}

// WithDetails 附加额外信息
func (e *AppError) WithDetails(details interface{}) *AppError {
	e.Details = details
	return e
}

// WithRequestID 设置请求 ID
func (e *AppError) WithRequestID(requestID string) *AppError {
	e.RequestID = requestID
	return e
}

// IsAppError 判断是否为 AppError
func IsAppError(err error) (*AppError, bool) {
	appErr, ok := err.(*AppError)
	return appErr, ok
}
