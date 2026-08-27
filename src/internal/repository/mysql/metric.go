package rm

import (
	"context"
	m "mlforge/internal/model"
	qm "mlforge/internal/query/mysql"

	"github.com/jmoiron/sqlx"
)

type metricRepository struct {
	ctx context.Context
	tx  *sqlx.Tx
}

func NewMetricRepository(ctx context.Context, tx *sqlx.Tx) *metricRepository {
	return &metricRepository{ctx: ctx, tx: tx}
}

func (r *metricRepository) CreateMetrics(metrics []m.CreateMetric) ([]uint64, error) {
	_, err := r.tx.NamedExecContext(
		r.ctx,
		qm.InsertMetricsQuery,
		metrics,
	)
	if err != nil {
		return nil, err
	}
	var metricIDs []uint64
	for _, metric := range metrics {
		metricIDs = append(metricIDs, metric.RunID)
	}
	return metricIDs, nil
}

func (r *metricRepository) GetMetrics(runID uint64) ([]m.Metric, error) {
	var metrics []m.Metric
	err := r.tx.SelectContext(
		r.ctx,
		&metrics,
		qm.GetMetricsQuery,
		runID,
	)
	if err != nil {
		return nil, err
	}
	return metrics, nil
}

func (r *metricRepository) GetMetricIDs(runID uint64) ([]uint64, error) {
	var metricIDs []uint64
	err := r.tx.SelectContext(
		r.ctx,
		&metricIDs,
		qm.GetMetricIDsQuery,
		runID,
	)
	if err != nil {
		return nil, err
	}
	return metricIDs, nil
}
