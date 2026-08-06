package r

import (
	hh "mlforge/internal/handler/http"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPServer(addr string) *http.Server {
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

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
