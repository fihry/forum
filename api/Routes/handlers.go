package Routes

import (
	"encoding/json"
	"forum/api/Controllers"
	"forum/api/Models"
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
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")

	// Validate user data
	ok, err := CheckDataForLogin(user)
	if err != nil || !ok {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// Check if user exists
	exists, err := Database.CheckUserExist(user.Username)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Get user from database
	storedUser, err := Database.GetUserByName(user.Username)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Compare the provided password with the stored hashed password
	if !ComparePasswords(storedUser.Password, user.Password) {
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
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    session.SessionKey,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	w.WriteHeader(http.StatusOK)
}

// ComparePasswords compares a hashed password with a plain password
func ComparePasswords(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user Models.User
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")
	user.Email = r.FormValue("email")

	// Check if user already exists
	exist, err := Database.CheckUserExist(user.Username)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if exist {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	// Validate user data
	ok, err := CheckDataForRegister(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Create new user
	err = Database.CreateUser(user)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set session cookie
	user, err = Database.NewSession(user)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: user.SessionKey,
	})
	w.WriteHeader(http.StatusCreated)
}

// those functions are not required for the task
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// id := r.FormValue("id")
	// fmt.Println(id)
	// Update user here
	w.Write([]byte("this api end point is not implemented yet"))
	w.WriteHeader(http.StatusOK)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// id := r.FormValue("id")
	// fmt.Println(id)
	// Delete user here
	w.Write([]byte("this api end point is not implemented yet"))
	w.WriteHeader(http.StatusNoContent)
}

//  Posts handlers ==================================

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	// Get posts here
	var posts = []Models.Poste{}
	//get all posts from the database
	posts, err := Database.GetAllPosts()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post Models.Poste
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
	user, err := Database.GetUserBySession(session.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	post.Author = user.Username
	// Validate post data
	ok, err := CheckDataForPost(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	// Create new post here and return the created post
	id, err := Database.CreatePoste(post)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	post.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
	w.WriteHeader(http.StatusCreated)
}

// those functions are not required for the task
func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// id := r.FormValue("id")
	// fmt.Println(id)
	// Update post here
	w.Write([]byte("this api end point is not implemented yet"))
	w.WriteHeader(http.StatusOK)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// id := r.FormValue("id")
	// fmt.Println(id)
	w.Write([]byte("this api end point is not implemented yet"))
	w.WriteHeader(http.StatusNoContent)
}
