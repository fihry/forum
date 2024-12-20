package controllers

import (
	"fmt"

	"forum/models"
)

func CheckUserExist(username string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE username = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return false, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(username).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to execute query: %w", err)
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func CheckEmailExist(email string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE email = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return false, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(email).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to execute query: %w", err)
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func GetUserByName(username string) (models.User, error) {
	user := models.User{}
	query := "SELECT id,username,password,email FROM users WHERE username = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return user, fmt.Errorf("failed to prepare statement: %w", err)
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
	query := "SELECT * FROM users WHERE id = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	stmt.QueryRow(id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email)
	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	query := "INSERT INTO users(username, password, email) VALUES(?, ?, ?)"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.Password, user.Email)
	if err != nil {
		return user, fmt.Errorf("failed to execute statement: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user, fmt.Errorf("failed to get last insert id: %w", err)
	}
	user.ID = int(id)
	return user, nil
}

func GetUserBySession(sessionKey string) (models.User, error) {
	user := models.User{}
	query := "SELECT * FROM users WHERE session_key = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return user, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(sessionKey).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.SessionKey,
		&user.ExpireDate)
	if err != nil {
		return user, fmt.Errorf("failed to execute query: %w", err)
	}

	return user, nil
}

func UpdateUser(user models.User) error {
	query := "UPDATE users SET username = ?, password = ?, email = ? WHERE id = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Password, user.Email, user.ID)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}
	return nil
}
