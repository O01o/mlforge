package m

type Experiment struct {
	ExperimentID          uint64 `db:"experiment_id"`
	ExperimentName        string `db:"experiment_name"`
	ExperimentDescription string `db:"experiment_description"`
}
