package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"

	"forum/api/models"
)


func NewSession(user models.User) (models.User, error) {
	// create a new session key and set the expiration date to 24 hours from now
	expireDate := time.Now().Add(24 * time.Hour)
	UUID, err := uuid.NewV4()
	if err != nil {
		log.Printf("Failed to generate UUID: %v", err)
		return user, fmt.Errorf("failed to generate UUID: %w", err)
	}
	user.SessionKey = UUID.String()
	user.ExpireDate = expireDate
	// update the user's session key and expiration date in the database
	query := `UPDATE users SET session_key = ?, expire_date = ? WHERE id = ?`
	stmt, err := Database.Prepare(query)
	if err != nil {
		log.Printf("Failed to prepare statement: %v", err)
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.SessionKey, user.ExpireDate, user.ID)
	if err != nil {
		log.Printf("Failed to execute statement: %v", err)
		return user, fmt.Errorf("failed to execute statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Failed to get rows affected: %v", err)
		return user, fmt.Errorf("failed to get rows affected: %w", err)
	}

	log.Printf("Rows affected: %d", rowsAffected)

	if rowsAffected == 0 {
		log.Printf("No rows were updated for user ID: %d", user.ID)
		return user, fmt.Errorf("no rows were updated")
	}

	log.Printf("Session created successfully for user ID: %d", user.ID)
	return user, nil
}

func DeleteSession(Key string) error {
	stmt, err := Database.Prepare("UPDATE  users SET session_key = NULL, expire_date = NULL WHERE session_key = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(Key)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}
	return nil
}

func CheckUserState(username string) error {
	var id int
	var sessionKey sql.NullString
	var expireDate sql.NullTime

	err := Database.QueryRow("SELECT id, session_key, expire_date FROM users WHERE username = ?", username).Scan(&id, &sessionKey, &expireDate)
	if err != nil {
		return fmt.Errorf("failed to query user state: %w", err)
	}

	log.Printf("User state - ID: %d, SessionKey: %v, ExpireDate: %v", id, sessionKey, expireDate)
	return nil
}
