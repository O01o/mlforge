package ri

import m "mlforge/internal/model"

type MetricRepository interface {
	CreateMetrics(metrics []m.CreateMetric) ([]uint64, error)
	GetMetrics(experimentID uint64, runID uint64) ([]m.Metric, error)
	GetMetricIDs(experimentID uint64, runID uint64) ([]uint64, error)
}
