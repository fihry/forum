package Routes

import "net/http"

// ApiRoutes function returns a http.Handler that handles all the api routes for the application a

func ApiRoutes() http.Handler {
	// create a new ServeMux
	mux := http.NewServeMux()
	// ============== user routes
	mux.Handle("/api/login", http.HandlerFunc(LoginHandler))
	mux.Handle("/api/register", http.HandlerFunc(RegisterHandler))
	mux.Handle("/api/updateUser/{id}", http.HandlerFunc(UpdateUserHandler))
	mux.Handle("/api/deleteUser/{id}", http.HandlerFunc(DeleteUserHandler))
	// ============== post routes
	mux.Handle("/api/posts", http.HandlerFunc(PostsHandler))
	mux.Handle("/api/createPost", http.HandlerFunc(CreatePostHandler))
	mux.Handle("/api/updatePost/{id}", http.HandlerFunc(UpdatePostHandler))
	mux.Handle("/api/deletePost/{id}", http.HandlerFunc(DeletePostHandler))
	return mux
}
