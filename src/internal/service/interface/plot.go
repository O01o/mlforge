package sei

import sc "mlforge/internal/schema"

type PlotService interface {
	CreatePlots(metricId uint64, req *sc.CreatePlotsRequest) error
	GetPlots(req *sc.GetPlotsRequest) (*sc.GetPlotsResponse, error)
}
