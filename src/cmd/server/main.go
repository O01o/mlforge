package main

import r "mlforge/internal/router"

func main() {
	server := r.NewHTTPServer(":8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
