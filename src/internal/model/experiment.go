package m

type Experiment struct {
	ExperimentID          uint64 `db:"experiment_id"`
	ExperimentName        string `db:"name"`
	ExperimentDescription string `db:"description"`
	CreatedAt             string `db:"created_at"`
	UpdatedAt             string `db:"updated_at"`
}

type CreateExperiment struct {
	ExperimentName        string `db:"name"`
	ExperimentDescription string `db:"description"`
}
