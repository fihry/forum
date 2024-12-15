package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"forum/api/controllers"
	"forum/models"
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

	var n int64
	if engagement.LikeAction == "add" {
		err = controllers.UpdateLikeToEngagement(engagement.PosteID, user.ID)
		if err != nil {
			log.Println(err)
		}
		n, err = controllers.LikePost(engagement.PosteID)
		if err != nil {
			log.Println(err)
		}
	} else if engagement.LikeAction == "remove" {
		err = controllers.RemLikeFromEngagement(engagement.PosteID, user.ID)
		if err != nil {
			log.Println(err)
		}
		n, err = controllers.RemoveLike(engagement.PosteID)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil || n == 0 {
		w.WriteHeader(http.StatusBadRequest)
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
	if engagement.DislikeAction == "add" {
		err = controllers.UpdateDislikeToEngagement(engagement.PosteID, user.ID)
		if err != nil {
			fmt.Println("updatedislike function return:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		n, err = controllers.DislikePost(engagement.PosteID)
	} else if engagement.DislikeAction == "remove" {
		err = controllers.RemDislikeFromEngagement(engagement.PosteID, user.ID)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		n, err = controllers.RemoveDislike(engagement.PosteID)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil || n == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
