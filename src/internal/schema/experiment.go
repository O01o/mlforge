package sc

type ExperimentByID struct {
	ExperimentID uint64 `json:"experiment_id"`
}

type ExperimentCreate struct {
	ExperimentName        string `json:"experiment_name"`
	ExperimentDescription string `json:"experiment_description"`
}

type Experiment struct {
	ExperimentID          uint64 `json:"experiment_id"`
	ExperimentName        string `json:"experiment_name"`
	ExperimentDescription string `json:"experiment_description"`
}
