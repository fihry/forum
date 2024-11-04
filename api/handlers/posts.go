package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"forum/api/controllers"
	"forum/api/models"
	"forum/utils"
)

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	// Check if user is authenticated
	session, err := r.Cookie("session")
	if err != nil {
		// Session cookie is not present; return all posts without engagement info
		posts, err := controllers.GetAllPosts()
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(posts)
		return
	}

	// User is authenticated; get user information based on session
	user, err := controllers.GetUserBySession(session.Value)
	log.Println("username by session", user.Username)
	if err != nil {
		// Invalid session; return all posts without engagement info
		posts, err := controllers.GetAllPosts()
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(posts)
		return
	}

	// User is authenticated; get posts with engagement info
	posts, err := controllers.GetAllPostsWithEngagement(user.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Respond with posts including engagement info
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Poste
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Check if user is authenticated
	session, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	user, err := controllers.GetUserBySession(session.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	post.Author = user.Username
	// Validate post data
	ok, err := utils.CheckDataForPost(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	// Create new post here and return the created post
	id, err := controllers.CreatePoste(post)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	post.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
	w.WriteHeader(http.StatusCreated)
}
