package handlers

import (
	"fmt"
	"log"
	"net/http"

	"forum/api/controllers"
	"forum/models"
	"forum/utils"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user models.User
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")

	// Validate user data
	ok, err := utils.CheckDataForLogin(user)
	if err != nil || !ok {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// Check if user exists
	exists, err := controllers.CheckUserExist(user.Username)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Get user from database
	storedUser, err := controllers.GetUserByName(user.Username)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Println("Stored user", storedUser)

	// Compare the provided password with the stored hashed password
	if !ComparePasswords(storedUser.Password, user.Password) {
		log.Println("Invalid password", user.Password, storedUser.Password)
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Create session
	newUser, err := controllers.NewSession(storedUser)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    newUser.SessionKey,
		Expires:  newUser.ExpireDate,
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
	var user models.User
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")
	user.Email = r.FormValue("email")

	// Validate user data
	ok, err := utils.CheckDataForRegister(user)
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
	// Set session cookie

	// Create user
	newUser, err := controllers.CreateUser(user)
	if err != nil {
		log.Println("Error creating user", err)
		log.Println("Error creating user", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Create new session
	newUser, err = controllers.NewSession(newUser)
	if err != nil {
		log.Println("Error creating session", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   newUser.SessionKey,
		Expires: newUser.ExpireDate,
		Secure:  true,
	})
	w.WriteHeader(http.StatusCreated)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
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
	err = controllers.DeleteSession(user.Username)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   "",
		Expires: user.ExpireDate,
		MaxAge:  -1,
		Secure:  true,
	})
	w.WriteHeader(http.StatusOK)
}

func SessionCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	session, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	user, err := controllers.GetUserBySession(session.Value)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.Write([]byte(user.Username))
}
