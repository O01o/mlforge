package rm

import (
	"context"
	m "mlforge/internal/model"
	qm "mlforge/internal/query/mysql"
	ri "mlforge/internal/repository/interface"

	"github.com/jmoiron/sqlx"
)

type experimentRepository struct {
	ctx context.Context
	tx  *sqlx.Tx
}

func NewExperimentRepository(ctx context.Context, tx *sqlx.Tx) ri.ExperimentRepository {
	return &experimentRepository{ctx: ctx, tx: tx}
}

func (r *experimentRepository) CreateExperiment(experiment *m.CreateExperiment) (uint64, error) {
	res, err := r.tx.NamedExecContext(r.ctx, qm.InsertExperimentQuery, experiment)
	if err != nil {
		return 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastID), nil
}

func (r *experimentRepository) GetExperimentByID(id uint64) (*m.Experiment, error) {
	var experiment m.Experiment
	err := r.tx.GetContext(r.ctx, &experiment, qm.GetExperimentByIDQuery, id)
	if err != nil {
		return nil, err
	}
	return &experiment, nil
}

func (r *experimentRepository) GetExperimentList() ([]*m.Experiment, error) {
	var experiments []*m.Experiment
	err := r.tx.SelectContext(r.ctx, &experiments, qm.GetExperimentsQuery)
	if err != nil {
		return nil, err
	}
	return experiments, nil
}

func (r *experimentRepository) DeleteExperimentByID(id uint64) error {
	_, err := r.tx.ExecContext(r.ctx, qm.DeleteExperimentByIDQuery, id)
	if err != nil {
		return err
	}
	return nil
}
