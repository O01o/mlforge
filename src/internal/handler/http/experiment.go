package hh

import (
	"encoding/json"
	"mlforge/internal/core"
	sc "mlforge/internal/schema"
	sei "mlforge/internal/service/interface"
	"net/http"
	"path"
	"strconv"
)

type experimentHandler struct {
	Log *core.Log
	s   sei.ExperimentService
}

func NewExperimentHandler(s sei.ExperimentService) *experimentHandler {
	return &experimentHandler{
		Log: core.NewLog(),
		s:   s,
	}
}

func (h *experimentHandler) CreateExperiment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req *sc.CreateExperimentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resr, err := h.s.CreateExperiment(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ressc, err := json.MarshalIndent(&sc.CreateExperimentResponse{
		ExperimentID: resr,
	}, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(ressc)
}

func (h *experimentHandler) GetExperimentList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	resr, err := h.s.GetExperimentList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ressc, err := json.MarshalIndent(resr, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
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
	resr, err := h.s.GetExperimentByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ressc, err := json.MarshalIndent(resr, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(ressc)
}

func (h *experimentHandler) DeleteExperimentByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := path.Base(r.URL.Path)
	id, err := strconv.ParseUint(req, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.s.DeleteExperimentByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
