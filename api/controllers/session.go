package controllers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"

	"forum/models"
)

func NewSession(user models.User) (models.User, error) {
	// create a new session key and set the expiration date to 24 hours from now
	expireDate := time.Now().Add(24 * time.Hour)
	UUID, err := uuid.NewV4()
	if err != nil {
		return user, fmt.Errorf("failed to generate UUID: %w", err)
	}
	user.SessionKey = UUID.String()
	user.ExpireDate = expireDate
	// update the user's session key and expiration date in the database
	query := `UPDATE users SET session_key = ?, expire_date = ? WHERE id = ?`
	stmt, err := Database.Prepare(query)
	if err != nil {
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.SessionKey, user.ExpireDate, user.ID)
	if err != nil {
		return user, fmt.Errorf("failed to execute statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return user, fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return user, fmt.Errorf("no rows affected for user with ID %d", user.ID)
	}
	return user, nil
}

func DeleteSession(Key string) error {
	query := "UPDATE  users SET session_key = NULL, expire_date = NULL WHERE session_key = ?"
	stmt, err := Database.Prepare(query)
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
	query := "SELECT id, session_key, expire_date FROM users WHERE username = ?"
	stm, err := Database.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stm.Close()
	err = stm.QueryRow(username).Scan(&id, &sessionKey, &expireDate)
	if err != nil {
		return fmt.Errorf("failed to query user state: %w", err)
	}
	if !sessionKey.Valid || !expireDate.Valid {
		return fmt.Errorf("user %s is not logged in", username)
	}
	if expireDate.Time.Before(time.Now()) {
		return fmt.Errorf("session for user %s has expired", username)
	}
	return nil
}
