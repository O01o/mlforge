package rm

import (
	"context"
	m "mlforge/internal/model"
	qm "mlforge/internal/query/mysql"

	"github.com/jmoiron/sqlx"
)

type runRepository struct {
	ctx context.Context
	tx  *sqlx.Tx
}

func NewRunRepository(ctx context.Context, tx *sqlx.Tx) *runRepository {
	return &runRepository{ctx: ctx, tx: tx}
}

func (r *runRepository) CreateRun(run m.CreateRun) (uint64, error) {
	res, err := r.tx.NamedExecContext(r.ctx, qm.InsertRunQuery, run)
	if err != nil {
		return 0, err
	}
	runID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(runID), nil
}

func (r *runRepository) GetRuns(experimentID uint64) ([]*m.Run, error) {
	var runs []*m.Run
	err := r.tx.SelectContext(
		r.ctx,
		&runs,
		qm.GetRunsQuery,
		experimentID,
	)
	if err != nil {
		return nil, err
	}
	return runs, nil
}

func (r *runRepository) FinishRunByID(runID uint64) error {
	_, err := r.tx.ExecContext(
		r.ctx,
		qm.FinishRunByIDQuery,
		runID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *runRepository) FailRunByID(runID uint64) error {
	_, err := r.tx.ExecContext(
		r.ctx,
		qm.FailRunByIDQuery,
		runID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *runRepository) DeleteRunByID(runID uint64) error {
	_, err := r.tx.ExecContext(
		r.ctx,
		qm.DeleteRunByIDQuery,
		runID,
	)
	if err != nil {
		return err
	}
	return nil
}
