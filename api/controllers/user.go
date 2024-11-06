package controllers

import (
	"fmt"
	"log"

	"forum/api/models"
)

func CheckUserExist(username string) (bool, error) {
	if Database == nil {
		return false, fmt.Errorf("database connection is not initialized")
	}

	var count int
	stmt, err := Database.Prepare("SELECT COUNT(*) FROM users WHERE username = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(username).Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func CheckEmailExist(email string) (bool, error) {
	if Database == nil {
		return false, fmt.Errorf("database connection is not initialized")
	}

	var count int
	stmt, err := Database.Prepare("SELECT COUNT(*) FROM users WHERE email = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(email).Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func GetUserByName(username string) (models.User, error) {
	user := models.User{}
	stmt, err := Database.Prepare("SELECT * FROM users WHERE username = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	stmt.QueryRow(username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email)
	return user, nil
}

func GetUserById(id int) (models.User, error) {
	user := models.User{}
	stmt, err := Database.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	stmt.QueryRow(id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email)
	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	log.Printf("Creating user: %s", user.Username)

	stmt, err := Database.Prepare("INSERT INTO users(username, password, email) VALUES(?, ?, ?)")
	if err != nil {
		log.Printf("Failed to prepare statement: %v", err)
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.Password, user.Email)
	if err != nil {
		log.Printf("Failed to execute statement: %v", err)
		return user, fmt.Errorf("failed to execute statement: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Failed to get last insert id: %v", err)
		return user, fmt.Errorf("failed to get last insert id: %w", err)
	}

	user.ID = int(id)
	log.Printf("User created successfully. ID: %d, Username: %s", user.ID, user.Username)
	return user, nil
}

func GetUserBySession(sessionKey string) (models.User, error) {
	user := models.User{}
	// get the first user with  the given session key
	stmt, err := Database.Prepare("SELECT * FROM users WHERE session_key = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(sessionKey).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.SessionKey,
		&user.ExpireDate)
	if err != nil {
		return user, fmt.Errorf("failed to execute query: %w", err)
	}

	return user, nil
}

func UpdateUser(user models.User) error {
	stmt, err := Database.Prepare("UPDATE users SET username = ?, password = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Password, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}
