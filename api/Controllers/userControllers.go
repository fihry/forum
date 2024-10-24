package Controllers

import (
	"database/sql"
	"fmt"
	"forum/api/Models"
)

type Database struct {
	DB *sql.DB
}

func (db *Database) AddUserIfNotExist(user Models.User) error {
	_, err := db.DB.Exec("INSERT OR IGNORE INTO users (fullname, username, password, email) VALUES (?, ?, ?, ?)",
		user.Fullname,
		user.Username,
		user.Password,
		user.Email)
	if err != nil {
		return err
	}
	return nil
}
func (db *Database) CheckUserExist(username string) (bool, error) {
	if db.DB == nil {
		return false, fmt.Errorf("database connection is not initialized")
	}

	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM Users WHERE username = ?", username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (db *Database) GetUserByName(username string) (Models.User, error) {
	user := Models.User{}
	err := db.DB.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(
		&user.ID,
		&user.Fullname,
		&user.Username,
		&user.Password,
		&user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *Database) GetUserById(id int) (Models.User, error) {
	user := Models.User{}
	err := db.DB.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(
		&user.ID,
		&user.Fullname,
		&user.Username,
		&user.Password,
		&user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *Database) CreateUser(user Models.User) error {
	_, err := db.DB.Exec("INSERT INTO users (fullname, username, password, email) VALUES (?, ?, ?, ?)",
		user.Fullname,
		user.Username,
		user.Password,
		user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateUser(user Models.User) error {
	_, err := db.DB.Exec("UPDATE users SET fullname = ?, username = ?, password = ?, email = ? WHERE id = ?",
		user.Fullname,
		user.Username,
		user.Password,
		user.Email,
		user.ID)
	if err != nil {
		return err
	}
	return nil
}
