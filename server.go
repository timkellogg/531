package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/timkellogg/531/server/config"
	"github.com/timkellogg/531/server/models"
)

func main() {
	bootstrap()

	router := NewRouter()

	handler := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:4200"},
	}).Handler(router)

	port := ":" + os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	fmt.Println("Server up on " + port)
	log.Fatal(http.ListenAndServe(port, handler))
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
