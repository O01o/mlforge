package main

import (
	"mlforge/internal/core"
	hh "mlforge/internal/handler/http"
	rm "mlforge/internal/repository/mysql"
	r "mlforge/internal/router"
)

func main() {
	db := core.ConnectDB()
	hhs := hh.NewSetupHandler(rm.NewSetupRepository(db))
	if err := hhs.SetupDB(); err != nil {
		panic(err)
	}
	server := r.NewHTTPServer(":8080", db)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
