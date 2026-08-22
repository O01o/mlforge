package sei

import sc "mlforge/internal/schema"

type PlotService interface {
	CreatePlots(req *sc.CreatePlotsRequest) (uint64, error)
	GetPlots(req *sc.GetPlotsRequest) (*sc.GetPlotsResponse, error)
}
