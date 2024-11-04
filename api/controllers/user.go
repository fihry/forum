package controllers

import (
	"fmt"

	"forum/api/models"
)

func CheckUserExist(username string) (bool, error) {
	if Database == nil {
		return false, fmt.Errorf("database connection is not initialized")
	}

	var count int
	stmt, err := Database.Prepare("SELECT COUNT(*) FROM users WHERE username = ?")
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

func CheckEmailExist(email string) (bool, error) {
	if Database == nil {
		return false, fmt.Errorf("database connection is not initialized")
	}

	var count int
	stmt, err := Database.Prepare("SELECT COUNT(*) FROM users WHERE email = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(email).Scan(&count)
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
	stmt, err := Database.Prepare("SELECT * FROM users WHERE username = ?")
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
	stmt, err := Database.Prepare("SELECT * FROM users WHERE id = ?")
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

func CreateUser(user models.User) error {
	stmt, err := Database.Prepare("INSERT INTO users (username, password, email) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func GetUserBySession(sessionKey string) (models.User, error) {
	user := models.User{}
	stmt, err := Database.Prepare("SELECT u.id, u.username, u.password, u.email FROM users u JOIN sessions s ON u.username = s.username WHERE s.key = ?")
	if err != nil {
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(sessionKey).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email)
	if err != nil {
		return user, fmt.Errorf("failed to execute query: %w", err)
	}

	return user, nil
}

func UpdateUser(user models.User) error {
	stmt, err := Database.Prepare("UPDATE users SET username = ?, password = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Password, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}
