package controllers

import (
	"fmt"
	"time"

	"forum/api/models"

	"github.com/gofrs/uuid"
)

func NewSession(user models.User) (models.User, error) {
	expireDate := time.Now().Add(24 * time.Hour)
	UUID, err := uuid.NewV4()
	if err != nil {
		return user, fmt.Errorf("failed to generate UUID: %w", err)
	}
	user.SessionKey = UUID.String()
	user.ExpireDate = expireDate
	stmt, err := Database.Prepare("INSERT OR IGNORE INTO sessions (userId, key, ExpireDate) VALUES (?, ?, ?)")
	if err != nil {
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.SessionKey, expireDate)
	if err != nil {
		return user, fmt.Errorf("failed to execute statement: %w", err)
	}
	return user, nil
}

func GetSession(Key string) (models.User, error) {
	user := models.User{}

	stmt, err := Database.Prepare("SELECT * FROM users WHERE session = ?")
	if err != nil {
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(Key).Scan(&user.ID, &user.SessionKey ,&user.ExpireDate)
	if err != nil {
		return user, fmt.Errorf("failed to query row: %w", err)
	}
	return GetUserById(user.ID)
}

func DeleteSession(Key string) error {
	stmt, err := Database.Prepare("DELETE FROM users WHERE session = ?")
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
