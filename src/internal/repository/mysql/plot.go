package rm

import (
	"context"
	m "mlforge/internal/model"
	qm "mlforge/internal/query/mysql"

	"github.com/jmoiron/sqlx"
)

type plotRepository struct {
	ctx context.Context
	tx  *sqlx.Tx
}

func NewPlotRepository(ctx context.Context, tx *sqlx.Tx) *plotRepository {
	return &plotRepository{ctx: ctx, tx: tx}
}

func (r *plotRepository) CreatePlots(plots []m.Plot) error {
	for _, plot := range plots {
		_, err := r.tx.NamedExecContext(
			r.ctx,
			qm.InsertPlotsQuery,
			plot,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *plotRepository) GetPlots(metricID uint64, startStep uint64, endStep uint64) ([]m.Plot, error) {
	var plots []m.Plot
	err := r.tx.SelectContext(
		r.ctx,
		&plots,
		qm.GetPlotsByRangeOfStepQuery,
		metricID,
		startStep,
		endStep,
	)
	if err != nil {
		return nil, err
	}
	return plots, nil
}
