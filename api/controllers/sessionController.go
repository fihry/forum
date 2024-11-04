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

func init() {
	var err error
	Database, err = sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
}

func NewSession(user models.User) (models.User, error) {
	expireDate := time.Now().Add(24 * time.Hour)
	UUID, err := uuid.NewV4()
	if err != nil {
		return user, fmt.Errorf("failed to generate UUID: %w", err)
	}
	user.SessionKey = UUID.String()
	user.ExpireDate = expireDate
	stmt, err := Database.Prepare("INSERT OR IGNORE INTO sessions (username, key, ExpireDate) VALUES (?, ?, ?)")
	if err != nil {
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.SessionKey, expireDate)
	if err != nil {
		return user, fmt.Errorf("failed to execute statement: %w", err)
	}
	return user, nil
}

func UpdateSessionByUser(user models.User) (models.User, error) {
	expireDate := time.Now().Add(24 * time.Hour)
	UUID, err := uuid.NewV4()
	if err != nil {
		return user, fmt.Errorf("failed to generate UUID: %w", err)
	}
	user.SessionKey = UUID.String()
	user.ExpireDate = expireDate
	stmt, err := Database.Prepare("UPDATE sessions SET key = ?, ExpireDate = ? WHERE username = ?")
	if err != nil {
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.SessionKey, expireDate, user.Username)
	if err != nil {
		return user, fmt.Errorf("failed to execute statement: %w", err)
	}
	log.Println("Session updated", user.SessionKey, user.ExpireDate, user.Username)
	return user, nil
}

func GetSession(Key string) (models.User, error) {
	user := models.User{}

	stmt, err := Database.Prepare("SELECT * FROM users WHERE session = ?")
	if err != nil {
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(Key).Scan(&user.ID, &user.SessionKey, &user.ExpireDate)
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
