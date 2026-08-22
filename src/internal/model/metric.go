package m

type Metric struct {
	MetricID   uint64 `db:"metric_id"`
	RunID      uint64 `db:"run_id"`
	MetricName string `db:"name"`
}

type CreateMetric struct {
	RunID      uint64 `db:"run_id"`
	MetricName string `db:"name"`
}
