package forum

import (
	"forum/api/routes"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()
	// handle api routes
	mux.Handle("/api/", routes.ApiRoutes())
	// handle frontend routes
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/css"))))
	mux.HandleFunc("/login", LoginPageHandler)
	mux.HandleFunc("/register", RegisterPageHandler)
	mux.HandleFunc("/", HomePageHandler)
	return mux
}
