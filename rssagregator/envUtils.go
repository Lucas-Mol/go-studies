package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadingFromEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not defined in the environment")
	}
	return port
}
