package config

import "github.com/joho/godotenv"
import "log"

// Bootstrap - load initial configuration vars
func Bootstrap() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file in environment.go")
	}
}
