package r

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPServer(addr string) *http.Server {
	r := mux.NewRouter()

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
