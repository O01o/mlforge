package sc

type ExperimentByID struct {
	ExperimentID uint64 `json:"experiment_id"`
}

type ExperimentCreate struct {
	ExperimentName        string `json:"name"`
	ExperimentDescription string `json:"description"`
}

type Experiment struct {
	ExperimentID          uint64 `json:"id"`
	ExperimentName        string `json:"name"`
	ExperimentDescription string `json:"description"`
}
