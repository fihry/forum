package routes

import (
	"encoding/json"
	"fmt"
	"forum/api/Models"
	"net/http"
)

// Users handlers ==================================
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := r.FormValue("username")
	passWord := r.FormValue("password")
	fmt.Println(userName, passWord)
	var user Models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	// Authenticate user here
	w.WriteHeader(http.StatusOK)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user Models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Create new user here
	w.WriteHeader(http.StatusCreated)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	fmt.Println(id)
	// Update user here
	w.WriteHeader(http.StatusOK)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	fmt.Println(id)
	// Delete user here
	w.WriteHeader(http.StatusNoContent)
}

//  Posts handlers ==================================

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	// Get posts here
	posts := []Models.Poste{}
	json.NewEncoder(w).Encode(posts)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post Models.Poste
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Create new post here
	w.WriteHeader(http.StatusCreated)
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	fmt.Println(id)
	// Update post here
	w.WriteHeader(http.StatusOK)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	fmt.Println(id)
	// Delete post here
	w.WriteHeader(http.StatusNoContent)
}
