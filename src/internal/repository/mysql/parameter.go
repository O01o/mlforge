package rm

import (
	"context"
	m "mlforge/internal/model"
	qm "mlforge/internal/query/mysql"

	"github.com/jmoiron/sqlx"
)

type parameterRepository struct {
	ctx context.Context
	tx  *sqlx.Tx
}

func NewParameterRepository(ctx context.Context, tx *sqlx.Tx) *parameterRepository {
	return &parameterRepository{ctx: ctx, tx: tx}
}

func (r *parameterRepository) CreateParameters(parameters []m.CreateParameter) error {
	_, err := r.tx.ExecContext(
		r.ctx,
		qm.InsertParametersQuery,
		parameters,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *parameterRepository) GetParameters(
	experimentID uint64,
	runID uint64,
) ([]m.Parameter, error) {
	var parameters []m.Parameter
	err := r.tx.SelectContext(
		r.ctx,
		&parameters,
		qm.GetParametersQuery,
		experimentID,
		runID,
	)
	if err != nil {
		return nil, err
	}
	return parameters, nil
}

func (r *parameterRepository) GetParameterIDs(experimentID uint64, runID uint64) ([]float64, error) {
	var parameterIDs []float64
	err := r.tx.SelectContext(
		r.ctx,
		&parameterIDs,
		qm.GetParameterIDsQuery,
		experimentID,
		runID,
	)
	if err != nil {
		return nil, err
	}
	return parameterIDs, nil
}
