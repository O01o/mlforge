package ri

import m "mlforge/internal/model"

type ParameterRepository interface {
	CreateParameters(parameters []m.CreateParameter) error
	GetParameters(experimentID uint64, runID uint64) ([]m.Parameter, error)
	GetParameterIDs(experimentID uint64, runID uint64) ([]uint64, error)
}
