package config

import (
	"os"
)

var PORT string

func init() {
	portStr := os.Getenv("API_PORT")
	if portStr != "" {
		PORT = portStr
	} else {
		PORT = "3000"
	}
}
