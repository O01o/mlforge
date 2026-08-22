package sei

import sc "mlforge/internal/schema"

type ExperimentService interface {
	CreateExperiment(req *sc.CreateExperimentRequest) (uint64, error)
	GetExperimentByID(id uint64) (*sc.GetExperimentResponse, error)
	GetExperimentList() ([]*sc.GetExperimentSummariesResponse, error)
	DeleteExperimentByID(id uint64) error
}
