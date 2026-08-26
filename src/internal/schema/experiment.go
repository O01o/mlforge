package sc

import "time"

type CreateExperimentRequest struct {
	ExperimentName        string `json:"name"`
	ExperimentDescription string `json:"description"`
}

type CreateExperimentResponse struct {
	ExperimentID uint64 `json:"id"`
}

type ExperimentSummary struct {
	ExperimentID          uint64    `json:"id"`
	ExperimentName        string    `json:"name"`
	ExperimentDescription string    `json:"description"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}

type ExperimentDetails struct {
	ExperimentSummary ExperimentSummary `json:"experimentSummary"`
	RunDetails        []RunDetails      `json:"runDetails"`
}

type GetExperimentResponse struct {
	ExperimentDetails *ExperimentDetails `json:"experimentDetails"`
}

type GetExperimentSummariesResponse struct {
	ExperimentSummaries []*ExperimentSummary `json:"experimentSummaries"`
}
