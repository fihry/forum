package src

import (
	"net/http"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.NotFound(w, r)
	}
	RenderTemplate(w, "login", nil)
}
func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.NotFound(w, r)
	} else {
		RenderTemplate(w, "register", nil)
	}
}
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	} else {
		RenderTemplate(w, "Rout", nil)
	}
}
