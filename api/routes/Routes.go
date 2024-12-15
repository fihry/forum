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
	mux.Handle("/api/logout", http.HandlerFunc(handlers.LogoutHandler))
	mux.Handle("/api/session_check", http.HandlerFunc(handlers.SessionCheck))

	// ============== post routes
	mux.Handle("/api/posts", http.HandlerFunc(handlers.PostsHandler))
	mux.Handle("/api/filterPosts", http.HandlerFunc(handlers.FilterPostHandler))
	mux.Handle("/api/createPost", http.HandlerFunc(handlers.CreatePostHandler))
	mux.Handle("/api/posts/like", http.HandlerFunc(handlers.LikePostHandler))
	mux.Handle("/api/posts/dislike", http.HandlerFunc(handlers.DislikePostHandler))

	// ============== comment routes
	mux.Handle("/api/createcomment", http.HandlerFunc(handlers.CreateCommentHandler))
	mux.Handle("/api/comments/like", http.HandlerFunc(handlers.LikeCommentHandler))
	mux.Handle("/api/comments/dislike", http.HandlerFunc(handlers.DislikeCommentHandler))

	return mux
}

// InitRouter function returns a http.Handler that handles all the routes for the application
func InitRouter() http.Handler {
	// create a new ServeMux
	router := http.NewServeMux()

	// handle api routes
	router.Handle("/api/", ApiRoutes())

	// handle frontend routes
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./view/assets"))))
	router.HandleFunc("/login", handlers.LoginPageHandler)
	router.HandleFunc("/register", handlers.RegisterPageHandler)
	router.HandleFunc("/", handlers.HomePageHandler)

	return router
}
