package handlers

import (
	"fmt"
	"log"
	"net/http"

	"forum/api/controllers"
	"forum/api/models"
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
		fmt.Println(err)
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")
	fmt.Println(user.Username, user.Password)
	// Validate user data
	ok, err := utils.CheckDataForLogin(user)
	if err != nil || !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	// Compare the provided password with the stored hashed password
	if !ComparePasswords(storedUser.Password, user.Password) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Create session
	session, err := controllers.NewSession(user)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    session.SessionKey,
		Expires:  session.ExpireDate,
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
	log.Println("Starting RegisterHandler")

	var user models.User
	err := r.ParseForm()
	if err != nil {
		log.Printf("Failed to parse form: %v", err)
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")
	user.Email = r.FormValue("email")

	log.Printf("Received registration request for username: %s, email: %s", user.Username, user.Email)

	// ... (your existing validation code)

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Create user
	newUser, err := controllers.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("User created successfully. ID: %d, Username: %s", newUser.ID, newUser.Username)

	// Create new session
	sessionUser, err := controllers.NewSession(newUser)
	if err != nil {
		log.Printf("Error creating session: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Session created successfully. SessionKey: %s, ExpireDate: %v", sessionUser.SessionKey, sessionUser.ExpireDate)

	err = controllers.CheckUserState(newUser.Username)
	if err != nil {
		log.Printf("Error checking user state: %v", err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    sessionUser.SessionKey,
		Expires:  sessionUser.ExpireDate,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	w.WriteHeader(http.StatusCreated)
	log.Println("RegisterHandler completed successfully")
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
