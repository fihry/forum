package handlers

import (
	"net/http"

	"forum/utils"
)

func AuthPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth" {
		http.NotFound(w, r)
		return
	}
	utils.RenderTemplate(w, "auth", nil)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	utils.RenderTemplate(w, "index", nil)
}
