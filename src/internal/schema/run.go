package sc

import "time"

type CreateRunRequest struct {
	RunName        string            `json:"name"`
	RunDescription string            `json:"description"`
	Parameters     map[string]string `json:"parameters"`
	Metrics        []string          `json:"metrics"`
}

type CreateRunResponse struct {
	RunID     uint64   `json:"id"`
	MetricIDs []uint64 `json:"metricIds"`
}

type RunSummary struct {
	RunID            uint64    `json:"id"`
	RunName          string    `json:"name"`
	RunDescription   string    `json:"description"`
	RunFailureReason string    `json:"failureReason"`
	RunStatus        uint16    `json:"status"`
	StartedAt        time.Time `json:"startedAt"`
	EndedAt          time.Time `json:"endedAt"`
}

type RunDetails struct {
	RunSummary RunSummary  `json:"runSummary"`
	Parameters []Parameter `json:"parameters"`
	Metrics    []Metric    `json:"metrics"`
}

type GetRunResponse struct {
	RunDetails RunDetails `json:"runDetails"`
}
