package m

import "time"

type Experiment struct {
	ExperimentID          uint64    `db:"experiment_id"`
	ExperimentName        string    `db:"name"`
	ExperimentDescription string    `db:"description"`
	CreatedAt             time.Time `db:"created_at"`
	UpdatedAt             time.Time `db:"updated_at"`
}

type CreateExperiment struct {
	ExperimentName        string `db:"name"`
	ExperimentDescription string `db:"description"`
}
