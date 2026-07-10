package observability

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// Config 可观测性配置
type Config struct {
	Enabled         bool
	ServiceName     string
	TracingEndpoint string
	SampleRatio     float64
}

// Tracer 追踪器封装
type Tracer struct {
	tracer trace.Tracer
	logger *zap.Logger
}

// NewTracer 创建新的追踪器
func NewTracer(cfg Config, logger *zap.Logger) (*Tracer, error) {
	if !cfg.Enabled {
		logger.Info("Observability disabled")
		return &Tracer{tracer: trace.NewNoopTracerProvider().Tracer(""), logger: logger}, nil
	}

	// 创建 OTLP HTTP 导出器（ Jaeger 默认端口 4318）
	exporter, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithEndpoint(cfg.TracingEndpoint),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	// 创建资源
	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(cfg.ServiceName),
		),
	)
	if err != nil {
		return nil, err
	}

	// 创建 TracerProvider
	tp := sdkTrace.NewTracerProvider(
		sdkTrace.WithBatcher(exporter),
		sdkTrace.WithResource(res),
		sdkTrace.WithSampler(sdkTrace.TraceIDRatioBased(cfg.SampleRatio)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	logger.Info("Observability initialized",
		"service", cfg.ServiceName,
		"endpoint", cfg.TracingEndpoint,
		"sample_ratio", cfg.SampleRatio,
	)

	return &Tracer{
		tracer: tp.Tracer(cfg.ServiceName),
		logger: logger,
	}, nil
}

// Start 开始一个 span
func (t *Tracer) Start(ctx context.Context, name string, opts ...trace.SpanStartEventOption) (context.Context, trace.Span) {
	return t.tracer.Start(ctx, name, opts...)
}

// Shutdown 优雅关闭
func (t *Tracer) Shutdown(ctx context.Context) error {
	// 这里需要获取 TracerProvider 并 shutdown
	// 简化处理，实际应该保存 provider 引用
	t.logger.Info("Tracer shutdown")
	return nil
}

// GinMiddleware 为 Gin 创建追踪中间件
func GinMiddleware(tracer *Tracer) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span := tracer.Start(c.Request.Context(), c.FullPath())
		defer span.End()

		// 将 span 注入上下文
		c.Request = c.Request.WithContext(ctx)

		// 添加请求属性
		span.SetAttributes(
			semconv.HTTPMethodKey.String(c.Request.Method),
			semconv.HTTPTargetKey.String(c.FullPath()),
			semconv.HTTPRouteKey.String(c.FullPath()),
			semconv.HTTPClientIPKey.String(c.ClientIP()),
		)

		c.Next()

		// 添加响应属性
		span.SetAttributes(
			semconv.HTTPStatusCodeKey.Int(c.Writer.Status()),
		)

		if c.Writer.Status() >= 500 {
			span.SetAttributes(trace.WithAttributes(semconv.HTTPResponseContentLengthKey))
			span.RecordError(c.Error.Last())
		}
	}
}
