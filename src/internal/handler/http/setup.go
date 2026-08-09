package hh

import (
	"mlforge/internal/core"
	ri "mlforge/internal/repository/interface"
)

type setupHandler struct {
	Log  *core.Log
	repo ri.SetupRepository
}

func NewSetupHandler(repo ri.SetupRepository) *setupHandler {
	return &setupHandler{
		Log:  core.NewLog(),
		repo: repo,
	}
}

func (h *setupHandler) SetupDB() error {
	return h.repo.Setup()
}
