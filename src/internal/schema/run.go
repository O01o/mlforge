package sc

type CreateRunRequest struct {
	RunName        string                   `json:"name"`
	RunDescription string                   `json:"description"`
	Parameters     []CreateParameterRequest `json:"parameters"`
	Metrics        []CreateMetricRequest    `json:"metrics"`
}

type CreateRunResponse struct {
	RunID     uint64   `json:"id"`
	MetricIDs []uint64 `json:"metricIds"`
}

type RunSummary struct {
	RunID            uint64 `json:"id"`
	RunName          string `json:"name"`
	RunDescription   string `json:"description"`
	RunFailureReason string `json:"failureReason"`
	RunStatus        string `json:"status"`
	StartedAt        string `json:"startedAt"`
	EndedAt          string `json:"endedAt"`
}

type RunDetails struct {
	RunSummary RunSummary  `json:"runSummary"`
	Parameters []Parameter `json:"parameters"`
	Metrics    []Metric    `json:"metrics"`
}

type GetRunResponse struct {
	RunDetails RunDetails `json:"runDetails"`
}
