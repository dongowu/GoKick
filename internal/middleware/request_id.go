package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID 请求链路 ID 生成与传递
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取或生成新的 Request ID
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}

		// 设置到上下文
		c.Set("X-Request-ID", requestID)

		// 设置到响应头
		c.Header("X-Request-ID", requestID)

		c.Next()
	}
}

func generateRequestID() string {
	return uuid.NewString()
}
