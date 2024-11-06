package src

import (
	"forum/api/Routes"
	"net/http"
)

// Routes function returns a http.Handler that handles all the routes for the application
func Routs() http.Handler {
	// create a new ServeMux
	mux := http.NewServeMux()
	// handle api routes
	mux.Handle("/api/", Routes.ApiRoutes())
	// handle frontend routes
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/css"))))
	mux.HandleFunc("/login", LoginPageHandler)
	mux.HandleFunc("/register", RegisterPageHandler)
	mux.HandleFunc("/", HomePageHandler)
	return mux
}
