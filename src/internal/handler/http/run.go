package hh

import (
	"encoding/json"
	"mlforge/internal/core"
	sc "mlforge/internal/schema"
	sei "mlforge/internal/service/interface"
	"net/http"
	"path"
	"strconv"
	"strings"
)

type runHandler struct {
	Log *core.Log
	s   sei.RunService
}

func NewRunHandler(s sei.RunService) *runHandler {
	return &runHandler{
		Log: core.NewLog(),
		s:   s,
	}
}

func (h *runHandler) CreateRun(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	experimentID, err := strconv.ParseUint(pathParts[len(pathParts)-2], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req *sc.CreateRunRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resr, resm, err := h.s.CreateRun(experimentID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ressc, err := json.MarshalIndent(&sc.CreateRunResponse{
		RunID:     resr,
		MetricIDs: resm,
	}, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(ressc)
}

func (h *runHandler) FinishRunByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := path.Base(r.URL.Path)
	id, err := strconv.ParseUint(req, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.s.FinishRunByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *runHandler) FailRunByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := path.Base(r.URL.Path)
	id, err := strconv.ParseUint(req, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.s.FailRunByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *runHandler) DeleteRunByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := path.Base(r.URL.Path)
	id, err := strconv.ParseUint(req, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.s.DeleteRunByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
