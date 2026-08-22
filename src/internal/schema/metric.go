package sc

type CreateMetricRequest struct {
	MetricName string `json:"name"`
}

type Metric struct {
	MetricID   uint64 `json:"id"`
	MetricName string `json:"name"`
}
