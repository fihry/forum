package utils

import (
	"encoding/json"
	"errors"
	"flag"
	"os"
	"strconv"
)

// Config struct
type Config struct {
	Server Server `json:"server"`
}

type Server struct {
	DefaultPort string `json:"DefaultPort"`
	Host        string `json:"Host"`
}

func LoadConfig() Config {
	file, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
	config.Server.DefaultPort, err = GetPort(config.Server.DefaultPort)
	if err != nil {
		panic(err)
	}
	return config
}

// GetPort gets the port from the command line arguments
func GetPort(DefaultPort string) (string, error) {
	// get the port
	port := flag.String("p", DefaultPort, "port to run the server on")
	flag.Parse()

	if !ValidPort(*port) {
		return *port, errors.New("invalid port number")
	}
	return *port, nil
}

// validPort checks if the port is valid
func ValidPort(Pr string) bool {
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
