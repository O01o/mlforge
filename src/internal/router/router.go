package r

import (
	hh "mlforge/internal/handler/http"
	sem "mlforge/internal/service/impl"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func NewHTTPServer(addr string, db *sqlx.DB) *http.Server {
	r := mux.NewRouter()

	hhd := hh.NewDocsHandler()
	r.HandleFunc("/docs", hhd.GetDocs).Methods("GET")
	r.PathPrefix("/docs/").Handler(
		http.StripPrefix(
			"/docs/",
			hhd.GetSwaggerHandler(),
		),
	)
	r.HandleFunc("/openapi.yaml", hhd.GetOpenAPI).Methods("GET")

	hhe := hh.NewExperimentHandler(sem.NewExperimentService(db))
	r.HandleFunc("/api/experiments", hhe.CreateExperiment).Methods("POST")
	r.HandleFunc("/api/experiments", hhe.GetExperimentList).Methods("GET")
	r.HandleFunc("/api/experiments/{experimentId:[0-9]+}", hhe.GetExperimentByID).Methods("GET")
	r.HandleFunc("/api/experiments/{experimentId:[0-9]+}", hhe.DeleteExperimentByID).Methods("DELETE")

	hhr := hh.NewRunHandler(sem.NewRunService(db))
	r.HandleFunc("/api/experiments/{experimentId:[0-9]+}/runs", hhr.CreateRun).Methods("POST")
	r.HandleFunc("/api/experiments/{experimentId:[0-9]+}/runs/{runId:[0-9]+}/finish", hhr.FinishRunByID).Methods("POST")
	r.HandleFunc("/api/experiments/{experimentId:[0-9]+}/runs/{runId:[0-9]+}/fail", hhr.FailRunByID).Methods("POST")
	r.HandleFunc("/api/experiments/{experimentId:[0-9]+}/runs/{runId:[0-9]+}", hhr.DeleteRunByID).Methods("DELETE")

	hhp := hh.NewPlotHandler(sem.NewPlotService(db))
	r.HandleFunc("/api/experiments/{experimentId:[0-9]+}/runs/{runId:[0-9]+}/metrics/{metricId:[0-9]+}/plots", hhp.CreatePlots).Methods("POST")
	r.HandleFunc("/api/experiments/{experimentId:[0-9]+}/plots", hhp.GetPlots).Methods("GET")

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
