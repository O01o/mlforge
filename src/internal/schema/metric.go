package sc

type CreateMetricRequest struct {
	RunID      uint64 `json:"run_id"`
	MetricName string `json:"name"`
}

type Metric struct {
	MetricID   uint64 `json:"id"`
	MetricName string `json:"name"`
}
