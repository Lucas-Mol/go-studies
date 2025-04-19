package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadingFromEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnvVar(envVar string) string {
	port := os.Getenv(envVar)
	if port == "" {
		log.Fatal(fmt.Sprintf("%s is not defined in the environment", envVar))
	}
	return port
}
