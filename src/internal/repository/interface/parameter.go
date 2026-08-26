package ri

import m "mlforge/internal/model"

type ParameterRepository interface {
	CreateParameters(parameters []m.CreateParameter) error
	GetParameters(runID uint64) ([]m.Parameter, error)
	GetParameterIDs(runID uint64) ([]uint64, error)
}
