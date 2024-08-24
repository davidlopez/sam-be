package utils

import (
	"os"
	"strings"
)

func GetPort() string {
	port := ":3000"

	val, ok := os.LookupEnv("PORT")
	if ok && strings.Contains(val, ":") {
		port = val
	} else if ok {
		port = ":" + val
	}

	return port
}
