package ri

import (
	m "mlforge/internal/model"
)

type RunRepository interface {
	CreateRun(run *m.CreateRun) (uint64, error)
	GetRunsByExperimentID(experimentID uint64) ([]*m.Run, error)
	FinishRunByID(experimentID uint64, runID uint64) error
	FailRunByID(experimentID uint64, runID uint64) error
	DeleteRunByID(experimentID uint64, runID uint64) error
}
