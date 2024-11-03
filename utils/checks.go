package utils

import (
	"errors"
	"regexp"

	"forum/api/models"
)

func CheckDataForRegister(userData models.User) (bool, error) {
	if userData.Username == "" {
		return false, errors.New("username is required")
	}
	if userData.Password == "" {
		return false, errors.New("password is required")
	}
	if userData.Email == "" {
		return false, errors.New("email is required")
	}

	// check if data includes special characters to avoid sql injection
	specialCharRegex := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
	if specialCharRegex.MatchString(userData.Username) {
		return false, errors.New("username cannot contain special characters")
	}
	// Validate username format
	if len(userData.Username) < 4 {
		return false, errors.New("username must contain at least 4 characters")
	}
	// Validate email format
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(userData.Email) {
		return false, errors.New("invalid email format")
	}
	// Validate password format

	if len(userData.Password) < 8 {
		return false, errors.New("password must contain at least 8 characters")
	}

	checkPasswordUppercase := regexp.MustCompile(`[A-Z]`)
	if !checkPasswordUppercase.MatchString(userData.Password) {
		return false, errors.New("password must contain at least one uppercase letter")
	}

	checkPasswordSpecialChar := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
	if !checkPasswordSpecialChar.MatchString(userData.Password) {
		return false, errors.New("password must contain at least one special character")
	}

	checkPasswordNumber := regexp.MustCompile(`[0-9]`)
	if !checkPasswordNumber.MatchString(userData.Password) {
		return false, errors.New("password must contain at least one number")
	}
	checkPasswordLowercase := regexp.MustCompile(`[a-z]`)
	if !checkPasswordLowercase.MatchString(userData.Password) {
		return false, errors.New("password must contain at least one lowercase letter")
	}

	return true, nil
}

func CheckDataForLogin(userData models.User) (bool, error) {
	if userData.Username == "" {
		return false, errors.New("username is required")
	}
	if len(userData.Username) < 4 {
		return false, errors.New("username must contain at least 4 characters")
	}
	if userData.Password == "" {
		return false, errors.New("password is required")
	}
	if len(userData.Password) < 8 {
		return false, errors.New("password must contain at least 8 characters")
	}
	return true, nil
}

func CheckDataForPost(postData models.Poste) (bool, error) {
	if postData.Title == "" {
		return false, errors.New("title is required")
	}
	if postData.Content == "" {
		return false, errors.New("content is required")
	}
	return true, nil
}
