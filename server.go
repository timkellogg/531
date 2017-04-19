package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/timkellogg/531/server/config"
)

func main() {
	config.Bootstrap()
	config.Database.InitDb()

	router := NewRouter()
	handler := cors.Default().Handler(router)

	fmt.Println("Server up on PORT")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
