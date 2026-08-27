package ri

import m "mlforge/internal/model"

type ExperimentRepository interface {
	CreateExperiment(experiment *m.CreateExperiment) (uint64, error)
	GetExperimentByID(id uint64) (*m.Experiment, error)
	GetExperimentList() ([]*m.Experiment, error)
	UpdateExperimentByID(id uint64, experiment *m.UpdateExperiment) error
	DeleteExperimentByID(id uint64) error
}
