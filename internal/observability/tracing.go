package observability

import (
	"context"
	"github.com/gin-gonic/gin"
)

// Config 可观测性配置
type Config struct {
	Enabled         bool
	ServiceName     string
	TracingEndpoint string
	SampleRatio     float64
}

// Tracer 追踪器封装
type Tracer struct{}

// NewTracer 创建新的追踪器
func NewTracer(cfg Config, logger interface{}) (*Tracer, error) {
	return &Tracer{}, nil
}

// Start 开始一个 span
func (t *Tracer) Start(ctx context.Context, name string, opts ...interface{}) (context.Context, interface{}) {
	return ctx, nil
}

// Shutdown 优雅关闭
func (t *Tracer) Shutdown(ctx context.Context) error {
	return nil
}

// GinMiddleware 为 Gin 创建追踪中间件
func GinMiddleware(tracer *Tracer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
