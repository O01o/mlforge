package hh

import (
	"io/fs"
	"mlforge/internal/assets"
	"mlforge/internal/core"
	"net/http"
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

func (h *BrowserHandler) GetBrowser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	http.Redirect(w, r, "/browser/", http.StatusTemporaryRedirect)
}

func (h *BrowserHandler) GetWebHandler() http.Handler {
	return http.FileServer(http.FS(h.webFS))
}
