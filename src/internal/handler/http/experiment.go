package hh

import (
	"encoding/json"
	"mlforge/internal/core"
	ri "mlforge/internal/repository/interface"
	sc "mlforge/internal/schema"
	"net/http"
	"path"
	"strconv"
)

type experimentHandler struct {
	Log  *core.Log
	repo ri.ExperimentRepository
}

func NewExperimentHandler(repo ri.ExperimentRepository) *experimentHandler {
	return &experimentHandler{
		Log:  core.NewLog(),
		repo: repo,
	}
}

func (h *experimentHandler) CreateExperiment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req *sc.ExperimentCreate
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resr, err := h.repo.CreateExperiment(req.ExperimentName, req.ExperimentDescription)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ressc, err := json.Marshal(&sc.ExperimentByID{
		ExperimentID: resr,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(ressc)
}

func (h *experimentHandler) GetExperimentByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := path.Base(r.URL.Path)
	id, err := strconv.ParseUint(req, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resr, err := h.repo.GetExperimentByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ressc, err := json.Marshal(&sc.Experiment{
		ExperimentID:          resr.ExperimentID,
		ExperimentName:        resr.ExperimentName,
		ExperimentDescription: resr.ExperimentDescription,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(ressc)
}

func (h *experimentHandler) GetExperimentList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	resr, err := h.repo.GetExperimentList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resscd := make([]*sc.Experiment, len(resr))
	for i, experiment := range resr {
		resscd[i] = &sc.Experiment{
			ExperimentID:          experiment.ExperimentID,
			ExperimentName:        experiment.ExperimentName,
			ExperimentDescription: experiment.ExperimentDescription,
		}
	}
	resscj, err := json.Marshal(resscd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resscj)
}

func (h *experimentHandler) DeleteExperimentByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := path.Base(r.URL.Path)
	id, err := strconv.ParseUint(req, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.repo.DeleteExperimentByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
