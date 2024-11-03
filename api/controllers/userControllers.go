package controllers

import (
	"fmt"

	"forum/api/models"
)

func CheckUserExist(username string) (bool, error) {
	if Database.DB == nil {
		return false, fmt.Errorf("database connection is not initialized")
	}

	var count int
	stmt, err := Database.DB.Prepare("SELECT COUNT(*) FROM users WHERE username = ?")
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

func GetUserByName(username string) (models.User, error) {
	user := models.User{}
	stmt, err := Database.DB.Prepare("SELECT * FROM users WHERE username = ?")
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
	stmt, err := Database.DB.Prepare("SELECT * FROM users WHERE id = ?")
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

// use prepare herz
func CreateUser(user models.User) error {
	_, err := Database.DB.Exec("INSERT INTO users (username, password, email, session) VALUES ( ?, ?, ?, ?)",
		user.Username,
		user.Password,
		user.Email,
		user.SessionKey)
	if err != nil {
		return err
	}
	return nil
}

func GetUserBySession(session string) (models.User, error) {
	user := models.User{}
	stmt, err := Database.DB.Prepare("SELECT * FROM users WHERE session = ?")
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

func UpdateUser(user models.User) error {
	_, err := Database.DB.Exec("UPDATE users SET username = ?, password = ?, email = ? WHERE id = ?",
		user.Username,
		user.Password,
		user.Email,
		user.ID)
	if err != nil {
		return err
	}
	return nil
}
