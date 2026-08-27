package ri

import m "mlforge/internal/model"

type PlotRepository interface {
	CreatePlots(plots []m.Plot) error
	GetPlots(metricID uint64, startStep uint64, endStep uint64) ([]m.Plot, error)
}
