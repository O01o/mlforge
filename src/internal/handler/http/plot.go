package hh

import (
	"encoding/json"
	"mlforge/internal/core"
	sc "mlforge/internal/schema"
	sei "mlforge/internal/service/interface"
	"net/http"
	"strconv"
	"strings"
)

type plotHandler struct {
	Log *core.Log
	s   sei.PlotService
}

func NewPlotHandler(s sei.PlotService) *plotHandler {
	return &plotHandler{
		Log: core.NewLog(),
		s:   s,
	}
}

func (h *plotHandler) CreatePlots(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	id, err := strconv.ParseUint(pathParts[len(pathParts)-2], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var req *sc.CreatePlotsRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.s.CreatePlots(id, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *plotHandler) GetPlots(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req *sc.GetPlotsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resr, err := h.s.GetPlots(req)
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
