package sei

import sc "mlforge/internal/schema"

type RunService interface {
	CreateRun(req *sc.CreateRunRequest) (uint64, error)
	FinishRunByID(id uint64) error
	FailRunByID(id uint64) error
	DeleteRunByID(id uint64) error
}
