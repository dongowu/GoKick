package observability

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// HTTP 请求指标
	HTTPRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gokick_http_requests_total",
			Help: "Total HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	HTTPRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "gokick_http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
		},
		[]string{"method", "path"},
	)

	// 数据库指标
	DBQueriesTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gokick_db_queries_total",
			Help: "Total database queries",
		},
		[]string{"operation", "table"},
	)

	DBQueryDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "gokick_db_query_duration_seconds",
			Help:    "Database query duration in seconds",
			Buckets: []float64{.001, .005, .01, .025, .05, .1, .25, .5, 1},
		},
		[]string{"operation", "table"},
	)

	// Redis 指标
	RedisCommandsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gokick_redis_commands_total",
			Help: "Total Redis commands",
		},
		[]string{"command"},
	)

	RedisCommandDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "gokick_redis_command_duration_seconds",
			Help:    "Redis command duration in seconds",
			Buckets: []float64{.0001, .0005, .001, .005, .01, .025, .05, .1},
		},
		[]string{"command"},
	)

	// 业务指标
	ActiveUsers = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "gokick_active_users",
			Help: "Current active users",
		},
	)
)
