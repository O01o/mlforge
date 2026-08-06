package ri

import m "mlforge/internal/model"

type ExperimentRepository interface {
	CreateExperiment(name string, description string) (uint64, error)
	GetExperimentByID(id uint64) (*m.Experiment, error)
	GetExperimentList() ([]*m.Experiment, error)
	DeleteExperimentByID(id uint64) error
}
