package ri

import (
	m "mlforge/internal/model"
)

type RunRepository interface {
	CreateRun(run *m.CreateRun) (uint64, error)
	GetRuns(experimentID uint64) ([]*m.Run, error)
	FinishRunByID(runID uint64) error
	FailRunByID(runID uint64) error
	DeleteRunByID(runID uint64) error
}
