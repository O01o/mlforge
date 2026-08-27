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

func (r *metricRepository) CreateMetrics(metrics []m.CreateMetric) error {
	for _, metric := range metrics {
		_, err := r.tx.NamedExecContext(
			r.ctx,
			qm.InsertMetricsQuery,
			metric,
		)
		if err != nil {
			return err
		}
	}
	return nil
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
