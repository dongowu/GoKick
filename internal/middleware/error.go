package middleware

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/dongowu/gokick/internal/pkg/apperror"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var log = zap.L()

// ErrorHandler 统一错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		// 取最后一个错误
		err := c.Errors.Last()
		var appErr *apperror.AppError
		var requestID string

		if rid, exists := c.Get("X-Request-ID"); exists {
			if id, ok := rid.(string); ok {
				requestID = id
			}
		}

		// 检查是否为 AppError
		if appErr, ok := err.(*apperror.AppError); ok {
			// 业务错误，返回结构化 JSON
			if requestID != "" {
				appErr.RequestID = requestID
			}

			// 记录日志（仅非 4xx 错误或详细信息）
			if appErr.HTTPStatus >= 500 {
				log.Errorw("internal server error",
					"code", appErr.Code,
					"message", appErr.Message,
					"details", appErr.Details,
					"request_id", requestID,
					"path", c.Request.URL.Path,
					"method", c.Request.Method,
				)
			} else {
				log.Infow("client error",
					"code", appErr.Code,
					"message", appErr.Message,
					"request_id", requestID,
					"path", c.Request.URL.Path,
					"method", c.Request.Method,
				)
			}

			c.JSON(appErr.HTTPStatus, gin.H{
				"code":       appErr.Code,
				"message":    appErr.Message,
				"request_id": requestID,
				"details":    appErr.Details,
			})
		} else {
			// 未知错误，脱敏后返回 500
			log.Errorw("unhandled error",
				"error", err.Error(),
				"request_id", requestID,
				"path", c.Request.URL.Path,
				"method", c.Request.Method,
			)

			c.JSON(http.StatusInternalServerError, gin.H{
				"code":       apperror.ErrInternal.Code,
				"message":    apperror.ErrInternal.Message,
				"request_id": requestID,
			})
		}

		c.Abort()
	}
}

// Recovery 恢复 panic 并转为 AppError
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var err error
				switch v := r.(type) {
				case string:
					err = fmt.Errorf("%s", v)
				case error:
					err = v
				default:
					err = fmt.Errorf("%v", v)
				}

				log.Errorw("panic recovered",
					"error", err,
					"request_id", c.GetString("X-Request-ID"),
					"path", c.Request.URL.Path,
					"method", c.Request.Method,
					"stack", getStack(),
				)

				c.Error(apperror.Wrap(err, apperror.ErrInternal.Code, "服务器内部错误"))
				c.Abort()
			}
		}()
		c.Next()
	}
}

func getStack() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}
