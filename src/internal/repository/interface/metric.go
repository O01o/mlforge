package ri

import m "mlforge/internal/model"

type MetricRepository interface {
	CreateMetrics(metrics []m.CreateMetric) ([]uint64, error)
	GetMetrics(runID uint64) ([]m.Metric, error)
	GetMetricIDs(runID uint64) ([]uint64, error)
}
