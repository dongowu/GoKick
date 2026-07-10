package observability

// Metrics 指标
type Metrics struct{}

// NewMetrics 创建指标
func NewMetrics(cfg interface{}) (*Metrics, error) {
	return &Metrics{}, nil
}
