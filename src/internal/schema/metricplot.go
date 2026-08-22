package sc

type Plot struct {
	Step  int     `json:"step"`
	Value float64 `json:"value"`
}

type CreatePlotsRequest struct {
	Plots []Plot `json:"plots"`
}

type PlotRange struct {
	StartStep int `json:"startStep"`
	EndStep   int `json:"endStep"`
}

type RunMetricPlots struct {
	RunID    uint64 `json:"runId"`
	MetricID uint64 `json:"metricId"`
	Plots    []Plot `json:"plots"`
}

type GetPlotsRequest struct {
	PlotRange PlotRange `json:"plotRange"`
	RunIDs    []uint64  `json:"runIds"`
}

type GetPlotsResponse struct {
	RunMetricPlots []RunMetricPlots `json:"runMetricPlots"`
}
