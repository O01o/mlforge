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
	_, err := r.tx.NamedExecContext(
		r.ctx,
		qm.InsertParametersQuery,
		parameters,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *parameterRepository) GetParameters(runID uint64) ([]m.Parameter, error) {
	var parameters []m.Parameter
	err := r.tx.SelectContext(
		r.ctx,
		&parameters,
		qm.GetParametersQuery,
		runID,
	)
	if err != nil {
		return nil, err
	}
	return parameters, nil
}

func (r *parameterRepository) GetParameterIDs(runID uint64) ([]uint64, error) {
	var parameterIDs []uint64
	err := r.tx.SelectContext(
		r.ctx,
		&parameterIDs,
		qm.GetParameterIDsQuery,
		runID,
	)
	if err != nil {
		return nil, err
	}
	return parameterIDs, nil
}
