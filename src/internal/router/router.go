package r

import (
	hh "mlforge/internal/handler/http"
	rm "mlforge/internal/repository/mysql"
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

	hhe := hh.NewExperimentHandler(rm.NewExperimentRepository(db))
	r.HandleFunc("/api/experiments", hhe.CreateExperiment).Methods("POST")
	r.HandleFunc("/api/experiments/{id:[0-9]+}", hhe.GetExperimentByID).Methods("GET")

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
