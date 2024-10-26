package Routes

import (
	"encoding/json"
	"fmt"
	"forum/api/Controllers"
	"forum/api/Models"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var Database = Controllers.Database{}

// Users handlers ==================================
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user Models.User
	r.ParseForm()
	// parse query parameters to user struct
	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")
	// Check if user exists
	exists, err := Database.CheckUserExist(user.Username)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Get user from database
	storedUser, err := Database.GetUserByName(user.Username)
	if err != nil {
		fmt.Println("her")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Println(storedUser)

	// Compare the provided password with the stored hashed password
	if !comparePasswords(storedUser.Password, user.Password) {
		fmt.Println("DATA FROM DB", storedUser.Password) // NOTE: This is the hashed password from the database
		fmt.Println("DATA FROM USER", user.Password)     // NOTE: This is the plain password from the user
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Create session
	session, err := Database.NewSession(storedUser)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{Name: "session", Value: session.SessionKey, HttpOnly: true})
	w.WriteHeader(http.StatusOK)
}

// comparePasswords compares a hashed password  with a plain password
func comparePasswords(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
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
