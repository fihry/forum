package Controllers

import (
	"fmt"
	"forum/api/Models"
	"log"

	"github.com/gofrs/uuid"
)

func (db *Database) NewSession(user Models.User) (Models.User, error) {
	UUID, err := uuid.NewV4()
	if err != nil {
		return user, fmt.Errorf("failed to generate UUID: %w", err)
	}
	user.SessionKey = UUID.String()

	stmt, err := db.DB.Prepare("INSERT OR IGNORE INTO users (session) VALUES (?)")
	if err != nil {
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.SessionKey)
	if err != nil {
		fmt.Println(err)
		return user, fmt.Errorf("failed to execute statement: %w", err)
	}
	log.Println("New session created", user.SessionKey)
	return user, nil
}

func (db *Database) GetSession(Key string) (Models.User, error) {
	user := Models.User{}

	stmt, err := db.DB.Prepare("SELECT * FROM users WHERE session = ?")
	if err != nil {
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(Key).Scan(&user.ID, &user.SessionKey)
	if err != nil {
		return user, fmt.Errorf("failed to query row: %w", err)
	}
	return db.GetUserById(user.ID)
}

func (db *Database) DeleteSession(Key string) error {
	stmt, err := db.DB.Prepare("DELETE FROM users WHERE session = ?")
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
