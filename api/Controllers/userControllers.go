package controllers

import (
	"database/sql"
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

func (db *Database) GetUser(id int) (Models.User, error) {
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
