package handlers

import (
	"encoding/json"
	"net/http"

	"forum/api/controllers"
	"forum/api/models"
)

// Reaction handlers =================================// Reaction handlers =================================
func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user, err := controllers.GetUserBySession(session.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	engagement := models.Engagement{}
	err = json.NewDecoder(r.Body).Decode(&engagement)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Update post likes and dislikes
	var n int64
	// post := models.Poste{}
	// controllers.GetPostByEngagement(user.ID, post)
	// fmt.Println(post.Liked)
	if engagement.LikeAction == "add" {
		// if post.Liked != nil {
		// 	w.WriteHeader(http.StatusConflict)
		// 	return
		// }
		controllers.AddLikeToEngament(engagement.PosteID, user.ID)
		n, err = controllers.LikePost(engagement.PosteID)
	} else if engagement.LikeAction == "remove" {
		// if post.Liked == nil {
		// 	w.WriteHeader(http.StatusConflict)
		// 	return
		// }
		controllers.RemLikeFromEngagement(engagement.PosteID, user.ID)
		n, err = controllers.RemoveLike(engagement.PosteID)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil || n == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DislikePostHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user, err := controllers.GetUserBySession(session.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	engagement := models.Engagement{}
	err = json.NewDecoder(r.Body).Decode(&engagement)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Update post likes and dislikes
	var n int64
	post := models.Poste{}
	if engagement.DislikeAction == "add" {
		controllers.GetPostByEngagement(user.ID, post)
		if post.Disliked != nil {
			w.WriteHeader(http.StatusConflict)
			return
		}
		controllers.AddDislikeToEngament(engagement.PosteID, user.ID)
		n, err = controllers.LikePost(engagement.PosteID)
	} else if engagement.DislikeAction == "remove" {
		if post.Disliked == nil {
			w.WriteHeader(http.StatusConflict)
			return
		}
		controllers.RemDislikeFromEngagement(engagement.PosteID, user.ID)
		n, err = controllers.RemoveDislike(engagement.PosteID)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil || n == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
