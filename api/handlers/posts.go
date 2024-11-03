package handlers

import (
	"encoding/json"
	"net/http"

	"forum/api/controllers"
	"forum/api/models"
	"forum/utils"
)

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	// get all posts from the database
	posts, err := controllers.GetAllPosts()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
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

// Reaction handlers ==================================
func PostReactionHandler(w http.ResponseWriter, r *http.Request) {
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
	// check type of reaction
	if post.Reaction == "like" {
		// check if already liked : remove like
		// if not liked :
		// check if already disliked : remove dislike
		// like the post
		_, err = controllers.LikePost(post)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
	if post.Reaction == "dislike" {
		// check if already disliked : remove dislike
		// if not disliked :
		// check if already liked : remove like
		// dislike the post
		_, err = controllers.DislikePost(post)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
