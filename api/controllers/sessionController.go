package controllers

import (
	"fmt"
	"log"

	"forum/api/models"

	"github.com/gofrs/uuid"
)

func NewSession(user models.User) (models.User, error) {
	UUID, err := uuid.NewV4()
	if err != nil {
		return user, fmt.Errorf("failed to generate UUID: %w", err)
	}
	user.SessionKey = UUID.String()

	stmt, err := Database.DB.Prepare("INSERT OR IGNORE INTO users (session) VALUES (?)")
	if err != nil {
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.SessionKey)
	if err != nil {
		return user, fmt.Errorf("failed to execute statement: %w", err)
	}
	log.Println("New session created", user.SessionKey)
	return user, nil
}

func GetSession(Key string) (models.User, error) {
	user := models.User{}

	stmt, err := Database.DB.Prepare("SELECT * FROM users WHERE session = ?")
	if err != nil {
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(Key).Scan(&user.ID, &user.SessionKey)
	if err != nil {
		return user, fmt.Errorf("failed to query row: %w", err)
	}
	return GetUserById(user.ID)
}

func DeleteSession(Key string) error {
	stmt, err := Database.DB.Prepare("DELETE FROM users WHERE session = ?")
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
