package utils

import (
	"errors"
	"os"
	"strconv"
)

// GetPort gets the port from the command line arguments
func GetPort() (string, error) {
	args := os.Args[1:]
	DefaultPort := ":8080"
	if len(args) == 0 {
		return DefaultPort, errors.New("\033[33mno port provided,forum using default port 8080\033[0m")
	} else if len(args) != 1 || !validPort(args[0]) {
		return DefaultPort, errors.New("invalid port :" + args[0] + ",forum using default port 8080")
	}
	return ":" + args[0], nil
}

// validPort checks if the port is valid
func validPort(Pr string) bool {
	Port, err := strconv.Atoi(Pr)
	if err != nil {
		return false
	}

	minPort := 1024
	maxPort := 65535
	if Port >= minPort && Port <= maxPort {
		return true
	}
	return false
}
