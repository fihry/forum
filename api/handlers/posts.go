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
	// get all posts from the database
	posts, err := controllers.GetAllPosts()
	if err != nil {
		log.Println(err)
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

// Reaction handlers =================================
func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	// session, err := r.Cookie("session")
	// if err != nil {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }
	// user, err := controllers.GetUserBySession(session.Value)
	// if err != nil {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }
	engagement := models.Engagement{}
	err := json.NewDecoder(r.Body).Decode(&engagement)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update post likes and dislikes
	var n int64
	if engagement.LikeAction == "add" {
		n, err = controllers.LikePost(engagement.PosteID)
		controllers.AddLikeToEngament(engagement.PosteID, 5)
	} else if engagement.LikeAction == "remove" {
		n, err = controllers.RemoveLike(engagement.PosteID)
		controllers.RemLikeFromEngagement(engagement.PosteID, 5)
	} else {
		http.Error(w, "Invalid action", http.StatusBadRequest)
		return
	}
	if err != nil || n == 0 {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
