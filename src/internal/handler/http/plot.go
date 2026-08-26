package hh

import (
	"encoding/json"
	"mlforge/internal/core"
	sc "mlforge/internal/schema"
	sei "mlforge/internal/service/interface"
	"net/http"
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

	var req *sc.CreatePlotsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resr, err := h.s.CreatePlots(req)
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

func (h *plotHandler) GetPlots(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Implement the logic to get plots here
	// This is a placeholder implementation
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"plots": []}`))
}
