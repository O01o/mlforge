package hh

import (
	"io/fs"
	"mlforge/internal/assets"
	"mlforge/internal/core"
	"net/http"
)

type DocsHandler struct {
	Log       *core.Log
	swaggerFS fs.FS
}

func NewDocsHandler() *DocsHandler {
	swaggerFS, err := fs.Sub(
		assets.Files,
		"swagger",
	)
	if err != nil {
		panic(err)
	}

	return &DocsHandler{
		Log:       core.NewLog(),
		swaggerFS: swaggerFS,
	}
}

func (h *DocsHandler) GetDocs(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	http.Redirect(w, r, "/docs/", http.StatusTemporaryRedirect)
}

func (h *DocsHandler) GetSwaggerHandler() http.Handler {
	return http.FileServer(http.FS(h.swaggerFS))
}

func (h *DocsHandler) GetOpenAPI(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	spec, err := assets.Files.ReadFile("openapi.yaml")
	if err != nil {
		// h.Log.Error("failed to read OpenAPI specification")

		http.Error(
			w,
			"failed to read OpenAPI specification",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/yaml; charset=utf-8",
	)
	w.Header().Set(
		"Cache-Control",
		"no-cache",
	)

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(spec); err != nil {
		// h.Log.Error("failed to write OpenAPI specification")
	}
}
