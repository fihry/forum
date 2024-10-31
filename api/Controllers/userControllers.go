package Controllers

import (
	"database/sql"
	"fmt"
	"forum/api/Models"
)

type Database struct {
	DB *sql.DB
}

func (db *Database) CheckUserExist(username string) (bool, error) {
	if db.DB == nil {
		return false, fmt.Errorf("database connection is not initialized")
	}

	var count int
	stmt, err := db.DB.Prepare("SELECT COUNT(*) FROM users WHERE username = ?")
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

func (db *Database) GetUserByName(username string) (Models.User, error) {
	user := Models.User{}
	stmt, err := db.DB.Prepare("SELECT * FROM users WHERE username = ?")
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

func (db *Database) GetUserById(id int) (Models.User, error) {
	user := Models.User{}
	stmt, err := db.DB.Prepare("SELECT * FROM users WHERE id = ?")
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

func (db *Database) CreateUser(user Models.User) error {
	_, err := db.DB.Exec("INSERT INTO users (username, password, email, session) VALUES ( ?, ?, ?, ?)",
		user.Username,
		user.Password,
		user.Email,
		user.SessionKey)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) GetUserBySession(session string) (Models.User, error) {
	user := Models.User{}
	stmt, err := db.DB.Prepare("SELECT * FROM users WHERE session = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	stmt.QueryRow(session).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email)
	return user, nil
}

func (db *Database) UpdateUser(user Models.User) error {
	_, err := db.DB.Exec("UPDATE users SET username = ?, password = ?, email = ? WHERE id = ?",
		user.Username,
		user.Password,
		user.Email,
		user.ID)
	if err != nil {
		return err
	}
	return nil
}
