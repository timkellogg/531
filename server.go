package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// initialize db
	router := NewRouter()
	handler := cors.Default().Handler(router)

	fmt.Println("Server up on PORT")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
