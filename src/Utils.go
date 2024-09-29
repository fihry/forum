package forum

import (
	"errors"
	"os"
	"strconv"
)

// GetPort gets the port from the command line arguments
func GetPort() (string, error) {
	Argv := os.Args[1:]
	DefaultPort := ":8080"
	if len(Argv) == 0 {
		return DefaultPort, errors.New("\033[33mno port provided,forum using default port 8080\033[0m")
	} else if len(Argv) != 1 || !validPort(Argv[0]) {
		return DefaultPort, errors.New("invalid port :" + Argv[0] + ",forum using default port 8080")
	}
	return ":" + Argv[0], nil
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
