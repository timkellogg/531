package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/timkellogg/531/server/models"
)

// Store - reference to db
type Store struct {
	DB *gorm.DB
}

// Database - reference to database used for reading/writing SQL trans
var Database = Store{}

// InitDb - bootstrap db connection
func (i *Store) InitDb() {
	var err error

	i.DB, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	i.DB.LogMode(true)
	i.CreateDb()
}

// CreateDb - creates table
func (i *Store) CreateDb() {
	// Migrate Tables
	i.DB.AutoMigrate(&models.User{})
	i.DB.AutoMigrate(&models.Program{})

	// Add Indices
	i.DB.Model(&models.User{}).AddUniqueIndex("idx_user_id", "id")
	i.DB.Model(&models.User{}).AddUniqueIndex("idx_user_email", "email")
}
