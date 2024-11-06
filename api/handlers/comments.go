package handlers

import (
	"encoding/json"
	"net/http"

	"forum/api/controllers"
	"forum/api/models"
	"forum/utils"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
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

	comment := models.Comment{}
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	err = utils.CheckDataForComment(comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comment.Author, comment.CreatedAt = user.Username, utils.GetCurrentTime()
	id, err := controllers.CreateComment(comment)
	if err != nil || id == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	comment.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}


func LikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
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
	
	engagement.UserID = user.ID
	var n int64
	if engagement.LikeAction == "add" {
		err = controllers.UpdateLikedComment(engagement.CommentID, user.ID, true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		n, err = controllers.LikeComment(engagement.CommentID)
	} else if engagement.LikeAction == "remove" {
		err = controllers.UpdateLikedComment(engagement.CommentID, user.ID, false)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		n, err = controllers.RemoveLikeComment(engagement.CommentID)
	}else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil || n == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DislikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
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
	engagement.UserID = user.ID
	if engagement.DislikeAction == "add" {
		err = controllers.UpdateDislikedComment(engagement.CommentID, user.ID, true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		n, err = controllers.DislikeComment(engagement.CommentID)
	} else if engagement.DislikeAction == "remove" {
		err = controllers.UpdateDislikedComment(engagement.CommentID, user.ID, false)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		n, err = controllers.RemoveDislikeComment(engagement.CommentID)
	}

	if err != nil || n == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}