package sc

type CreatePlotsRequest struct {
	Plots map[string]float64 `json:"plots"`
}

type PlotRange struct {
	StartStep uint64 `json:"startStep"`
	EndStep   uint64 `json:"endStep"`
}

type RunMetricPlots struct {
	RunID    uint64             `json:"runId"`
	MetricID uint64             `json:"metricId"`
	Plots    map[string]float64 `json:"plots"`
}

type GetPlotsRequest struct {
	PlotRange PlotRange `json:"plotRange"`
	RunIDs    []uint64  `json:"runIds"`
}

type GetPlotsResponse struct {
	RunMetricPlots []RunMetricPlots `json:"runMetricPlots"`
}
