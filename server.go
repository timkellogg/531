package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/timkellogg/531/server/config"
	"github.com/timkellogg/531/server/models"
)

func main() {
	bootstrap()

	router := NewRouter()
	handler := cors.Default().Handler(router)

	fmt.Println("Server up on PORT")
	log.Fatal(http.ListenAndServe(":3000", handler))
}

func bootstrap() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	config.Database.InitDb()
	store := config.Database.DB

	store.AutoMigrate(&models.User{})
	store.AutoMigrate(&models.Program{})

	// Add Indices
	store.Model(&models.User{}).AddUniqueIndex("idx_user_id", "id")
	store.Model(&models.User{}).AddUniqueIndex("idx_user_email", "email")
}
