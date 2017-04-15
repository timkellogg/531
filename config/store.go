package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

// Store - reference to db
type Store struct {
	DB *gorm.DB
}

// InitDb - bootstrap db connection
func (i *Store) InitDb() {
	var err error
	dbHost := os.Getenv("DATABASE_HOST")
	dbUser := os.Getenv("DATABASE_USER")
	dbName := os.Getenv("DATABASE_NAME")
	dbSslm := os.Getenv("DATABASE_SSL_MODE")
	dbPass := os.Getenv("DATABASE_PASS")

	i.DB, err = gorm.Open("postgres", "host=%s user=%s dbname=%s, sslmode=%s, dbPass", dbHost, dbUser, dbName, dbSslm, dbPass)
	if err != nil {
		log.Fatal("Error opening database connection in store.go")
	}
}
