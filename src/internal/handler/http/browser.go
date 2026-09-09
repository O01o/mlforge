package hh

import (
	"io/fs"
	"mlforge/internal/assets"
	"mlforge/internal/core"
	"net/http"
	"path"
)

type BrowserHandler struct {
	Log   *core.Log
	webFS fs.FS
}

func NewBrowserHandler() *BrowserHandler {
	webFS, err := fs.Sub(
		assets.Files,
		"web",
	)
	if err != nil {
		panic(err)
	}

	return &BrowserHandler{
		Log:   core.NewLog(),
		webFS: webFS,
	}
}

func (h *BrowserHandler) GetWebHandler() http.Handler {
	fileServer := http.FileServer(http.FS(h.webFS))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if path.Ext(r.URL.Path) == "" {
			indexHTML, err := fs.ReadFile(h.webFS, "index.html")
			if err != nil {
				http.Error(w, "web application is not available", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_, _ = w.Write(indexHTML)
			return
		}
		fileServer.ServeHTTP(w, r)
	})
}
