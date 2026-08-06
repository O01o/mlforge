package rdm

import (
	m "mlforge/internal/model"
	qm "mlforge/internal/query/mysql"
	ri "mlforge/internal/repository/db/interface"

	"github.com/jmoiron/sqlx"
)

type experimentRepository struct {
	db *sqlx.DB
}

func NewExperimentRepository(db *sqlx.DB) ri.ExperimentRepository {
	return &experimentRepository{db: db}
}

func (r *experimentRepository) CreateExperiment(name string, description string) (uint64, error) {
	res, err := r.db.Exec(qm.ExperimentInsertQuery, name, description)
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
	err := r.db.Get(&experiment, qm.ExperimentGetByIDQuery, id)
	if err != nil {
		return nil, err
	}
	return &experiment, nil
}

func (r *experimentRepository) GetExperimentList() ([]*m.Experiment, error) {
	var experiments []*m.Experiment
	err := r.db.Select(&experiments, qm.ExperimentGetListQuery)
	if err != nil {
		return nil, err
	}
	return experiments, nil
}

func (r *experimentRepository) DeleteExperimentByID(id uint64) error {
	_, err := r.db.Exec(qm.ExperimentDeleteByIDQuery, id)
	if err != nil {
		return err
	}
	return nil
}
