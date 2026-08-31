package hh

import "net/http"

func RedirectToBrowser(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/browser", http.StatusSeeOther)
}
