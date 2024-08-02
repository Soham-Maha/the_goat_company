package utils

import (
	"log"
	"os"
)

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Warning: Environment variable %s is not set", key)
	}
	return value
}
