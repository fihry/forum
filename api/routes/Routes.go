package routes

import (
	"net/http"

	"forum/api/handlers"
)

// ApiRoutes function returns a http.Handler that handles all the api routes for the application

func ApiRoutes() http.Handler {
	// create a new ServeMux
	mux := http.NewServeMux()
	// ============== user routes
	mux.Handle("/api/login", http.HandlerFunc(handlers.LoginHandler))
	mux.Handle("/api/register", http.HandlerFunc(handlers.RegisterHandler))

	// ============== post routes
	mux.Handle("/api/posts", http.HandlerFunc(handlers.PostsHandler))
	mux.Handle("/api/createPost", http.HandlerFunc(handlers.CreatePostHandler))
	mux.Handle("/api/postReaction", http.HandlerFunc(handlers.PostReactionHandler))
	return mux
}

func InitRouter() http.Handler {
	// create a new ServeMux
	router := http.NewServeMux()

	// handle api routes
	router.Handle("/api/", ApiRoutes())

	// handle frontend routes
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/css"))))
	router.HandleFunc("/login", handlers.LoginPageHandler)
	router.HandleFunc("/register", handlers.RegisterPageHandler)
	router.HandleFunc("/", handlers.HomePageHandler)

	return router
}
